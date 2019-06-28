output kops_sg_id {
    value = "${aws_security_group.kops.id}"
    description = "The ID of kops Security Group"
}