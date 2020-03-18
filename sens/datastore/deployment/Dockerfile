FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine as builder
RUN git config --system credential.helper gcloud.sh
RUN git clone https://source.developers.google.com/p/sensconsole/r/github_senslabs_alpha /github_senslabs_alpha
RUN apk add go

WORKDIR /github_senslabs_alpha/sens/datastore/generated/main
RUN go build -o datastore

FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine
COPY --from=builder /github_senslabs_alpha/sens/datastore/generated/main/datastore /datastore
ENTRYPOINT [ "/datastore" ]
