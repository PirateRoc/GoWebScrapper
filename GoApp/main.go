package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PirateRoc/GoWebScrapper/GoApp/AyuntamientosAragon"
	"github.com/PirateRoc/GoWebScrapper/GoApp/MySqlConnection"
)

func handler(w http.ResponseWriter, r *http.Request) {

	ayuntamientos := AyuntamientosAragon.Get()

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
