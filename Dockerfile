FROM golang:1.14

WORKDIR /app
# COPY . . 

# RUN go build -o app
RUN go get -u github.com/cosmtrek/air

# CMD ["./app"]
ENTRYPOINT air