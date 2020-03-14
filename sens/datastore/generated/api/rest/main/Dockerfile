FROM gcr.io/cloud-builders/gcloud as builder
RUN git config --system credential.helper gcloud.sh
RUN gcloud config set account developer@dozee.io
RUN gcloud source repos clone https://source.developers.google.com/p/sensconsole/r/github_senslabs_alpha

RUN apk add go
RUN cd sens/datastore/generated/api/rest/main
RUN go build -o datastore

COPY --from=builder datastore /datastore
ENTRYPOINT [ "/datastore" ]