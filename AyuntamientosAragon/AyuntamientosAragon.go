package AyuntamientosAragon

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Modelo con la informacion de un ayuntamiento
type Ayuntamiento struct {
	Nombre string
	Email  string
}

func GetAyuntamientosAragon() []Ayuntamiento {

	//Inicializacion
	c := colly.NewCollector()
	var ayuntamientos []Ayuntamiento
	//Buscamos todos los links que empiecen con /aragon y los visitamos
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
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
	c.Visit("https://www.todoslosayuntamientos.es/index.php?bbaf263a157d2b4561bd2ad296554729=1&option=com_xsbayuntamientos&view=comunidades&task=getAyuntamientos&id_comunidad=2&xsb_elements=1000&xsb_offset=0")

	return ayuntamientos

}
