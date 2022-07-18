ARG build_image
ARG base_image

# Build stage.
FROM ${build_image} AS build
ARG branch=main
WORKDIR /build
COPY . .
RUN go build .

# Copy binary and default config to slim image.
FROM ${base_image}
COPY --from=build /build/verboserver /usr/local/bin/
ENTRYPOINT ["verboserver"]
