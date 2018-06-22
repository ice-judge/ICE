FROM alpine
WORKDIR /home
ADD bin/api-amd64 .
CMD ["/home/api-amd64"]
