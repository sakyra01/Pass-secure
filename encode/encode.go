package encode

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"weakpass/models"
)

func Encryption(word string) {
	h := sha512.New()
	h.Write([]byte(word))
	bs := h.Sum(nil)

	k := fmt.Sprintf("%x", bs)
	DbUpdate(k)
}

func Enumeration() {
	f, err := os.Open("weakpass.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		word := reader.Text()
		Encryption(word)
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

}

func DbUpdate(k string) {
	dbUri := models.GetDbUri()
	db, err := gorm.Open("postgres", dbUri)
	db.LogMode(false)
	if err != nil {
		panic("failed to connect database")
	}

	result := models.Hashes{}
	db.FirstOrCreate(&result, models.Hashes{Hash: k})
	db.Save(result)

	db.Close()
}
