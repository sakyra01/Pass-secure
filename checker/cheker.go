package checker

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"weakpass/models"
)

//func PassControl() {
//	arg := os.Args[1:]
//	if arg[0] == "-p" || arg[0] == "--pass" {
//		pass := arg[1]
//		h := sha512.New()
//		h.Write([]byte(pass))
//		bs := h.Sum(nil)
//
//		hash := fmt.Sprintf("%x\n", bs)
//		DataBaseGates(hash)
//	}
//}
//
//func DataBaseGates(hash string) {
//	dbUri := models.GetDbUri()
//	db, err := gorm.Open("postgres", dbUri)
//	db.LogMode(false)
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	result := models.Hashes{}
//	db.First(&result, models.Hashes{Hash: hash})
//
//	res, _ := json.Marshal(result)
//	str := fmt.Sprintf(string(res))
//
//	var data map[string]interface{}
//	error := json.Unmarshal([]byte(str), &data)
//	if error != nil {
//		fmt.Printf("could not unmarshal json: %s\n", err)
//		return
//	}
//
//	for key, val := range data {
//		if key == "Hash" {
//			if val == "" {
//			} else {
//				db.Close()
//				fmt.Println("Password unsecure")
//			}
//		}
//	}
//	db.Close()
//}

var HashGate = func(w http.ResponseWriter, r *http.Request) {
	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(true)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	hash := r.FormValue("hash")

	var model = models.Hashes{}
	db.First(&model, models.Hashes{Hash: hash})

	res, err := json.Marshal(model)
	str := fmt.Sprintf(string(res))

	var data map[string]interface{}
	error := json.Unmarshal([]byte(str), &data)
	if error != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Println(data)

	for key, val := range data {
		if key == "Hash" {
			if val == "" {
				fmt.Fprintf(w, "Secure")
			} else {
				fmt.Fprintf(w, "Unsecure")
			}
		}
	}
}
