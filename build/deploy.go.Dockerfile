FROM --platform=linux/amd64 golang:1.22.6-alpine3.19 AS base
RUN apk --no-cache add \
    bash \
    build-base \
    git 

#################

FROM base AS builder

WORKDIR /reminoassignment/api

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 go build -o /reminoassignment/api/cmd/entrypoint ./cmd/entrypoint

###################

FROM --platform=linux/amd64 alpine:3.19
RUN apk --no-cache add \
    ca-certificates \
    tzdata
COPY --from=builder /reminoassignment/api/cmd/entrypoint /
COPY ./templates ./templates

RUN adduser -D -H -u 1000 reminoassignment
USER reminoassignment

EXPOSE 3001
CMD /entrypoint
