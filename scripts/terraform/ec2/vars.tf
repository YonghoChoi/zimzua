variable "region" {}
variable "vpc" {}
variable "subnet_1" {}
variable "subnet_2" {}
variable "az_1" {}
variable "az_2" {}
variable "amazon_linux_ami" {}
variable "ubuntu_ami" {}

variable "ec2_passwd_prameter_name" {
   default = "kops-ubuntu-password"
   description = "Parameter name of AWS Parameter Store"
}