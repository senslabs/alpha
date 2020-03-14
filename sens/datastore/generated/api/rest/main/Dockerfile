FROM golang:1.13.8-alpine3.11 as builder
RUN git clone https://source.developers.google.com/p/sensconsole/r/github_senslabs_alpha

RUN sens/datastore/generated/api/rest/main

RUN go build -o datastore

COPY --from=builder datastore /datastore

ENTRYPOINT [ "/datastore" ]