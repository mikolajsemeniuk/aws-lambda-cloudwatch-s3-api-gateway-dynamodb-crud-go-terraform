terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region                   = var.region
  shared_credentials_files = ["./credentials"]
  profile                  = "default"
}

variable "region" {
  description = "The AWS region"
  default     = "eu-central-1"
}
