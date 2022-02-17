package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// DB set up
func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s port=%s dbname=%s sslmode=disable", "root", "password", 5432, "postgresql")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
	  panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
  
	
}

type JsonResponse struct {
	Type    string        `json:"type"`
	Data    []AccountData `json:"data"`
	Message string        `json:"message"`
}
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

//  errors handling
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	// Initiate the mux router
	r := mux.NewRouter()

	// Route handles & endpoints

	// Get all account
	r.HandleFunc("/account/", GetAccount).Methods("GET")

	// Create a account
	r.HandleFunc("/account/", CreateAccount).Methods("POST")

	// Delete a account
	r.HandleFunc("/account/{account_ID}", DeleteAccount).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", r))
}


// Get all account

func GetAccount(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()

	// Get all account from account struct table
	rows, err := db.Query("SELECT * FROM Account")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var Account []AccountData

	// Foreach AccountData
	for rows.Next() {
		var Attributes *AccountAttributes
		var ID string
		var OrganisationID string
		var Type string
		var Version *int64
		err = rows.Scan(&ID, &Type, &OrganisationID, &Version)

		// check errors
		checkErr(err)

		Account = append(Account, AccountData{Attributes: attributes, ID: id, OrganisationID: organisation_id, Type: type, Version: version})
	}

	var response = JsonResponse{Type: "success", Data: Account}

	json.NewEncoder(w).Encode(response)
}

// function Create a account

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	Attributes := r.FormValue("attributes")
	OrganisationID := r.FormValue("organisation_id")
	Type := r.FormValue("type")
	Version := r.FormValue("version")

	var response = JsonResponse{}

	if organisation_id == "" || version == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a parameter."}
	} else {
		db := ConnectDB()

		fmt.Println("Inserting new account: " + organisation_id + " and name: " + version)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO Account(organisation_id, type) VALUES($1, $2) returning id;", organisation_id, version).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The account has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a account

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	organisation_id := params["organisationid"]

	var response = JsonResponse{}

	if organisation_id == "" {
		response = JsonResponse{Type: "error", Message: "You are missing organisationid parameter."}
	} else {
		db := setupDB()

		_, err := db.Exec("DELETE FROM Account where organisation_id = $1", organisation_id)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The account has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
