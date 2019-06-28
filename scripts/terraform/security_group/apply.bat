@echo off

terraform plan -out=planfile -var-file="..\common\variable.tfvars"
if errorlevel 1 (
    pause
    goto :EOF
)

terraform apply "planfile"
if errorlevel 1 (
    pause
    goto :EOF
)