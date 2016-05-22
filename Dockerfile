FROM busybox
COPY ca-certificates.crt /etc/ssl/certs/
COPY app /
RUN mkdir public
EXPOSE 5000
CMD ["/app"]
