FROM golang:1.16-alpine AS builder

# Copy the code from the host and compile it
WORKDIR /github.com/deissh/highloadcup
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static" -w' -o ./bin/server ./cmd && mv bin/* /


FROM scratch

WORKDIR /app
COPY --from=builder /server /server

EXPOSE 8000
ENTRYPOINT ["/server"]