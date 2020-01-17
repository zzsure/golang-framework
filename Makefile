MAIN_PKG:=golang-framework
MAIN_PREFIX=$(dir $(MAIN_PKG))
MAIN=$(subst $(MAIN_PREFIX), , $(MAIN_PKG))
BIN=$(strip $(MAIN))

export GOPATH=$(shell pwd)/../../../../
export AZBIT_KUBERNETES_IDC=suzhou

build:
	go build -tags=jsoniter -x -o run/$(BIN) gitlab.azbit.cn/web/$(MAIN_PKG)

dev:
	go run main.go $(ARG)

run: build
	cd run && ./$(BIN) $(ARG)

init:
	cd run && TARGET='run' ARG='init' docker-compose run --rm golang-framework-devel

docker-build:
	#cd run && \
	#TARGET='build' ARG='server' docker-compose run --rm golang-framework-devel && cp $(BIN) ../build/ &&
	cd run && cp $(BIN) ../build/ && \
	cd ../build && \
	docker build -t zzsure/golang-framework:$(TAG) . && \
	docker push zzsure/golang-framework:$(TAG)

.PHONY: build
