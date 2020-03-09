FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./bank .
ENV PATH /root/bank:$PATH
#ADD ./assets /root/assets
ENTRYPOINT ["./bank"]