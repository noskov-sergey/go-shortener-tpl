tests:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration1 -binary-path=shortener

tests2:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta -test.v -test.run=^TestIteration2 -source-path=.

build:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
