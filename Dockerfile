FROM golang:alpine3.12

# Meta information
LABEL mantainer="Makuza Mugabo Verite"
LABEL description="Youtube video downloader"

# Copy files into a container
WORKDIR /usr/src/app


#Download Dependancies
COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .


# Install dependancies
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app .

CMD [ "./app" ]
