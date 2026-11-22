package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Tilke123/weather/weather"

	"github.com/urfave/cli/v3"
)

func main() {
	if err := cmd().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func cmd() *cli.Command {
	return &cli.Command{
		Name:  "weather",
		Usage: "beleženje temperaturnih meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve v °C",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve v °C",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minTemp := c.Float("minTemp")
			maxTemp := c.Float("maxTemp")
			stMeritev := c.Int("stMeritev")

			if minTemp > maxTemp {
				return fmt.Errorf("minTemp (%g) ne sme biti večji od maxTemp (%g)", minTemp, maxTemp)
			}
			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			zazeni(minTemp, maxTemp, stMeritev)
			return nil
		},
	}
}

// zazeni ustvari postaje, doda meritve in izpiše rezultate.
func zazeni(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]weather.Postaja)

	weather.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
	weather.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
	weather.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
	weather.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

	meritve := map[string][]float64{
		"KP01": {24.5, 27, 26.1, 28.4},
		"KR01": {2, -1.5, 3.5},
		"LJ01": {18.2, 21.5, 19.8},
		"MB01": {12.4},
	}
	for oznaka, temperature := range meritve {
		for _, t := range temperature {
			weather.DodajMeritev(postaje, oznaka, t, minTemp, maxTemp)
		}
	}

	// primeri neveljavnih meritev
	weather.DodajMeritev(postaje, "LJ01", 65, minTemp, maxTemp) // previsoka temperatura
	weather.DodajMeritev(postaje, "PT33", 20, minTemp, maxTemp) // postaja ne obstaja

	fmt.Println("\nMeritve:")
	weather.IzpisMeritev(postaje)

	fmt.Println("\nPoročilo:")
	weather.IzpisPorocila(postaje, stMeritev)
}
