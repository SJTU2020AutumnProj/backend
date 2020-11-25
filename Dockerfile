FROM golang
COPY --from=build /build/api-auth .
COPY --from=build /config.json .

CMD ["./api-auth","--registry=consul"]