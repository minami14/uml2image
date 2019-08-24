FROM ubuntu:18.04

COPY uml2image uml2image

COPY entrypoint.sh entrypoint.sh

COPY plantuml.jar plantuml.jar

COPY uml uml

RUN chmod +x uml2image

RUN chmod +x entrypoint.sh

RUN chmod +x plantuml.jar

RUN apt update

RUN apt upgrade -y

RUN apt install -y default-jre graphviz

ENTRYPOINT ["bash", "entrypoint.sh"]
