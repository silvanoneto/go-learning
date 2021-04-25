FROM golang:alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v cmd/go_learning.go

RUN go build -o ./go-learning cmd/go_learning.go

FROM alpine
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/go-learning /app/go-learning

COPY --from=build_base /tmp/app/examples /app/examples

WORKDIR /app

CMD ["./go-learning"]