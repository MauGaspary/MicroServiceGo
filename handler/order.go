package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criar uma ordem")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listar todas as ordens")
}

func (o *Order) getByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pegar uma ordem por ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Atualizar uma ordem por ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletar uma ordem por ID")
}

