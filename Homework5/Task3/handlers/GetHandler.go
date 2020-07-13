package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var userList []string

// type data represents data, that will appear on website
// (Users - list of users (name: address))
// (Cookies - cookies (name, address)
type data struct{
	Users []string
	Cookies User
}

// function GetHandler handles GET requests.
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// server responds to "/" (404 on any other page)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var dataToTemplate data
	dataToTemplate.Users = userList

	// getting cookie for user name
	if cookieName, err := r.Cookie("name"); err==nil{
		dataToTemplate.Cookies.Name = cookieName.Value
	} else{
		dataToTemplate.Cookies.Name = ""
	}

	// getting cookie for user address
	if cookieAddress, err := r.Cookie("address"); err==nil{
		dataToTemplate.Cookies.Address = cookieAddress.Value
	} else {
		dataToTemplate.Cookies.Address = ""
	}

	templ, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := templ.ExecuteTemplate(w, "index.html", dataToTemplate); err != nil{
		log.Fatal(err)
		return
	}

}