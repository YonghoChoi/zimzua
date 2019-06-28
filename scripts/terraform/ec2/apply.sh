#!/bin/bash
terraform init -var-file="../common/variable.tfvars"
if [ "$?" -ne "0" ]; then
  echo "fail plan"
  exit 1
fi

terraform plan -out=planfile -var-file="../common/variable.tfvars"
if [ "$?" -ne "0" ]; then
  echo "fail plan"
  exit 1
fi

terraform apply "planfile"
if [ "$?" -ne "0" ]; then
  echo "fail apply"
  exit 1
fi