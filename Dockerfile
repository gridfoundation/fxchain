# Simple usage with a mounted data directory:
# > docker build -t fxchain .
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fxchaind:/root/.fxchaind -v ~/.fxchaincli:/root/.fxchaincli fxchain fxchaind init mynode
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.fxchaind:/root/.fxchaind -v ~/.fxchaincli:/root/.fxchaincli fxchain fxchaind start
FROM golang:1.17.2-alpine AS build-env

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

# Set working directory for the build
WORKDIR /go/src/github.com/gridfx/fxchain

# Add source files
COPY . .

ENV GO111MODULE=on \
    GOPROXY=http://goproxy.cn
# Build GRIDFxChain
RUN make install

# Final image
FROM alpine:edge

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/fxchaind /usr/bin/fxchaind
COPY --from=build-env /go/bin/fxchaincli /usr/bin/fxchaincli

# Run fxchaind by default, omit entrypoint to ease using container with fxchaincli
CMD ["fxchaind"]
