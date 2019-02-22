FROM golang:1.11 as build

ENV CGO_ENABLED=0

WORKDIR /go/src/app

# explicitly install dependencies to improve Docker re-build times
RUN go get -v github.com/kisielk/errcheck
RUN go get -v github.com/pkg/errors 
RUN go get -v github.com/calvinbrewer/section-sdk-go/api 
RUN go get -v github.com/stretchr/testify/assert 
RUN go get -v golang.org/x/lint/golint 
RUN go get -v gopkg.in/h2non/gock.v1

# Use specific version of terraform
RUN mkdir -p "${GOPATH}/src/github.com/hashicorp/" && \
    cd "${GOPATH}/src/github.com/hashicorp/" && \
      git clone --verbose --branch v0.11.10 --depth 1 https://github.com/hashicorp/terraform

WORKDIR /go/src/github.com/section-io/terraform-provider-section
COPY main.go .
COPY section ./section/

# Capture dependency versions
RUN \
  go list -f '{{ join .Imports "\n" }}' ./... \
  | xargs --max-lines=1 -I % go list -f '{{ .Dir }}' % \
  | xargs --max-lines=1 -I % bash -c 'cd %; git rev-parse --show-toplevel 2>/dev/null || true ' \
  | sort | uniq \
  | xargs --max-lines=1 -I % bash -c 'cd %; echo $(git rev-parse HEAD) %'

RUN gofmt -e -s -d /go/src/app 2>&1 | tee /gofmt.out && test ! -s /gofmt.out
RUN go tool vet /go/src/app
#RUN errcheck ./...

RUN go install ./...

RUN cd "/go/src/$(go list -e -f '{{.ImportComment}}')" && \
  go test -bench=. -v ./...

### END FROM build

FROM hashicorp/terraform:0.11.10

RUN mkdir -p /work/ && \
  printf 'providers {\n  section = "/go/bin/terraform-provider-section"\n}\n' >/root/.terraformrc

WORKDIR /work

COPY --from=build /go/bin/terraform-provider-section /go/bin/

COPY example.tf ./main.tf

RUN terraform fmt -diff -check ./

# https://www.terraform.io/docs/internals/debugging.html
ARG TF_LOG=WARN

RUN terraform init && \
  TF_LOG=DEBUG SECTION_USERNAME=fff SECTION_PASSWORD=fff terraform plan -out=a.tfplan

RUN TF_LOG=DEBUG SECTION_USERNAME=fff SECTION_PASSWORD=fff terraform apply a.tfplan
