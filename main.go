package main

import (
	"encoding/json"
	"fmt"

	"github.com/PirateRoc/GoWebScrapper/AyuntamientosAragon"
)

func main() {

	ayuntamientos := AyuntamientosAragon.Get()
	jsonAyuntamientos, _ := json.Marshal(ayuntamientos)
	fmt.Println(string(jsonAyuntamientos))
	fmt.Println(len(ayuntamientos))

}
