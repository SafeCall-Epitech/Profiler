# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /src
RUN export GO111MODULE="on"

COPY . .

RUN apt-get remove -y cmake; \
    wget https://github.com/Kitware/CMake/releases/download/v3.21.1/cmake-3.21.1.tar.gz ; \
    tar xzvf cmake-3.21.1.tar.gz ; \
    cd cmake-3.21.1 ; \
    ./bootstrap  -- -DCMAKE_USE_OPENSSL=OFF ; \
    make -j 8 && make install ; \
    cd .. && rm -rf cmake-3.21.1 cmake-3.21.1.tar.gz

RUN git clone --branch=master https://github.com/zeromq/libzmq.git ~/libzmq
RUN cd ~/libzmq ; \
    git submodule init  ; \
    git submodule update ; \
    mkdir build ; \
    cd build  ; \
    cmake .. ; \
    make -j 12 ; \
    make install

RUN ./build.sh


FROM golang:1.17-alpine
WORKDIR /root
COPY --from=build /src/profiler .
COPY --from=build /src/config.json .


EXPOSE 8080

CMD [ "./profiler" ]
