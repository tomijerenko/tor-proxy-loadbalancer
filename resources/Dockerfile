FROM alpine:latest

RUN apk update && apk add tor
RUN echo "SocksPort 0.0.0.0:9050" > /etc/tor/torrc
RUN chown -R tor /etc/tor

USER tor

EXPOSE 9050

ENTRYPOINT ["tor"]
CMD ["-f", "/etc/tor/torrc"]
