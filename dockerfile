# dev-build
FROM golang:1.26.0-alpine AS dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

CMD ["air"]

# bygg en docker-image FRÅN denna existerande image, 
# officiell image från docker med go 1.26 installerat byggt på alpine linux-distro.
# är en multi-stage build, detta steg kallar vi för builder, kommer att bygga binären för programmet
FROM golang:1.26.0-alpine AS builder

# kopiera in dependencies och dependency controll från lokalt wdir till composen ./app dir (om den ej finns skapas den nu), hämta dom
# - görs för optimering, docker kör dessa rader som i lager, så när vi laddar hem dependencies separat gör vi det möjligt
# - för docker att använda det lagret om inte någon dependency är uppdaterad, därav behöver vi endast hämta alla dependencies
# - om någon är uppdaterad
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# kompilera go-servern enbart med go-implementationer. Förebygg ev länkning med C-implementationer vilket skulle
# göra binären dynamiskt länkad och därmed crasha servern (dynamiskt länkad är att binären förväntar sig att vissa
# bibliotek finns på operativsystemet)
RUN CGO_ENABLED=0 go build -o app ./cmd/server

# starta en helt ny image med alpine linux, som vi kommer att köra => mindre image
# på denna image ska det endast innehålla det som behövs för att köra applikationen, vilket görs
# i säkerhetssyfte för att minimera attackytan.
FROM alpine:latest
WORKDIR /app

# kopiera binären och frontend från builder-builden till nuvarande build
COPY --from=builder /app/app .
COPY --from=builder /app/web ./web 
COPY --from=builder /app/migrations ./migrations


# containern lyssnar på port 8080
# kör applikationen
CMD ["./app"]