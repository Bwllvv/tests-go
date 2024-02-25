package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type user struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func main() {
	http.HandleFunc("/bro", bro_page)
	http.HandleFunc("/pages", home_page)
	http.ListenAndServe(":8080", nil)
}

//func (u user) GetAllInfo() string {
//	return
//}

func home_page(w http.ResponseWriter, r *http.Request) {
	lev := &user{"lev", 19, 24000, 1.43, 1.1}
	fmt.Fprintf(w, "username:%s\nold:%d\nbalance:%d dollars\ngrades:%f\nhapiness:%f", lev.Name, lev.Age, lev.Money, lev.Avg_grades, lev.Happiness)
}

func bro_page(w http.ResponseWriter, r *http.Request) {
	lev := user{"lev", 19, 24000, 1.43, 1.1}
	templ, err := template.ParseFiles("templates/Site.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	templ.Execute(w, lev)
}
