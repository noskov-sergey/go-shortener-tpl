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

tests9:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration9 -binary-path=shortener -source-path=. -file-storage-path=/tmp/mHkznYW

tests10:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration10 -binary-path=shortener -source-path=. -database-dsn=postgres://shortener:shortener@localhost:5439/shortener?sslmode=disable

tests14:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
	shortenertestbeta --test.v --test.run=^TestIteration14 -binary-path=shortener -database-dsn=postgres://shortener:shortener@localhost:5439/shortener?sslmode=disable

build:
	GOARCH=amd64 GOOS=windows go build -o cmd/shortener/shortener.exe cmd/shortener/main.go
