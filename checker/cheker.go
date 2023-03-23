package checker

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"weakpass/models"
)

// function which get post request from main program

var HashGate = func(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	if hash == "" {
		http.Error(w, "Validation Error", http.StatusUnprocessableEntity)
		return
	}

	dbUri := models.GetDbUri() // open DataBase with credentials
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(false)
	defer db.Close() // Database close when function working end

	if err != nil {
		panic("failed to connect database")
	}

	var model = models.Hashes{}
	db.First(&model, models.Hashes{Hash: hash}) // checking hash in DB

	res, err := json.Marshal(model) // json format data
	str := fmt.Sprintf(string(res))

	var data map[string]interface{} // using map for enumeration data
	error := json.Unmarshal([]byte(str), &data)
	if error != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	for key, val := range data { // condition for the presence of hash in the database
		safety := ""
		if key == "Hash" {
			if val == "" {
				safety = "Secure"
			} else {
				safety = "Unsecure"

			}
			ReturnResponse(w, safety)
		}
	}
}

func ReturnResponse(w http.ResponseWriter, safety string) { // function return message response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = safety
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
