FROM scratch
ADD bin/main.static /main
CMD ["/main"]
