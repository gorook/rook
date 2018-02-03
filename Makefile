install:
	go install -v
test:
	go test -cover -race ./...
testv:
	go test -v -cover -race ./...
lint: prepare
	gometalinter --vendor --deadline=1m ./...
prepare:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
