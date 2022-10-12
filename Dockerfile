FROM golang:latest AS builder
WORKDIR /app/emp_server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED 0
RUN go build -o app -tags netgo ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/emp_server/app .
ENV PORT 3000
CMD [ "/app/app" ]
