package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
	"duomly.com/go-bank-backend/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email string
	Password string
}

// Create readBody function
func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	
	return body
}
// Create apiResponse function
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// Refactor login to use readBody
	body := readBody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)
	// Refactor login to use apiResponse function
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Refactor register to use apiResponse function
	apiResponse(register, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func StartApi() {
	router := mux.NewRouter()
	// Add panic handler middleware
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
