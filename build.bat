cd infra-robot

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build 

mv infra-robot.exe ..\dist\
mv infra-robot ..\dist\

cd ..

cp -r config dist\
cp -r sample dist\
cp README.md dist\