FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add -U tzdata
ENV TZ=Asia/Shanghai
WORKDIR /root/
COPY ./tools .
ENV PATH /root/tools:$PATH
ENTRYPOINT ["./tools"]