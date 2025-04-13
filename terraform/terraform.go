package terraform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/biograph-health/interview-platform-external-asrtdd/config"
)

const terraformDir = "terraform"

// Setup sets up the terraform directory with the necessary files
func Setup(cfg *config.Config) error {
	// Create the terraform directory if it doesn't exist
	if err := os.MkdirAll(terraformDir, 0755); err != nil {
		return fmt.Errorf("failed to create terraform directory: %w", err)
	}
	tfvars := fmt.Sprintf(`
minio_server   = "%s"
minio_user     = "%s"
minio_password = "%s"
minio_ssl      = %t
bucket_name    = "%s"
file_content   = "%s"
`, cfg.Minio.Server, cfg.Minio.User, cfg.Minio.Password, cfg.Minio.SSL, cfg.BucketName, cfg.FileContent)

	err := os.WriteFile(filepath.Join(terraformDir, "terraform.tfvars"), []byte(tfvars), 0644)
	if err != nil {
		return fmt.Errorf("failed to write terraform.tfvars file: %w", err)
	}
	return nil
}

func Init() error {
	cmd := exec.Command("terraform", "init")
	cmd.Dir = terraformDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
	// err := cmd.Run()
	// if err != nil {
	//     return fmt.Errorf("error initializing terraform: %w", err)
	// }
	// m.logger.Info("Terraform initialized successfully")
	// return nil
}

func Apply() error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = terraformDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Destroy() error {
	cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd.Dir = terraformDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
