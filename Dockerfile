# syntax=docker/dockerfile:1
FROM golang:1-alpine AS builder
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build from source
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o "/echo"

# create a non-root user
RUN addgroup -S appgroup && adduser -S -G appgroup appuser

FROM scratch AS runner
WORKDIR /app

COPY --from=builder /echo /echo

# copy the non-user from builder stage
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# switch to the non-root user
USER appuser

EXPOSE 8080

CMD [ "/echo" ]
