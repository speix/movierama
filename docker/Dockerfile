FROM golang

ADD . /go/src/github.com/speix/movierama

WORKDIR /go/src/github.com/speix/movierama

ARG MR_SERVER_PORT

ENV MR_SERVER_PORT $MR_SERVER_PORT

RUN go get github.com/jmoiron/sqlx
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/twinj/uuid
RUN go get github.com/gorilla/mux

RUN go install github.com/speix/movierama

ENTRYPOINT /go/bin/movierama

EXPOSE $MR_SERVER_PORT