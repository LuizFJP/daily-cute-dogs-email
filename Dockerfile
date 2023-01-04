FROM golang:1.19
WORKDIR /home/luizportel4/projects/daily-cute-dogs-email

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .
CMD ["./main"]