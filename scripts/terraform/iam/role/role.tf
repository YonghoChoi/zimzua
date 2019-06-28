resource "aws_iam_role" "zimzua_role" {
  name = "zimzua_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

  tags = {
    Name = "zimzua_role"
  }
}

# 인스턴스 프로파일은 IAM 역할을 위한 컨테이너로서 인스턴스 시작 시 EC2 인스턴스에 역할 정보를 전달하는 데 사용됨
# EC2 인스턴스는 가상화된 리소스이기 때문에 권한을 부여하기 위해 임시 자격 증명이 필요한데 이를 수행해주는 것이 인스턴스 프로파일
# instance profile은 콘솔에서 제거할 수 없다. awscli를 사용하여 제거한다. (aws iam delete-instance-profile --instance-profile-name zimzua_profile)
resource "aws_iam_instance_profile" "zimzua_profile" {
  name = "zimzua_profile"
  role = "${aws_iam_role.zimzua_role.name}"
}

resource "aws_iam_role_policy" "zimzua_policy" {
  name = "zimzua_policy"
  role = "${aws_iam_role.zimzua_role.id}"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeTags",
                "ecs:CreateCluster",
                "ecs:DeregisterContainerInstance",
                "ecs:DiscoverPollEndpoint",
                "ecs:Poll",
                "ecs:RegisterContainerInstance",
                "ecs:StartTelemetrySession",
                "ecs:UpdateContainerInstancesState",
                "ecs:Submit*",
                "ecr:GetAuthorizationToken",
                "ecr:BatchCheckLayerAvailability",
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}