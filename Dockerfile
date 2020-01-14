FROM golang:latest as builder

ADD . /go/src/sflogapi
WORKDIR /go/src/sflogapi
RUN go get github.com/futurenda/google-auth-id-token-verifier
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/lib/pq
RUN go get github.com/joho/godotenv
RUN go get github.com/jinzhu/gorm
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go install sflogapi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/sflogapi/main .
COPY --from=builder /go/src/sflogapi/.env .
COPY --from=builder /go/src/sflogapi/pgserver.crt .
COPY --from=builder /go/src/sflogapi/pgserver.pem .
COPY --from=builder /go/src/sflogapi/apiserver.crt .
COPY --from=builder /go/src/sflogapi/apiserver.key .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 