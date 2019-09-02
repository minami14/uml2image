FROM alpine:latest

RUN apk add openjdk8-jre graphviz ttf-droid

COPY uml2image uml2image

COPY plantuml.jar plantuml.jar

COPY uml uml

COPY entrypoint.sh entrypoint.sh

ENTRYPOINT ["ash", "entrypoint.sh"]