FROM icejudge/golang-node

WORKDIR /go/src/github.com/ice-judge/ICE

ADD . .
RUN set -ex \
	&& make deps-go \
	&& make deps-js
RUN set -ex \
	&& make build-api \
	&& make build-public \
	&& mv ./ice/api/api / \
	&& mv ./ice/public/dist /

CMD ["/api"]
