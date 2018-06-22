# https://github.com/sol-prog/Clang-in-Docker/blob/master/Dockerfile
FROM ubuntu:16.04

RUN apt update && apt install -y \
  xz-utils \
  build-essential \
  curl \
  && rm -rf /var/lib/apt/lists/* \
  && curl -SL http://releases.llvm.org/6.0.0/clang+llvm-6.0.0-x86_64-linux-gnu-ubuntu-16.04.tar.xz \
  | tar -xJC . && \
  mv clang+llvm-6.0.0-x86_64-linux-gnu-ubuntu-16.04 clang_6.0.0 && \
  echo 'export PATH=/clang_6.0.0/bin:$PATH' >> ~/.bashrc && \
  echo 'export LD_LIBRARY_PATH=/clang_6.0.0/lib:LD_LIBRARY_PATH' >> ~/.bashrc

CMD [ "/bin/bash" ]

