del /q out
cd ..
go build -ldflags "-s -w" -o _build\release\
cd _build
upx -9 release\mtools-backend.exe