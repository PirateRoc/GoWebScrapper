package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Ayuntamiento struct {
	Nombre string
	Email  string
}

func main() {

	//Inicializacion
	c := colly.NewCollector()
	var ayuntamientos []Ayuntamiento
	//Buscamos todos los links que empiecen con /aragon y los visitamos
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/aragon") {
			e.Request.Visit(e.Attr("href"))
		}
	})
	//Dentro de cada link buscamos Nombre y email y los guardamos en nuestro objeto
	c.OnHTML("main", func(e *colly.HTMLElement) {

		nombres := e.ChildTexts("h1[class]")
		var nombre string
		if len(nombres) > 0 {
			nombre = nombres[0]
		} else {
			nombre = ""
		}
		emails := e.ChildTexts("span[itemprop=email]")
		var email string
		if len(emails) > 0 {
			email = emails[0]
		} else {
			email = ""
		}
		if strings.Contains(nombre, ",") {
			posicionComa := strings.Index(nombre, ",")
			nombre = nombre[0:posicionComa]
		}
		ayuntamiento := Ayuntamiento{
			Nombre: nombre,
			Email:  email,
		}

		ayuntamientos = append(ayuntamientos, ayuntamiento)
	})
	//Pagina de inicio
	c.Visit("https://www.todoslosayuntamientos.es/index.php?e56138359dca2145fd1cc148831b88b6=1&option=com_xsbayuntamientos&view=comunidades&task=getAyuntamientos&id_comunidad=2&xsb_elements=1000&xsb_offset=0")
	//Parseamos y sacmos por pantalla la informacion y la longitud
	jsonAyuntamientos, _ := json.Marshal(ayuntamientos)
	fmt.Println(string(jsonAyuntamientos))
	fmt.Println(len(ayuntamientos))

}
