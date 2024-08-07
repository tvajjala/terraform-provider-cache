TEST?=$$(go list -buildvcs=false ./... | grep -v 'vendor')
#If you are accessing local plugins , make sure create .terraformrc file under home directory
HOSTNAME=tvajjala.github.io# make sure no space after each name
NAMESPACE=service
NAME=cache
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64#IMPTONANT update your platform by running `go version`
# you need to build for your platform 
# to know your platform run `go version`

# Use -buildvcs=false to always omit version control information, 
#or -buildvcs=true to error out if version control information is available but cannot be included due to a missing tool or ambiguous directory structure.
default: install
build:
	go build -buildvcs=false -o ${BINARY}
release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign
install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m