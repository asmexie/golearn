set cur=%~dp0
set GOARCH=amd64
set GOOS=windows
go build -o release/cool.exe -ldflags="-H windowsgui"
cd %cur%
pause