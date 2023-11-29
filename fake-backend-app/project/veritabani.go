package project

import (
	"database/sql"
	"fmt"
	"log"
)

func Veritabani() {
	server := "server ismi"
	database := "veritabani ismi"
	username := "kullanici ismi"
	password := 0 //sifre

	db, err := sql.Open("mssql", fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s", server, database, username, password)) // kullanilan veritabanina göre burasi degistirilebilir
	if err != nil {
		fmt.Println("sql.Open da hata olustu")
		return
	}
	defer db.Close()

	datas, _ := GetAllProduct()
	for _, item := range datas {
		_, err := db.Exec("INSERT INTO product (Id, ProductName, CategoryId, UnitPrice) VALUES ($1, $2, $3,$4)", item, datas[1], datas[2], datas[3])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("veritabanina baglandi")

}
