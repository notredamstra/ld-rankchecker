FROM golang:stretch as build

# get nodejs and webpack
RUN apt-get update
RUN apt-get install -y git python jq curl
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN npm install webpack -g

RUN go build -o /bin/ldjam-rank .
RUN npm install
RUN npm run build

COPY . /app
WORKDIR /app

# Heroku setup
FROM heroku/heroku:18
COPY --from=build /bin/ldjam-rank /bin/ldjam-rank
RUN chmod a+x /bin/ldjam-rank
CMD ["/bin/ldjam-rank"]