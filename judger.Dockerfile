FROM alpine

ADD ./ice/judger/out/judger /

CMD ["/judger"]
