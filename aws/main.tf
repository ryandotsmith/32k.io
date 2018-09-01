variable "ssh_public_key" {
  type    = "string"
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCzM2/wM4EzrudXaBKVSD8nT7et43pr5SzeGdyr6mPFZghLmDLoH1Mo55f4hnYZtt9So5JKMLUR1b5Eppc2mUg1j+hQIPtphuA0TRyAWFWfuTS7ppn+a7Up+Kd6DVPFebgvFENfY2BqjyzmkbzC2dPomZL/3oCfid6OPkSLs26oqO7SmbBvGnEjyQGjoN5ev6nzf78ba4mBoidL65PjkzBs3tRwRkAA8dLijvV/7O9PwL6AZPCznv3oy3Pc/URo0GxvuaI7IcrChB+cjJ4TjabsLqQ2YpnLheMLO1EQL8cO5kkFp+viK04qUGcx0InOfEABBrmG680qqMBx9ugAnrsv tf"
}

variable "vpc_name" {
  type    = "string"
  default = "32k"
}

variable "cidr_block" {
  type    = "string"
  default = "10.0.0.0/16"
}

provider "aws" {
  version = "~> 1.22"
  region  = "us-west-1"
}

resource "aws_vpc" "32k" {
  cidr_block           = "${var.cidr_block}"
  enable_dns_hostnames = true

  tags {
    Name = "${var.vpc_name}"
  }
}

resource "aws_internet_gateway" "32k" {
  vpc_id = "${aws_vpc.32k.id}"
}

resource "aws_route" "internet_access" {
  route_table_id         = "${aws_vpc.32k.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.32k.id}"
}

resource "aws_subnet" "32k" {
  vpc_id                  = "${aws_vpc.32k.id}"
  cidr_block              = "${cidrsubnet(var.cidr_block, 8, 1)}"
  map_public_ip_on_launch = true

  tags {
    Name = "${var.vpc_name}"
  }
}

resource "aws_security_group" "32k" {
  vpc_id      = "${aws_vpc.32k.id}"
  name        = "${var.vpc_name}"
  description = "SSH + Peer traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Public SSH access"
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Public peer access"
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Should redirect to https"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "All outgoing traffic allowed"
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_key_pair" "32k" {
  key_name   = "${var.vpc_name}-deployer-key"
  public_key = "${var.ssh_public_key}"
}

resource "aws_instance" "server32k" {
  connection {
    user = "ubuntu"
  }

  ami                    = "${data.aws_ami.ubuntu.id}"
  instance_type          = "t3.small"
  vpc_security_group_ids = ["${aws_security_group.32k.id}"]
  subnet_id              = "${aws_subnet.32k.id}"
  key_name               = "${aws_key_pair.32k.key_name}"

  root_block_device = {
    volume_size = 10
  }

  tags {
    Name = "${var.vpc_name}-1"
  }

  provisioner "local-exec" "build-server" {
    working_dir = "/Users/r/src/github.com/ryandotsmith/32k.io"

    environment {
      GOARCH = "amd64"
      GOOS   = "linux"
    }

    command = "go build -o /tmp/server ./cmd/server"
  }

  provisioner "file" {
    source      = "/tmp/server"
    destination = "/home/ubuntu/server"
  }

  provisioner "file" {
    source      = "./server.service"
    destination = "/home/ubuntu/server.service"
  }

  provisioner "file" {
    source      = "./server.socket"
    destination = "/home/ubuntu/server.socket"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /home/ubuntu/server",
      "sudo mv /home/ubuntu/server.service /etc/systemd/system/",
      "sudo mv /home/ubuntu/server.socket /etc/systemd/system/",
      "sudo systemctl daemon-reload",
      "sudo systemctl enable server.socket",
      "sudo systemctl enable server",
      "sudo systemctl start server.socket",
    ]
  }
}

resource "aws_eip" "32k" {
  instance = "${aws_instance.server32k.id}"
  vpc      = true
}

resource "aws_route53_zone" "32k" {
  name = "32k.io."
}

resource "aws_route53_record" "32k" {
  zone_id = "${aws_route53_zone.32k.zone_id}"
  name    = "server.32k.io."
  type    = "A"
  ttl     = "60"
  records = ["${aws_eip.32k.public_ip}"]
}
