@echo off

REM ex1) test.bat test
REM ex2) test.bat /verbose true test
REM ex3) test.bat /verbose true cover

set GOPATH=E:\works\go
set GOROOT=C:\Go\

:loop
IF NOT "%1"=="" (
    IF /I "%1"=="/verbose" (
        set verbose=%2
        echo * Verbose     : %verbose%
        SHIFT
    )

    set command=%1
    SHIFT
    GOTO :loop
)
:theend

echo * Command     : %command%

if "%command%" == "test" (
    if "%verbose%" == "true" (
        go test -v .\...
    ) else (
        go test .\...
    )

    EXIT /B 0
)

if "%command%" == "cover" (
    if "%verbose%" == "true" (
        go test -coverprofile=coverage.out .\...
        go tool cover -func=coverage.out
        go tool cover -html=coverage.out
    ) else (
        go test -coverprofile=coverage.out .\...
    )

    EXIT /B 0
)

ECHO "invalid command"
EXIT /B 1