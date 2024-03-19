FROM golang:1.22-bullseye as base

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY . .
RUN make mod

RUN make build

FROM gcr.io/distroless/static-debian12
COPY --from=base /app/dist/dock-ecr-linux /dock-ecr

ENTRYPOINT [ "./dock-ecr" ]