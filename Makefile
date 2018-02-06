install:
	go install -v
test:
	go test -race ./...
testv:
	go test -v -cover -race ./...
lint: prepare
	gometalinter --vendor --deadline=1m --skip assets ./...
prepare:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
dep:
	go get -u github.com/golang/dep
	dep ensure
	dep prune
bindata:
	go-bindata -o assets/newsite/assets.go -pkg newsite -prefix "data/newsite" data/newsite/...
