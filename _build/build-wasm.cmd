del /q out
cd ..
set GOOS=js
set GOARCH=wasm
go build -ldflags "-s -w" -o _build\release\
cd _build
upx -9 release\mtools-backend.exe