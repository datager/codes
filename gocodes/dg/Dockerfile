FROM ubuntu:latest

ARG path
ARG name

COPY $path/* /$name/

WORKDIR /$name

RUN mv $name run

ENTRYPOINT ["./run","-logtostderr=true"]
