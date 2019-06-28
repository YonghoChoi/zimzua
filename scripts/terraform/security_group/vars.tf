variable "region" {}
variable "vpc" {}
variable "subnet_1" {}
variable "subnet_2" {}
variable "az_1" {}
variable "az_2" {}
variable "amazon_linux_ami" {}
variable "ubuntu_ami" {}

variable "vpn_ip" {
    default = "119.206.206.251/32"
    description = "VPN IP for accessing the EC2 instance"
}