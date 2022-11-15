package main

import "fmt"

// const (
// 	OCAK = iota
// 	SUBAT
// 	MART
// 	NISAN

// )

type Insan struct {
	boy int
}

type Hayvan struct {
	boy int
}

type Canli interface {
	Boy() int
}

func BanaBoyuGoster(c Canli)int {
	return c.Boy()
}

func (i *Insan) Boy() int {
	return i.boy
}

func (h *Hayvan) Boy() int {
	return h.boy
}

func main() {

	insan := Insan {
		boy: 180,
	}

	hayvan := Hayvan {
		boy: 60,
	}

	insanBoyu := BanaBoyuGoster(&insan)
	hayvanBoyu := BanaBoyuGoster(&hayvan)

	fmt.Println(insanBoyu)
	fmt.Println(hayvanBoyu)
	// teachers := []string{}

	// teachers = append(teachers, "bekir")

	// haritalar := map[string]int{}
	
	// for _,value := range teachers {
	// 	fmt.Println(value)
	// }

	// haritalar["izmir"] = 35
	// haritalar["bolu"] = 14
	// for key, value := range haritalar {
	// 	fmt.Printf("key: %s, value: %d", key, value)
	// }

	// if _, ok := haritalar["elaz"];ok {
	// 	fmt.Print("izmir guzel bir yer")
	// } else {
	// 	fmt.Print("kotu bir yer")
	// }

	// switch input {
	// case OCAK:
	// 	fmt.Println("Kış kapıdan baktırır")
	// case SUBAT:
	// 	fmt.Println("kazma kurek yaktirir")
	// }

}