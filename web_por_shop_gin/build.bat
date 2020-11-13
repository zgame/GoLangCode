for /f %%i in ('powershell -c "Get-Date -uformat '%%Y%%m%%d'"') do (
    set "Today=%%i"
)
go build  -o portia_shop_%Today%.exe

