# GoWebScrapper
WebScrapper  hecho en Go


Para hacer el scrapper de los ayuntamientos de aragon visitar la pagina
 https://www.todoslosayuntamientos.es/aragon
  darle a consiltar mas y ver la llamada que hace, sera de este estilo
https://www.todoslosayuntamientos.es/index.php?bbaf263a157d2b4561bd2ad296554729=1&option=com_xsbayuntamientos&view=comunidades&task=getAyuntamientos&format=json&id_comunidad=2&xsb_elements=10&xsb_offset=20

Quitamos el offsett y el formato y aumentamos el tamaño a 1000

https://www.todoslosayuntamientos.es/index.php?bbaf263a157d2b4561bd2ad296554729=1&option=com_xsbayuntamientos&view=comunidades&task=getAyuntamientos&id_comunidad=2&xsb_elements=1000&xsb_offset=0



Separados los proyectos de mysql y el de go y creado docker compose y dockerfiles separados
Pasos a seguir  crear mas metodos, CRUD y añadir swagger
