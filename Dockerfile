FROM golang:1.20-bullseye AS build

WORKDIR /build

COPY Makefile ./

COPY ./ ./

RUN make build

FROM debian:bullseye-slim AS final

COPY --from=build /build/unused-file-remover /usr/local/bin/unused-file-remover

ENTRYPOINT ["/usr/local/bin/unused-file-remover"]
