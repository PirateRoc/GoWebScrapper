package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PirateRoc/GoWebScrapper/Ayuntamiento"
	"github.com/PirateRoc/GoWebScrapper/MySqlConnection"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//ayuntamientos := AyuntamientosAragon.Get()
	ayuntamientos := []Ayuntamiento.Ayuntamiento{
		{
			Poblacion: "Casa",
			Email:     "aa",
			Telefono:  "aa",
			Web:       "aa",
		},
		{
			Poblacion: "Casa1",
			Email:     "aa",
			Telefono:  "aa",
			Web:       "aa",
		},
	}
	MySqlConnection.Insert(ayuntamientos)
	jsonAyuntamientos, _ := json.Marshal(ayuntamientos)
	// dump results
	body, err := json.Marshal(jsonAyuntamientos)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)

}
func main() {
	addr := ":8080"

	http.HandleFunc("/", handler)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
