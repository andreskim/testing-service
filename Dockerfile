FROM golang:1.15-buster AS build

WORKDIR /go/src/app
ADD . /go/src/app/

RUN make

# Now copy binaries into our base image.
FROM gcr.io/distroless/base 

COPY --from=build /go/src/app/bin/github.com/resideo/testing-service /bin/testing-service

CMD ["/bin/testing-service"]
