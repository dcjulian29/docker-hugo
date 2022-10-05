@echo off
setlocal

pushd %~dp0

for /f "delims=" %%x in (version) do set VERSION=%%x

docker build --build-arg VERSION=%VERSION% -t dcjulian29/hugo:%VERSION% .

if %errorlevel% neq 0 GOTO FINAL

docker tag dcjulian29/hugo:%VERSION% dcjulian29/hugo:latest

:FINAL

popd

endlocal
