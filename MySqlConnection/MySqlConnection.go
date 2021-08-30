package MySqlConnection

import (
	"database/sql"
	"log"

	"github.com/PirateRoc/GoWebScrapper/Ayuntamiento"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(ayuntamientos []Ayuntamiento.Ayuntamiento) {

	db, err := sql.Open("mysql", "root:password@tcp(gowebscrapper_database_1:3306)/Ayuntamientos")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	stmtIns, err := db.Prepare("INSERT INTO AyuntamientosAragon(Poblacion, Email,Telefono, Web) VALUES( ?, ?, ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for _, ayuntamiento := range ayuntamientos {
		_, err = stmtIns.Exec(ayuntamiento.Poblacion, ayuntamiento.Email, ayuntamiento.Telefono, ayuntamiento.Web) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

}
