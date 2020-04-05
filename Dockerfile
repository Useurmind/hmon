FROM golang as build
WORKDIR /build
COPY . .
RUN go build -o hmon ./app
RUN chmod +x ./hmon

FROM debian as run
WORKDIR /app
COPY --from=build /build/hmon .
ENTRYPOINT [ "./hmon" ]
