FROM golang:alpine as builder
WORKDIR /build 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o urn-lookup .
FROM scratch
COPY --from=builder /build/urn-lookup /app/
WORKDIR /app
ENV PATH=/app/:$PATH
ENV GIN_MODE=release
CMD ["urn-lookup"]