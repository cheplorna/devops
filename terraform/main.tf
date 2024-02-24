terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.38.0"
    }
  }
}

provider "aws" {
  profile= "default"
}

# Create a VPC
resource "aws_vpc" "test1" {
  cidr_block = "10.0.0.0/16"
}
