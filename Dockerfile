# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /src
RUN export GO111MODULE="on"

COPY . .
RUN ./build.sh

FROM golang:1.17-alpine
WORKDIR /root
COPY --from=build /src/profiler .
COPY --from=build /src/config.json .


EXPOSE 8080

CMD [ "./profiler" ]

