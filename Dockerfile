FROM scratch

COPY app /

COPY --from=alpine:latest /etc/ssl/certs /etc/ssl/certs

COPY --from=gcr.io/distroless/static:nonroot /etc/passwd /etc/passwd
COPY --from=gcr.io/distroless/static:nonroot /etc/group /etc/group

USER nonroot:nonroot

CMD ["/app"]
