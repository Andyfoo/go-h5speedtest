md _release
rice embed-go -i github.com/Andyfoo/go-h5speedtest
go mod tidy

set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%.exe

set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%.exe


set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%

set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%


set GOOS=darwin
set GOARCH=386
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%

set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%

set GOOS=linux
set GOARCH=arm
set GOARM=5
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%v%GOARM%

set GOOS=linux
set GOARCH=arm
set GOARM=6
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%v%GOARM%

set GOOS=linux
set GOARCH=arm
set GOARM=7
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%v%GOARM%

set GOOS=linux
set GOARCH=arm64
set GOARM=7
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%v%GOARM%

set GOOS=linux
set GOARCH=mips
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%


set GOOS=linux
set GOARCH=mips64
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%

set GOOS=linux
set GOARCH=mips64le
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%

set GOOS=linux
set GOARCH=mipsle
go build -ldflags="-s -w" -o _release/go-h5speedtest_%GOOS%_%GOARCH%
