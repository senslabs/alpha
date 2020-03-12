FROM golang:1.14.0-alpine3.11
WORKDIR /workspace

RUN ls

RUN cd sens/datastore/generated/api/rest/main

RUN CGO_ENABLED=0 GOOS=linux RUN go build -o /datastore 

ENTRYPOINT [ "/datastore" ]