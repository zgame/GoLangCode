for /f %%i in ('powershell -c "Get-Date -uformat '%%Y%%m%%d'"') do (
    set "Today=%%i"
)

start "" portia_shop_%Today%.exe -https=1
start "" portia_shop_%Today%.exe -https=0