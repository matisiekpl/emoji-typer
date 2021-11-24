FROM golang AS build
# utworz katalog roboczy
RUN mkdir -p /go/src/emoji_typer
# ustaw go jako aktualny katalog roboczy
WORKDIR /go/src/emoji_typer
# skopiuj kod do niego
COPY . .
# zainstaluj program pakujacy pliki statyczne Packr2
RUN go install github.com/gobuffalo/packr/v2/packr2@v2.8.3
# niech Packr2 spakuje pliki statyczne
RUN packr2
# skompiluj program
RUN go build -ldflags "-linkmode external -extldflags -static" emoji_typer

# utworz nowy kontener, zawierajacy jedynie skompilowany plik z kodem maszynowym
FROM scratch
# skopiuj plik wykonywalny
COPY --from=build /go/src/emoji_typer/emoji_typer /emoji_typer
# powiedz Dockerowi, by podczas uruchamiania kontenera, wywolal ten plik
CMD ["/emoji_typer"]
