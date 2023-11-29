package project

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func UrunEkleme() (string, error) { // fonksiyon bir string deger ve bir hata degeri donduruyor
	urun := product{ProductName: "mikrofon", CategoryId: 1, UnitPrice: 250.00} // urunler(product) srtuctina bi veri eklendi
	jsonProduct, err := json.Marshal(urun)                                     // urun adli struct tipteki data json formatina çevrildi
	if err != nil {
		return "json.Marshal hatasi olustu..!", err // fonksiyon bir string ve bir err dondurdugu icin return yapisina dikkat...
	}
	reponse, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct)) //response isimli bir degişkenle bir post istegi olusturdum
	if err != nil {
		return "http.post hatasi olustu..!", err
	}
	defer reponse.Body.Close() //http.post istegi bize bir deger dondurur bu degerin body kismini UrunEkleme fonksiyonum bitince kessin kapanacak sekilde bir defer olusturdum

	return "kaydedildi", err //urunun kaydedildiği mesaji donduruldu ve hata olması durumunda hata mesaji donduruldu

}
