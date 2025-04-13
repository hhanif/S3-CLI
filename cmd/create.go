package cmd

import (
	"fmt"
	"log"

	"github.com/biograph-health/interview-platform-external-asrtdd/config"
	"github.com/biograph-health/interview-platform-external-asrtdd/terraform"
	"github.com/spf13/cobra"
)

type createCmd struct {
	cmd *cobra.Command
}

func newCreateCmd() *createCmd {
	root := &createCmd{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Provision infrastructure based on configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Loading infrastructure...")
			cfg, err := config.LoadConfig(configFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}
			log.Println("Setting up Terraform configuration...")
			if err := terraform.Setup(cfg); err != nil {
				return fmt.Errorf("failed to setup terraform: %w", err)
			}
			log.Println("Initializing Terraform...")
			if err := terraform.Init(); err != nil {
				return fmt.Errorf("failed to initialize terraform: %w", err)
			}
			log.Println("Applying Terraform configuration...")
			if err := terraform.Apply(); err != nil {
				return fmt.Errorf("failed to apply terraform: %w", err)
			}
			log.Println("Successfully created bucket '%s' and sample file", cfg.BucketName)
			return nil
		},
	}

	root.cmd = cmd
	return root
}
