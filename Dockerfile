FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o bin/ldjam-rank .

FROM heroku/heroku:18
COPY --from=build bin/ldjam-rank bin/ldjam-rank
CMD ["/bin/ldjam-rank"]