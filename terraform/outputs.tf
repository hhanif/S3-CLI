output "minio_id" {
  value = minio_s3_bucket.default.id
}

output "minio_url" {
  value = minio_s3_bucket.default.bucket_domain_name
}

output "minio_object_id" {
  value = minio_s3_object.txt_file.id
}
