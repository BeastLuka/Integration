FROM golang:1.23

WORKDIR /app

COPY ./Booking-app .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /Booking-app

EXPOSE 3000

CMD ["/Booking-app"]

