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
	cmd := &cli.Command{
		Name:  "weather",
		Usage: "weather beleži temperature postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "najnižja dovoljena temperatura",
				Value: -50,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "najvišja dovoljena temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "najmanjše število meritev",
				Value: 3,
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

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func zazeni(minTemp, maxTemp float64, stMeritev int) {

	weather.DodajPostajo("KP01", "Koper Luka", "Koper")
	weather.DodajPostajo("KR01", "Kredarica", "Triglav")
	weather.DodajPostajo("LJ01", "Ljubljana Bežigrad", "Ljubljana")
	weather.DodajPostajo("MB01", "Maribor Tabor", "Maribor")

	meritve := map[string][]float64{
		"KP01": {24.5, 27, 26.1, 28.4},
		"KR01": {2, -1.5, 3.5},
		"LJ01": {18.2, 21.5, 19.8},
		"MB01": {12.4},
	}
	for oznaka, temperature := range meritve {
		for _, temp := range temperature {
			weather.DodajMeritev(oznaka, temp, minTemp, maxTemp)
		}
	}

	weather.DodajMeritev("LJ01", 65, minTemp, maxTemp)
	weather.DodajMeritev("PT33", 20, minTemp, maxTemp)
	weather.DodajMeritev("MB01", 20, minTemp, maxTemp)
	weather.DodajMeritev("MB01", 25, minTemp, maxTemp)

	fmt.Println()
	fmt.Println("Meritve:")
	weather.IzpisMeritev()

	fmt.Println()
	fmt.Println("Poročilo:")
	weather.IzpisPorocila(stMeritev)
}
