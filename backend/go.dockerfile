#build stage
FROM golang:1.16.0-alpine3.13 AS builder
#Backend dir context
ARG BOT_USER
ARG BOT_PRIVATE_TOKEN
ARG CI_PROJECT_NAMESPACE
RUN apk add --no-cache ca-certificates git

WORKDIR /go/src/gitlab.com/mcuc/monorepo/backend
ADD go.mod /go/src/gitlab.com/mcuc/monorepo/backend/go.mod
RUN git config --global url."https://${BOT_USER}:${BOT_PRIVATE_TOKEN}@gitlab.com/".insteadOf https://gitlab.com/ &&\
    go env -w GOPRIVATE=gitlab.com && \
    go mod download && cp go.sum /tmp/go.sum
ADD . /go/src/gitlab.com/mcuc/monorepo/backend
ARG service
RUN cp /tmp/go.sum ./go.sum && go build -o ./$service/cmd/main ./$service/cmd/main.go && mkdir -p ./$service/scripts

#final stage
FROM alpine:3.13
ARG env=dev
ARG service
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /go/src/gitlab.com/mcuc/monorepo/backend/$service/cmd/main /app/server
COPY --from=builder /go/src/gitlab.com/mcuc/monorepo/backend/$service/configs /app
COPY --from=builder /go/src/gitlab.com/mcuc/monorepo/backend/wait-for-it.sh ./wait-for-it.sh
COPY --from=builder /go/src/gitlab.com/mcuc/monorepo/backend/$service/scripts /app/scripts
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]