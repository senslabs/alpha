FROM golang:1.14.0-alpine3.11 as builder
RUN apk add --no-cache git
RUN git clone https://source.cloud.google.com/sensconsole/github_senslabs_alpha
WORKDIR $WORKDIR/sens/datastore/generated/api/rest/main
RUN CGO_ENABLED=0 GOOS=linux RUN go build -o datastore 

COPY --from=builder datastore /datastore
ENTRYPOINT [ "/datastore" ]