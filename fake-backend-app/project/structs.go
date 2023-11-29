package project

type product struct { //ürünler için bir struct oluşturuldu
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float32 `json:"unitPrice"`
}

type category struct { //kategoriler için bir struct oluşturuldu
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}
