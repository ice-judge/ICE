FROM ksunhokim/golang

WORKDIR /go/src/github.com/ice-judge/ICE

ADD . .
RUN set -ex \
	&& make deps-go
RUN set -ex \
	&& make build-scheduler \
	&& mv ./ice/scheduler/scheduler /

CMD ["/scheduler"]
