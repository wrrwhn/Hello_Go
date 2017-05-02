package server

import (
	"fmt"
	"hello/web/page"
	"hello/web/security"
	"html/template"
	"log"
	"net/http"
)

func CommonHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// common
		title, err := security.CheckPath(w, r)
		if nil != err {
			return
		}

		// func call
		fn(w, r, title)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request, title string) {
	log.Println("index")
	fmt.Fprintf(w, "Hello, %s!", title)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {

	// init
	tmpl := "view"

	// query
	p, err := page.LoadPage(title)
	if nil != err {
		// 404
		http.Redirect(w, r, "../edit/"+title, http.StatusNotFound)
		return
	}

	// template
	log.Printf("view [%v]", title)
	renderTemplate(w, tmpl, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {

	// init
	tmpl := "edit"

	// query
	p, err := page.LoadPage(title)
	if nil != err {
		p = &page.Page{Title: title}
	}

	// template
	log.Printf("edit [%v]", title)
	renderTemplate(w, tmpl, p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// init
	body := r.FormValue("body")

	// save
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect
	log.Printf("save [%v] with [%v]", title, body)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {

	log.Printf("render to %v \n", tmpl)
	t, err := template.ParseFiles(fmt.Sprintf("template/%v.html", tmpl))
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
