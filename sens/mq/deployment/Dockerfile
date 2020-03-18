FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine as builder
RUN git config --system credential.helper gcloud.sh
RUN git clone https://source.developers.google.com/p/sensconsole/r/github_senslabs_alpha -b mq /github_senslabs_alpha
RUN apk add go

WORKDIR /github_senslabs_alpha/sens/mq/datastore
RUN go build -o datastoreconsumer

FROM gcr.io/google.com/cloudsdktool/cloud-sdk:alpine
COPY --from=builder /github_senslabs_alpha/sens/datastore/generated/api/rest/main/datastore /datastore
ENTRYPOINT [ "/datastore" ]