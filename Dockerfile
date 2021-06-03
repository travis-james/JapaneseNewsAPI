FROM golang:1.15.11-alpine3.12

# Now set up the webapp.
WORKDIR /go/jpnewsapi
COPY ./ ./
# downloads the app dependencies and prints the progress to standard out.
RUN go mod download -x
RUN go build ./cmd/web

CMD [ "./web" ]

# docker build --no-cache -t jpnewsapi .
# docker run --env-file=jpnewsapi.env -p 4000:4000 jpnewsapi