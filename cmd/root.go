package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var configFile string

type rootCmd struct {
	cmd *cobra.Command
}

func Execute(args []string) {
	rootCmd := newRootCmd()
	rootCmd.cmd.SetArgs(args)
	if err := rootCmd.cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func newRootCmd() *rootCmd {
	root := &rootCmd{}
	cmd := &cobra.Command{
		Use:               "infra-cli",
		Short:             "Infrastructure Automation CLI Tool",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "Path to the configuration file")
	_ = cmd.MarkFlagFilename("config", "yaml", "yml")

	cmd.AddCommand(
		newCreateCmd().cmd,
		newDeleteCmd().cmd,
	)
	root.cmd = cmd
	return root
}
