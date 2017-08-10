
# 1st Stage - build binary

FROM fedora:26

RUN yum install -y golang make

ENV GOPATH="/go/"

WORKDIR $GOPATH/src/github.com/kadel/kedge-demo/first/backend
COPY . .

RUN make build



# 2th Stage - create image

FROM fedora:26

COPY --from=0 /go/src/github.com/kadel/kedge-demo/first/backend/server /opt/api/

EXPOSE 3000

ENV HOST 0.0.0.0
ENV DATA_FILE /opt/api/data/comments.json

RUN mkdir -p /opt/api/data/
RUN chmod a+w  /opt/api/data/

ENTRYPOINT ["/opt/api/server"]
