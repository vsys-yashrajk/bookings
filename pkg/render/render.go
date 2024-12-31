package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"github.com/vsys-yashrajk/bookings/pkg/config"
	"github.com/vsys-yashrajk/bookings/pkg/models"
	"net/http"
	"path/filepath"
)

// var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tempCache map[string]*template.Template

	if app.UseCache {
		tempCache = app.TemplateCache
	} else {
		tempCache, _ = CreateTemplateCache()
	}

	t, ok := tempCache[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of the files name *.page.html
	htmlPages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//iterate over the slice htmlPages
	for _, page := range htmlPages {

		//gives the name of the file excluding path
		fileName := filepath.Base(page)

		//parse the file
		templateSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//get all of the files name *.layout.html
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		//adds the base layout to the pages
		if len(matches) > 0 {

			templateSet, err = templateSet.ParseGlob("./templates/*layout.html")
			if err != nil {
				return myCache, err
			}
		}

		//adds the page in the cache
		myCache[fileName] = templateSet
	}

	return myCache, nil
}

//Advanced render template
//var templateCache = make(map[string]*template.Template)
// func RenderTemplate(w http.ResponseWriter,tname string){
// 	var tmpl *template.Template
// 	var err error

// 	_,inMap := templateCache[tname]

// 	if !inMap{
// 		//need to cache new template
// 		log.Println("creating and adding to cache")
// 		err = createTemplateCache(tname)
// 		if err !=  nil{
// 			log.Println(err)
// 		}
// 	} else{
// 		//already in chache
// 		log.Println("Using cache template")
// 	}
// 	tmpl = templateCache[tname]
// 	err = tmpl.Execute(w,nil)
// 	if err != nil{
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(tname string) error{
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s",tname),
// 		"./templates/base.layout.html"
// 	}
// 	tmpl,err := template.ParseFiles(templates...)
// 	if err != nil{
// 		return err
// 	}
// 	templateCache[tname] = tmpl
// 	return nil
// }

//Basic render template
// func RenderTemplate(w http.ResponseWriter, tmpl string){
// 	parsedTemplate,_ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.html")
// 	err := parsedTemplate.Execute(w,nil)
// 	if err != nil{
// 		fmt.Println("error parsing template",err)
// 		return
// 	}
// }
