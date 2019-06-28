resource "aws_security_group" "zimzua" {
  name        = "zimzua"
  description = "docker based zimzua server"

  vpc_id = "${var.vpc}"

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "zimzua"
  }
}

resource "aws_security_group_rule" "zimzua_ssh" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "TCP"
  cidr_blocks       = ["${var.vpn_ip}"]
  security_group_id = "${aws_security_group.zimzua.id}"

  lifecycle { create_before_destroy = true }
}

resource "aws_security_group_rule" "zimzua_web" {
  type              = "ingress"
  from_port         = 8000
  to_port           = 8000
  protocol          = "TCP"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = "${aws_security_group.zimzua.id}"

  lifecycle { create_before_destroy = true }
}
