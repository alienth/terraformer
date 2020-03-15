package cmd

import (
	wavefront_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/wavefront"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdWavefrontImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "newrelic",
		Short: "Import current state to Terraform configuration from Wavefront",
		Long:  "Import current state to Terraform configuration from Wavefront",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newNewRelicProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newWavefrontProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "alert", "newrelic_dashboard=id1:id2:id4")
	return cmd
}

func newWavefrontProvider() terraform_utils.ProviderGenerator {
	return &wavefront_terraforming.WavefrontProvider{}
}
