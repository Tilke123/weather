// Package weather omogoča beleženje temperaturnih meritev vremenskih postaj
// ter izpis meritev in poročila s statistiko za vsako postajo.
//
// Postaje dodajamo s funkcijo [DodajPostajo], meritve s funkcijo [DodajMeritev],
// rezultate pa izpišeta funkciji [IzpisMeritev] in [IzpisPorocila].
package weather

import "fmt"

// Postaje hrani vse dodane vremenske postaje. Ključ je oznaka postaje.
var Postaje = map[string]Postaja{}

// Postaja predstavlja vremensko postajo z imenom, krajem in seznamom
// izmerjenih temperatur v stopinjah Celzija.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo postajo z oznako oznaka, imenom ime in krajem kraj.
// Nova postaja nima meritev.
func DodajPostajo(oznaka string, ime string, kraj string) {
	Postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev postaji z oznako oznaka doda meritev temperatura, če je ta
// strogo večja od minTemp in strogo manjša od maxTemp.
// Če meritev ni veljavna ali postaja ne obstaja, izpiše sporočilo o napaki.
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

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje z oznako
// oznaka. Če ima postaja manj kot stMeritev meritev, je ok enak false.
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

// IzpisMeritev za vsako postajo izpiše oznako, ime, kraj in vse meritve.
func IzpisMeritev() {
	for oznaka, x := range Postaje {
		//fmt.Println(oznaka, " - ", x.ime, "(", x.kraj, "): ", x.meritve)
		fmt.Printf("%s  -  %s ( %s ):  %g\n", oznaka, x.ime, x.kraj, x.meritve)
	}
}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno
// temperaturo ter opis (Vroče, Toplo, Hladno, Mrzlo) glede na povprečje.
// Za postajo z manj kot stMeritev meritvami izpiše "Premalo meritev".
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
