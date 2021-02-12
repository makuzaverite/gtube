FROM golang:alpine3.12

# Meta information
LABEL mantainer="Makuza Mugabo Verite"
LABEL description="Youtube video downloader"

# Copy files into a container
COPY . /usr/src/app
WORKDIR /usr/src/app


# Install dependancies
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app .

FROM alpine:latest
COPY --from=builder /usr/src/app/app .
CMD [ "./app" ]
