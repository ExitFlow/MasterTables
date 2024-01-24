FROM golang:latest AS build

WORKDIR /src/
COPY . /src/

RUN CGO_ENABLED=0 go build -o mastertables_beta

EXPOSE 80

CMD ["/src/mastertables_beta"]
