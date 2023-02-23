go env -w GOARCH=arm
go env -w GOARM=7
go env -w GOOS=linux
go env -w CGO_ENABLED=0
go build -ldflags "-s -w" -buildvcs=false  -o linux-armv7-keepLogin

go env -w GOARCH=amd64
go env -w GOOS=linux
go build -ldflags "-s -w" -buildvcs=false  -o linux-amd64-keepLogin

go env -w GOARCH=amd64
go env -w GOOS=windows
go build -ldflags "-s -w" -buildvcs=false  -o windows-amd64-keepLogin.exe