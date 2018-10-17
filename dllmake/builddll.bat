go build -ldflags "-s -w" -buildmode=c-shared -o exportgo.dll dllmake.go

pause