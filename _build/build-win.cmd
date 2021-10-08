del /q out
cd ..
go build -ldflags "-s -w" -o _build\release\
copy mtools-backend.db _build\release\
cd _build
upx -9 release\mtools-backend.exe