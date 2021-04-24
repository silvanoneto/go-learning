FROM golang:alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/go-learning

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v cmd/go_learning.go

RUN go build -o ./out/go-learning cmd/go_learning.go

FROM alpine
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-learning/out/go-learning /app/go-learning

COPY --from=build_base /tmp/go-learning/_data_examples /app/_data_examples

CMD ["/app/go-learning"]