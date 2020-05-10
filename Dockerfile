FROM golang:stretch as build
COPY . /app
WORKDIR /app
MKDIR bin
RUN go build -o /ldjam-rank .

FROM heroku/heroku:18
COPY --from=build /ldjam-rank /ldjam-rank
CMD ["/bin/ldjam-rank"]