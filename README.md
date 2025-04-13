# Purpose

This code aims to simplify cloud infrastructure provisioning by developing a CLI tool that automates the
creation and deletion of S3 resources using Terraform.

## Prerequisites

- Docker Compose or equivalent

## Requirements

The CLI supports the following commands:

- create - Provisions a S3 bucket and object.
- delete - Destroys the infrastructure.

## Implementation Details

### Terraform Integration

- The CLI invokes Terraform commands (init, apply, destroy) using a system process (e.g., Python’s subprocess,
  Go’s os/exec, or equivalent in other languages).

```

## State Management

- Store Terraform state locally by default.


## Example CLI Usage

```bash
# Provision infrastructure
$ infra-cli create --config config.yaml

# Destroy infrastructure
$ infra-cli delete --config config.yaml
```
