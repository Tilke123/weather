package weather

import "fmt"

var Postaje = map[string]Postaja{}

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(oznaka string, ime string, kraj string) {
	Postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}
func DodajMeritev(oznaka string, temperatura, minTemp, maxTemp float64) {
	if minTemp < temperatura && temperatura < maxTemp {
		if x, ok := Postaje[oznaka]; ok {
			x.meritve = append(x.meritve, temperatura)
			Postaje[oznaka] = x
		} else {
			fmt.Println("Postaja ni veljavna")
		}
	} else {
		fmt.Println("Temperatura ni veljavna")
	}
}
func statistika(oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	if len(Postaje[oznaka].meritve) < stMeritev {
		return 0.0, 0.0, 0.0, false
	}
	var min, max, avg float64 = Postaje[oznaka].meritve[0], Postaje[oznaka].meritve[0], 0
	for _, x := range Postaje[oznaka].meritve {
		avg += x
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}
	}
	return min, max, avg / float64(len(Postaje[oznaka].meritve)), true
}
func IzpisMeritev() {
	for oznaka, x := range Postaje {
		//fmt.Println(oznaka, " - ", x.ime, "(", x.kraj, "): ", x.meritve)
		fmt.Printf("%s  -  %s ( %s ):  %g\n", oznaka, x.ime, x.kraj, x.meritve)
	}
}
func IzpisPorocila(stMeritev int) {
	tmp := " "
	for oznaka, x := range Postaje {
		min, max, avg, ok := statistika(oznaka, stMeritev)
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
