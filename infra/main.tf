# version.tf
terraform {
  required_version = "~> 1.4"
  backend "s3" {
    bucket = "espresso-terraform-state"
    region = "us-east-1"
    key    = "terraform/chatserver"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.60"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  # shared_config_files      = ["/Users/$USER/.aws/config"]
  # shared_credentials_files = ["/Users/$USER/.aws/credentials"]
  # profile                  = "Espresso"
}


resource "aws_security_group" "chatserver" {
  name_prefix = "chatserver-"
  description = "Security group for chat server"

  vpc_id = "vpc-01c711da1a5c01a70"

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "chatserver" {
  ami           = "ami-0323c3dd2da7fb37d"
  instance_type = "t2.micro"

  subnet_id                   = "subnet-0573fca7a512a2820"
  associate_public_ip_address = true

  vpc_security_group_ids = [aws_security_group.chatserver.id]

  key_name = "espresso-useast1"

}

output "chatserver_public_ip" {
  value = aws_instance.chatserver.public_ip
}

