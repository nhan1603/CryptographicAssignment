FROM golang:1.22

RUN apt-get update

RUN GO111MODULE=on go install golang.org/x/tools/cmd/goimports@latest
RUN GO111MODULE=on go install github.com/vektra/mockery/v2@v2.20.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@latest && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest \
