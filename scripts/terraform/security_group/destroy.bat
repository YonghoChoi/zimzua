@echo off

terraform destroy -var-file="..\common\variable.tfvars"
if errorlevel 1 (
    pause
    goto :EOF
)