// Copyright Contributors to the Open Cluster Management project
package init

import (
	"fmt"

	"open-cluster-management.io/clusteradm/pkg/helpers"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	genericclioptionsclusteradm "open-cluster-management.io/clusteradm/pkg/genericclioptions"
)

var example = `
# Init the hub
%[1]s init
`

// NewCmd ...
func NewCmd(clusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(clusteradmFlags, streams)

	cmd := &cobra.Command{
		Use:          "init",
		Short:        "init the hub",
		Example:      fmt.Sprintf(example, helpers.GetExampleHeader()),
		SilenceUsage: true,
		PreRun: func(c *cobra.Command, args []string) {
			helpers.DryRunMessage(o.ClusteradmFlags.DryRun)
		},
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(c, args); err != nil {
				return err
			}
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&o.outputFile, "output-file", "", "The generated resources will be copied in the specified file")
	cmd.Flags().BoolVar(&o.useBootstrapToken, "use-bootstrap-token", false, "If set then the boostrap token will used instead of a service account token")
	cmd.Flags().BoolVar(&o.force, "force", false, "If set then the hub will be reinitialized")
	cmd.Flags().StringVar(&o.ImageRepo, "image-repo", "quay.io", "replace the default image repo quay.io with provided image")
	return cmd
}
