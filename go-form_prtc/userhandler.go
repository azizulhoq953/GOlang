package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Person is a struct decsribing its properties
type Person struct {
	Username      string `json:"username"`
	Password   string `json:"password"`
	 
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve people from postgresql database using our `store` interface variable's
	// `func (*dbstore) GetPerson` pointer receiver method defined in `store.go` file
	personList, err := storeInfo.GetPerson()


	// Convert the `personList` variable to JSON
	personListBytes, err := json.Marshal(personList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON list of persons to response
	w.Write(personListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data received in the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract the field information about the person from the form info
	person := Person{}
	person.Username = r.Form.Get("username")
	person.Password = r.Form.Get("password")
	

	// Write new person details into postgresql database using our `store` interface variable's
	// `func (*dbstore) CreatePerson` pointer receiver method defined in `store.go` file
	err = storeInfo.CreatePerson(&person)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
	}	//Redirect to the originating HTML page