default: build

BUILD_VERSION=${BUILD}
PROJECT=k8s-load

clean:
	rm $(PROJECT)*
	rm -r out

test:
	go test -v

build-multi:
	gox -osarch="linux/amd64 darwin/amd64" -ldflags "-X main.buildVersion=$(BUILD_VERSION) -extldflags \"-static\" " -tags netgo -output="out/$(PROJECT)_{{.OS}}_{{.Arch}}"

build:
	go build -ldflags "-X main.buildVersion=$(BUILD_VERSION) -extldflags \"-static\" " -tags netgo 

docker-build:
	docker build --build-arg build=$(BUILD_VERSION) -t quay.io/mad01/$(PROJECT):$(BUILD_VERSION) --file Dockerfile-beta .

docker-push:
	docker push quay.io/mad01/$(PROJECT):$(BUILD_VERSION)

docker-login:
	docker login -u $(QUAY_LOGIN) -p="$(QUAY_PASSWORD)" quay.io

build-publish-gh:
	ghr -u mad01 $(BUILD_VERSION) out
