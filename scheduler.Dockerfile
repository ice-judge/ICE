FROM ksunhokim/golang

WORKDIR /go/src/github.com/ice-judge/ICE

ADD . .
RUN make deps-go
RUN make build-scheduler \
	&& mv ./ice/scheduler/scheduler /

CMD ["/scheduler"]
