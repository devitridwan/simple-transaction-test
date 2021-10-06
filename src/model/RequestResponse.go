package model

type ReqRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqCreateProduct struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	// Files    *multipart.FileHeader `json:"files"`
}

type ReqListProduct struct {
	Email string `json:"email"`
}

type ReqUpdateProduct struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type ReqDeleteProduct struct {
	Id int `json:"id"`
}

type ReqOrder struct {
	ProductId int `json:"product_id"`
	Amount    int `json:"amount"`
}

type ReqPayment struct {
	Email    string `json:"email"`
	OrderId  string `json:"order_id"`
	TotPrice int    `json:"total_price"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type ResLogin struct {
	Response
	Email string `json:"email"`
	Token string `json:"token"`
}

type DataProductList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Path     string `json:"path"`
}

type ResProductList struct {
	Response
	Data []DataProductList `json:"data"`
}

type ResOrder struct {
	Response
	OrderId string `json:"order_id"`
	Total   int    `json:"total"`
}
