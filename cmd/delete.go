package cmd

import (
	"fmt"
	"log"

	"github.com/biograph-health/interview-platform-external-asrtdd/config"
	"github.com/biograph-health/interview-platform-external-asrtdd/terraform"
	"github.com/spf13/cobra"
)

type deleteCmd struct {
	cmd *cobra.Command
}

func newDeleteCmd() *deleteCmd {
	root := &deleteCmd{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Destroy provisioned infrastructure",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Loading infrastructure...")
			_, err := config.LoadConfig(configFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}
			log.Println("Destroying infrastructure...")
			if err := terraform.Destroy(); err != nil {
				return fmt.Errorf("failed to destroy terraform: %w", err)
			}
			log.Println("Infrastructure destroyed successfully.")
			return nil
		},
	}

	root.cmd = cmd
	return root
}

// docker -compose up -d
// go run main.go create --config config.yaml
//go run main.go delete --config config.yaml
