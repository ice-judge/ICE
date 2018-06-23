FROM gcc

WORKDIR /ice-judge

ADD . .
RUN set -ex \
	&& make build-judger \
	&& mv ./ice/judger/out/judger /

CMD ["/judger"]
