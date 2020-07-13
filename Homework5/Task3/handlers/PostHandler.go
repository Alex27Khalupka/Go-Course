package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// type User represents information about user
// (Name - users name, Address - users address)
type User struct{
	Name string
	Address string
}

// function PostHandler handles POST requests
func PostHandler(w http.ResponseWriter, r *http.Request){
	// server responds to "/" (404 on any other page)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := r.ParseForm(); err != nil{
		log.Fatal(err)
		return
	}

	// getting user name and user address from forms
	user := User{r.PostForm.Get("name"), r.PostForm.Get("address")}
	userList = append(userList, fmt.Sprintf("%s: %s", user.Name, user.Address))

	// creating cookie for user name
	cookieName := http.Cookie{
		Name: "name",
		Value: user.Name,
	}
	http.SetCookie(w, &cookieName)

	// creating cookie for user address
	cookieAddress := http.Cookie{
		Name: "address",
		Value: user.Address,
	}
	http.SetCookie(w, &cookieAddress)

	// redirecting user to default page after submitting
	http.Redirect(w, r,"/", 302)

}
