data "terraform_remote_state" "sg_data" {
    backend = "local"
    config = {
        path = "${path.module}/../security_group/terraform.tfstate"
    }
}

data "terraform_remote_state" "iam_role_data" {
    backend = "local"
    config = {
        path = "${path.module}/../iam/role/terraform.tfstate"
    }
}

data "aws_ssm_parameter" "ec2-password" {
  name = "${var.ec2_passwd_prameter_name}"
}

resource "aws_instance" "zimzua" {
  ami                  = "${var.ubuntu_ami}"
  availability_zone    = "${var.az_1}"
  instance_type        = "t2.micro"
  key_name             = "${var.key_pair}"
  iam_instance_profile = "${data.terraform_remote_state.iam_role_data.outputs.zimzua_iam_instance_profile_name}"
  vpc_security_group_ids = [
    "${data.terraform_remote_state.sg_data.outputs.zimzua_sg_id}",
  ]

  subnet_id                   = "${var.subnet_1}"
  associate_public_ip_address = true
  user_data = <<EOF
#!/bin/bash
echo "ubuntu:${data.aws_ssm_parameter.ec2-password.name}" | chpasswd
sed -i "/^[^#]*PasswordAuthentication[[:space:]]no/c\PasswordAuthentication yes" /etc/ssh/sshd_config
service sshd restart
  EOF

  tags = {
    Name = "zimzua"
  }

  connection {
    type        = "ssh"
    user        = "ubuntu"
    password    = "${data.aws_ssm_parameter.ec2-password.name}"
    host        = "${self.public_ip}"
  }

  rovisioner "file" {
    source      = "../../compose/docker-compose.yml"
    destination = "/home/ubuntu/docker-compose.yml"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common",
      "sudo sh -c \"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -\"",
      "sudo apt-key fingerprint 0EBFCD88",
      "sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
      "sudo apt-get update",
      "sudo apt-get install -y docker-ce docker-ce-cli containerd.io awscli",
      "sudo usermod -aG docker ubuntu",
      "sudo curl -L \"https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)\" -o /usr/local/bin/docker-compose",
      "sudo chmod +x /usr/local/bin/docker-compose",
      "sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose",
      "sudo docker-compose -f /home/ubuntu/docker-compose.yml up -d",
    ]
  }
}