# docker build --no-cache --target bin --output .  .

FROM golang:1.17-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/pricefetch .
FROM scratch AS bin
COPY --from=build /out/* /
