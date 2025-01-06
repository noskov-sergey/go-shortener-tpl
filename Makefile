tests:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration1 -binary-path=shortener

tests2:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta -test.v -test.run=^TestIteration2 -source-path=.

tests3:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration3 -binary-path=shortener -source-path=.

tests4:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration4 -binary-path=shortener -source-path=. --server-port=8667

tests6:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration6 -binary-path=shortener -source-path=.

tests7:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration7 -binary-path=shortener -source-path=.

build:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
