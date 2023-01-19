# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /src
RUN export GO111MODULE="on"

COPY . .

RUN ./build.sh
RUN ls

FROM golang:1.17-alpine
WORKDIR /root
COPY --from=build /src/profiler .
COPY --from=build /src/config.json .
COPY --from=build /libzmq.so .
COPY --from=build libzmq.so.5 .
COPY --from=build libzmq.so.5.2.4 .


EXPOSE 8080

CMD [ "./profiler" ]
