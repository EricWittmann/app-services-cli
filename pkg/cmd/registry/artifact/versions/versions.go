package versions

import (
	"context"
	"errors"
	"fmt"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"

	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type Options struct {
	artifact     string
	group        string
	outputFormat string

	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

func NewVersionsCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "versions",
		Short: "Get latest artifact versions by artifact-id and group",
		Long:  "Get latest artifact versions by specifying group and artifact-id",
		Example: `
## Get latest artifact versions for default group
rhoas service-registry artifact versions --artifact-id=my-artifact

## Get latest artifact versions for my-group group
rhoas service-registry artifact versions --artifact-id=my-artifact --group mygroup 
		`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.artifact == "" {
				return errors.New("artifact id is required. Please specify artifact by using --artifact-id flag")
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service registry selected. Please specify registry by using --instance-id flag")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runGet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Artifact group")
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Output format (json, yaml, yml)")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == util.DefaultArtifactGroup {
		logger.Info("Group was not specified. Using 'default' artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	logger.Info("Fetching artifact versions")

	ctx := context.Background()
	request := dataAPI.VersionsApi.ListArtifactVersions(ctx, opts.group, opts.artifact)
	response, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	logger.Info("Successfully fetched artifact versions")

	dump.PrintDataInFormat(opts.outputFormat, response, opts.IO.Out)

	return nil
}