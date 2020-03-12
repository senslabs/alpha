FROM golang:1.14.0-alpine3.11

RUN echo $(ls -l && pwd)
RUN echo $WORKDIR

WORKDIR /workspace

RUN echo $(ls -l && pwd)

RUN cd sens/datastore/generated/api/rest/main

RUN CGO_ENABLED=0 GOOS=linux RUN go build -o /datastore 

ENTRYPOINT [ "/datastore" ]