package weather

import "fmt"

var Postaje = map[string]Postaja{}

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(oznaka string, ime string, kraj string, meritve []float64) {
	Postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	if minTemp < temperatura && temperatura < maxTemp {
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
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	if len(postaje[oznaka].meritve) < stMeritev {
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
func IzpisMeritev(postaje map[string]Postaja) {
	for oznaka, x := range postaje {
		//fmt.Println(oznaka, " - ", x.ime, "(", x.kraj, "): ", x.meritve)
		fmt.Printf("%s  -  %s ( %s ):  %g\n", oznaka, x.ime, x.kraj, x.meritve)
	}
}
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	tmp := " "
	for oznaka, x := range postaje {
		min, max, avg, ok := statistika(postaje, oznaka, stMeritev)
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
