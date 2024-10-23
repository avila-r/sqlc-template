resource "aws_ecr_repository" "Tasker ECR" {
  name                 = "tasker-app"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_iam_policy" "ec2_ecr_policy" {
  name        = var.ec2_to_ecr_policy
  path        = "/"
  description = "EC2 policy to access ECR"
  policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        "Effect" : "Allow",
        "Action" : [
          "ecr:*",
          "cloudtrail:LookupEvents"
        ],
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "iam:CreateServiceLinkedRole"
        ],
        "Resource" : "*",
        "Condition" : {
          "StringEquals" : {
            "iam:AWSServiceName" : [
              "replication.ecr.amazonaws.com"
            ]
          }
        }
      }
    ]
  })
}

resource "aws_iam_role" "ec2_ecr_role" {
  name = var.ec2_ecr_name
  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "name" {
  role       = aws_iam_role.ec2_ecr_role.name
  policy_arn = aws_iam_policy.ec2_ecr_policy.arn
}

resource "aws_iam_instance_profile" "profile" {
  name = var.ec2_instance_profile
  role = aws_iam_role.ec2_ecr_role.name
}

resource "aws_instance" "Tasker Server" {
  ami                    = "ami-06b21ccaeff8cd686"
  instance_type          = "t2.micro"
  
  tags = {
    Name = "Tasker App-Server"
  }

  user_data = <<-EOL
  #!/bin/bash -xe
  sudo yum update -y
  sudo amazon-linux-extras install docker -y
  sudo service docker start
  sudo usermod -a -G docker ec2-user
  EOL

}
