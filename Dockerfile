FROM alpine:3.8
RUN adduser -D -u 1000 k2w


FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
USER 1000
ADD /output/k2w /k2w
ENTRYPOINT [ "/k2w" ]