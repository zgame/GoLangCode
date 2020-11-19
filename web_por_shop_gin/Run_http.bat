for /f %%i in ('powershell -c "Get-Date -uformat '%%Y%%m%%d'"') do (
    set "Today=%%i"
)

portia_shop_%Today%.exe -https=0