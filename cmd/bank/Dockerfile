FROM alpine:latest
#RUN apk --no-cache add ca-certificates
RUN apk add -U tzdata
ENV TZ=Asia/Shanghai
WORKDIR /root/
COPY ./bank .
ENV PATH /root/bank:$PATH
#ADD ./assets /root/assets
ENTRYPOINT ["./bank"]