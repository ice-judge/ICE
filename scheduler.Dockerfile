FROM alpine

ADD ./ice/scheduler/scheduler /

CMD ["/scheduler"]
