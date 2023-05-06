FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /fxchain
# Add source files
COPY . .

# build fxchain
RUN make build-linux

# Final image
FROM golang:1.17 as final

WORKDIR /fxchaind
# Copy over binaries from the build-env
COPY --from=build-env /fxchain/build/fxchaind /usr/bin/fxchaind
COPY --from=build-env /fxchain/build/fxchaincli /usr/bin/fxchaincli
COPY --from=build-env /fxchain/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

