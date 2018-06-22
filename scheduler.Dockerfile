FROM alpine
WORKDIR /home
ADD bin/scheduler-amd64 .
CMD ["/home/scheduler-amd64"]
