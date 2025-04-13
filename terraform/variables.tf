variable "minio_server" {
  description = "Minio host and port"
  type        = string
  default     = "127.0.0.1:9000"
}

variable "minio_user" {
  description = "Minio user"
  type        = string
}

variable "minio_password" {
  description = "Minio password"
  type        = string
}

variable "minio_ssl" {
  description = "Minio SSL enabled"
  type        = bool
  default     = false
}

variable "bucket_name" {
  description = "Name of the bucket"
  type        = string
}

variable "file_content" {
  description = "Content type of the object"
  type        = string
}
