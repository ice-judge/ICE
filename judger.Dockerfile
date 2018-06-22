FROM gcc
WORKDIR /home
ADD bin/scheduler .
CMD ["/home/scheduler"]
