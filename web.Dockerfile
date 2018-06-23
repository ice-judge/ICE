FROM alpine

ADD ./ice/api/api ./ice/public/dist /

CMD ["/api"]
