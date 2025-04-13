resource "minio_s3_bucket" "default" {
  bucket = var.bucket_name
  acl    = "public"
}

resource "minio_s3_object" "txt_file" {
  depends_on   = [minio_s3_bucket.default]
  bucket_name  = minio_s3_bucket.default.bucket
  object_name  = "text.txt"
  content      = var.file_content
  content_type = "text/plain"
}
