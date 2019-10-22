rice embed-go -i github.com/Andyfoo/go-h5speedtest
go mod tidy
go build -ldflags="-s -w"
