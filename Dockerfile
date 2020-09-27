FROM golang:1.14-alpine AS builder
COPY . /src/
WORKDIR /src/cmd
RUN CGO_ENABLED=0 go build -o social-media-analysis-app

FROM alpine:latest  
WORKDIR /bin/
COPY --from=builder /src/cmd .
CMD ["/bin/social-media-analysis-app"]  