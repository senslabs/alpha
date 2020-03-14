FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine as builder
RUN apk add go
RUN gcloud source repos clone https://source.developers.google.com/p/sensconsole/r/github_senslabs_alpha
RUN cd sens/datastore/generated/api/rest/main
RUN go build -o datastore

COPY --from=builder datastore /datastore
ENTRYPOINT [ "/datastore" ]