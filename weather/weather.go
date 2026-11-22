package weather

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	if -50 < temperatura && temperatura < 50 {
		if x, ok := postaje[oznaka]; ok {
			x.meritve = append(x.meritve, temperatura)
			postaje[oznaka] = x
		} else {
			fmt.Println("Postaja ni veljavna")
		}
	} else {
		fmt.Println("Temperatura ni veljavna")
	}
}
func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	if len(postaje[oznaka].meritve) < 3 {
		return 0.0, 0.0, 0.0, false
	}
	var min, max, avg float64 = postaje[oznaka].meritve[0], postaje[oznaka].meritve[0], 0
	for _, x := range postaje[oznaka].meritve {
		avg += x
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}
	}
	return min, max, avg / float64(len(postaje[oznaka].meritve)), true
}
func izpisMeritev(postaje map[string]Postaja) {
	for oznaka, x := range postaje {
		//fmt.Println(oznaka, " - ", x.ime, "(", x.kraj, "): ", x.meritve)
		fmt.Printf("%s  -  %s ( %s ):  %g\n", oznaka, x.ime, x.kraj, x.meritve)
	}
}
func izpisPorocila(postaje map[string]Postaja) {
	tmp := " "
	for oznaka, x := range postaje {
		min, max, avg, ok := statistika(postaje, oznaka)
		if !ok {
			fmt.Println("Premalo meritev")
		} else {
			switch {
			case avg >= 25:
				tmp = "Vroče"
			case avg >= 15 && avg <= 25:
				tmp = "Toplo"
			case avg >= 5 && avg <= 15:
				tmp = "Hladno"
			case avg < 5:
				tmp = "Mrzlo"
			default:
				tmp = "X"
			}
			fmt.Printf("%s, : min , %g, , max , %g, , povprečje , %g, ->, %s\n", x.ime, min, max, avg, tmp)
		}
	}
}
