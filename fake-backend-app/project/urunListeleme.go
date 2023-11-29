package project

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAllProduct() ([]product, error) { // fonksiyon dizi yapisinda bir struct geri dondurecek
	response, err := http.Get("http://localhost:3000/products/") // get methodu ile url'e istek atip donen degeri response adli bir değişkene atandi
	if err != nil {                                              //hata kontrolu yapildi
		fmt.Println("http.get kisminda hata meydana geldi!!!")
		return nil, err
	}
	defer response.Body.Close() //response degerinin body kismi fonksiyon bitince kessin kapanmasi için defer kullanildi

	bodyByte, err := io.ReadAll(response.Body) //response body byte cinsinden dondurulur io.ReadAll ile bu deger okunup bodyByte diye bir degiskene atandi
	if err != nil {                            //hata kontrolü yapildi
		fmt.Println("io.ReadAll kisminda hata meydana geldi!!!")
		return nil, err
	}

	var urunler []product              // urunler adli bir struct array olusturuldu product structindan
	json.Unmarshal(bodyByte, &urunler) //donen deger json formatinda oldugu için bunu struct tipine aldik
	return urunler, err                // urunler adli struct array dondurulu ve olasi hata icin err donduruldu

}
