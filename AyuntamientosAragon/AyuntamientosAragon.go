package AyuntamientosAragon

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

// Modelo con la informacion de un ayuntamiento
type Ayuntamiento struct {
	Poblacion string
	Email     string
	Telefono  string
	Web       string
}

func Get() []Ayuntamiento {

	//Inicializacion
	c := colly.NewCollector(
		colly.Async(),
	)
	var ayuntamientos []Ayuntamiento
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 16})
	//Buscamos todos los links que empiecen con /aragon y los visitamos
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/aragon") {
			e.Request.Visit(e.Attr("href"))
		}
	})
	//Dentro de cada link buscamos Nombre y email y los guardamos en nuestro objeto
	c.OnHTML("main", func(e *colly.HTMLElement) {

		poblacionArray := e.ChildTexts("h1[class]")
		var poblacion string
		if len(poblacionArray) > 0 {
			poblacion = poblacionArray[0]
		} else {
			poblacion = ""
		}
		if strings.Contains(poblacion, ",") {
			posicionComa := strings.Index(poblacion, ",")
			poblacion = poblacion[0:posicionComa]
		}

		emails := e.ChildTexts("span[itemprop=email]")
		var email string
		if len(emails) > 0 {
			email = emails[0]
		} else {
			email = ""
		}
		telefonos := e.ChildTexts("span[itemprop=telephone]")

		var telefono string
		if len(telefonos) > 0 {
			telefono = telefonos[0]
		} else {
			telefono = ""
		}
		webs := e.ChildTexts("a[itemprop=url]")

		var web string
		if len(webs) > 0 {
			web = webs[0]
		} else {
			web = ""
		}

		ayuntamiento := Ayuntamiento{
			Poblacion: poblacion,
			Email:     email,
			Telefono:  telefono,
			Web:       web,
		}

		ayuntamientos = append(ayuntamientos, ayuntamiento)
	})
	//Pagina de inicio
	c.Visit("https://www.todoslosayuntamientos.es/index.php?bbaf263a157d2b4561bd2ad296554729=1&option=com_xsbayuntamientos&view=comunidades&task=getAyuntamientos&id_comunidad=2&xsb_elements=1000&xsb_offset=0")

	c.Wait()
	return ayuntamientos

}
