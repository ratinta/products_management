package models

import "github.com/dgrijalva/jwt-go"

//Products detail
type Products struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Exp      string `json:"expire_date"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Price    int    `json:"price"`
}

type Body struct {
	Name     string `json:"name"`
	Exp      string `json:"expire_date"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Price    int    `json:"price"`
}

type Category struct {
	ID   string
	Name string
}

type ProductDetail map[string]interface{}

type ProductResult []ProductDetail

type ResponseErrors map[string]string

type ResponseCategory []map[string]string

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
