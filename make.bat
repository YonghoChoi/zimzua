@echo off
    REM ex) make.bat /app zimzua-api /platform linux build

FOR /F "tokens=*" %%g IN ('git rev-parse --short HEAD') do (SET revVersion=%%g)
set "timex=%time: =%"
set buildTime=%date%T%timex%
set version=%revVersion%.%buildTime%
set baseDir=%cd%
set debugMode=false

:loop
IF NOT "%1"=="" (
    IF /I "%1"=="/app" (
        set app=%2
        SHIFT
    )

    IF /I "%1"=="/platform" (
        set platform=%2
        SHIFT
    )

    IF /I "%1"=="/debug" (
        set debugMode=true
        if not "%2"=="" (
            SHIFT
        )
    )

    set command=%1
    SHIFT
    GOTO :loop
)
:theend

if "%app%" == "" (
    ECHO "app is Empty. Please input /app argument."
    EXIT /B 1
)

if "%platform%"=="" (
    ECHO "You have not entered a platform. Select Windows by default."
    set platform=windows
)

if "%command%" == "" (
    ECHO "command is Empty. Please input command argument."
    EXIT /B 1
)

set output=bin\%app%
set src=cmd\%app%

echo * Platform     : %platform%
echo * Build Time   : %buildTime%
echo * Git Revision : %revVersion%
echo * Output Dir   : %output%
echo * Build App    : %app%
echo * Command      : %command%

IF /I "%command%"=="build" (
    @echo * delete output directory
    rmdir /S /Q %output%
    cd %src%

    IF /I "%platform%"=="linux" (
        IF /I "%debugMode%"=="true" (
            @echo * build debug binary
            set GOARCH=amd64
            set GOOS=linux
            set GOHOSTOS=windows

            go build -gcflags "all=-N -l" -o %baseDir%\%output%\%app%-debug
            goto finishbuild
        )

        @echo * build linux binaries
        set GOARCH=amd64
        set GOOS=linux

        go build -o %baseDir%\%output%\%app%
        goto finishbuild
    )

    IF /I "%platform%"=="windows" (
        @echo * build windows binaries
        set GOARCH=amd64
        set GOOS=windows

        go build -o %baseDir%\%output%\%app%.exe
        goto finishbuild
    )

    :finishbuild

    @echo * build complete
)

EXIT /B 0