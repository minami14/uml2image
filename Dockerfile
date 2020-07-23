  
FROM golang AS build

WORKDIR $GOPATH/src/github.com/minami14/uml2image

COPY . .

RUN CGO_ENABLED=0 go build -o uml2image cmd/main.go


FROM alpine:latest

RUN apk add openjdk8-jre graphviz ttf-droid

COPY --from=build uml2image uml2image

COPY plantuml.jar plantuml.jar

COPY uml uml

COPY entrypoint.sh entrypoint.sh

ENTRYPOINT ["ash", "entrypoint.sh"]
