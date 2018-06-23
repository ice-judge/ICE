FROM icejudge/golang-node

WORKDIR /go/src/github.com/ice-judge/ICE

ADD . .

RUN make deps-go && \
	make deps-js

RUN make build-api \
	&& make build-public \
	&& mv ./ice/api/api / \
	&& mv ./ice/public/dist /

CMD ["/api"]
