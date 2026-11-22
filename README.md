# weather

Modul v jeziku Go za beleženje temperaturnih meritev vremenskih postaj.

Modul vsebuje:

- paket [`weather`](./weather), ki hrani vremenske postaje in njihove meritve ter izpiše meritve in poročilo s statistiko,
- program `weather` (paket `main`), ki paket uporablja prek ukaznega vmesnika, narejenega s paketom [`github.com/urfave/cli/v3`](https://github.com/urfave/cli).

Dokumentacija paketa: [pkg.go.dev/github.com/Tilke123/weather/weather](https://pkg.go.dev/github.com/Tilke123/weather/weather)

## Namestitev

Namestitev programa:

```bash
go install github.com/Tilke123/weather@v1.0.0
```

Uporaba paketa v lastnem projektu:

```bash
go get github.com/Tilke123/weather@v1.0.0
```

## Uporaba programa

```bash
weather [stikala]
```

| Stikalo | Pomen | Privzeto |
|---|---|---|
| `--minTemp T` | najnižja dovoljena temperatura meritve v °C | `-50` |
| `--maxTemp T` | najvišja dovoljena temperatura meritve v °C | `50` |
| `--stMeritev N` | najmanjše število meritev, potrebnih za izračun statistike | `3` |
| `--help` | izpiše pomoč | |

Meritev je veljavna, če je strogo večja od `minTemp` in strogo manjša od `maxTemp`.

Primeri:

```bash
weather                                # privzete nastavitve
weather --stMeritev 4                  # statistika samo za postaje z vsaj 4 meritvami
weather --minTemp 0 --maxTemp 25       # strožje meje za veljavne meritve
weather --help
```

Iz izvorne kode program poženete z `go run .` v korenski mapi repozitorija.

Program doda štiri postaje (Koper, Kredarica, Ljubljana, Maribor) in nekaj meritev, preizkusi dodajanje neveljavne meritve in meritve za neobstoječo postajo, nato izpiše vse meritve in poročilo.

## Uporaba paketa `weather`

```go
package main

import "github.com/Tilke123/weather/weather"

func main() {
	weather.DodajPostajo("LJ01", "Ljubljana Bežigrad", "Ljubljana")

	weather.DodajMeritev("LJ01", 18.2, -50, 50)
	weather.DodajMeritev("LJ01", 21.5, -50, 50)
	weather.DodajMeritev("LJ01", 19.8, -50, 50)

	weather.IzpisMeritev()
	weather.IzpisPorocila(3)
}
```

Izvožene funkcije:

| Funkcija | Opis |
|---|---|
| `DodajPostajo(oznaka, ime, kraj)` | doda novo postajo brez meritev |
| `DodajMeritev(oznaka, temperatura, minTemp, maxTemp)` | postaji doda meritev, če je veljavna |
| `IzpisMeritev()` | za vsako postajo izpiše oznako, ime, kraj in meritve |
| `IzpisPorocila(stMeritev)` | za vsako postajo izpiše najnižjo, najvišjo in povprečno temperaturo ter opis |

Opis v poročilu je določen glede na povprečno temperaturo:

| Povprečje | Opis |
|---|---|
| vsaj 25 °C | Vroče |
| vsaj 15 °C | Toplo |
| vsaj 5 °C | Hladno |
| pod 5 °C | Mrzlo |

## Struktura repozitorija

```
weather/
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── main.go          program (paket main)
└── weather/
    └── weather.go   paket weather
```

## Licenca

Projekt je objavljen pod licenco, navedeno v datoteki [LICENSE](./LICENSE).