FROM registry.ci.openshift.org/open-cluster-management/builder:go1.17-linux AS builder

WORKDIR /tmp

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -i -o hostname main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

RUN microdnf update -y && microdnf clean all

WORKDIR /tmp

COPY --from=builder /tmp/hostname hostname

ENTRYPOINT ["/tmp/hostname"]
