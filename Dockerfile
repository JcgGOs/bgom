FROM  golang:1.18-alpine as builder
RUN apk update && apk add -U --no-cache git

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux go build -o /app/bgom


FROM alpine:3.15

# RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories
# RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositorie



RUN apk update  \
    && apk add --no-cache tzdata openjdk11 font-noto-cjk graphviz \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY plantuml /usr/bin/
RUN chmod 755 /usr/bin/plantuml

RUN apk add curl

ARG PLANTUML_VERSION="1.2022.4"
RUN mkdir /usr/share/plantuml \
    && curl -L https://github.com/plantuml/plantuml/releases/download/v${PLANTUML_VERSION}/plantuml-${PLANTUML_VERSION}.jar > /usr/share/plantuml/plantuml.jar

WORKDIR /app
RUN mkdir /app/static/ /app/templates/ /app/posts/
COPY --from=builder /app/bgom       /app
COPY --from=builder /app/static/    /app/static/
COPY --from=builder /app/templates/ /app/templates/
COPY --from=builder /app/posts/     /app/posts/

CMD [ "/app/bgom"]