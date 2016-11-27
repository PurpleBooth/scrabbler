FROM golang:1.6-onbuild

RUN mkdir -p  /go/src/github.com/purplebooth/scrabbler
WORKDIR /go/src/github.com/purplebooth/scrabbler
COPY . /go/src/github.com/purplebooth/scrabbler
RUN (go get -v -d . && cd scrabbler-http && go get -v -d . && go install )
CMD ["/go/bin/scrabbler-http", "/go/src/github.com/purplebooth/scrabbler/wordlists/words.txt"]
EXPOSE 8080
