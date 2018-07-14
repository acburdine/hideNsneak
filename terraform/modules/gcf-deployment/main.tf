provider "google" {
  credentials = "${file(var.google_credentials_path)}"
  project     = "${var.gcp_project}"
}

resource "random_string" "bucket_name" {
  length  = 8
  special = false
  number  = false
  upper   = false
}

resource "google_storage_bucket" "bucket" {
  name = "hidensneak${random_string.bucket_name.result}"

  force_destroy = true
}

data "archive_file" "http_trigger" {
  type        = "zip"
  output_path = "${path.module}/files/redirector.zip"

  source {
    content  = "${file(var.redirector_file)}"
    filename = "index.js"
  }

  source {
    content  = "${file(var.package_file)}"
    filename = "package.json"
  }
}

resource "google_storage_bucket_object" "archive" {
  name       = "redirector.zip"
  bucket     = "${google_storage_bucket.bucket.name}"
  source     = "${path.module}/files/redirector.zip"
  depends_on = ["data.archive_file.http_trigger"]
}

resource "google_cloudfunctions_function" "hidensneak" {
  name                  = "${var.function_name}"
  entry_point           = "redirector"
  available_memory_mb   = 128
  timeout               = 61
  project               = "${var.gcp_project}"
  region                = "${var.region}"
  trigger_http          = "${var.enabled}"
  source_archive_bucket = "${google_storage_bucket.bucket.name}"
  source_archive_object = "${google_storage_bucket_object.archive.name}"

  labels = "${var.labels}"

  depends_on = ["google_storage_bucket_object.archive"]
}
