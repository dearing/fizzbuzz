## BUILDTIME
FROM golang:latest as builder
LABEL maintainer="jacob.dearing@gmail.com"
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

## RUNTIME
FROM alpine:latest  
RUN apk --no-cache add ca-certificates nano curl wget
WORKDIR /root/
COPY --from=builder /app/main .
#EXPOSE 3300
CMD ["./main"] 
