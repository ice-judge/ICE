FROM node
WORKDIR /home
ADD public .
CMD ["node start --prefix /home/publc"]
