FROM gcc

WORKDIR /ice-judge

ADD . .
RUN make build-judger \
	&& mv ./ice/judger/out/judger /

CMD ["/judger"]
