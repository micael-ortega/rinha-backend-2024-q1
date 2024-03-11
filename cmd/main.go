package main

import (
	"log"
	"net/http"

	"github.com/micael-ortega/crebito/internal/controller"
	"github.com/micael-ortega/crebito/internal/database"
	"github.com/micael-ortega/crebito/internal/repository"
	"github.com/micael-ortega/crebito/internal/service"
	"github.com/micael-ortega/crebito/internal/utils"
)

func main() {

	utils.LoadEnv()

	pool := database.NewPool()
	r := repository.NewRepository(pool)
	s := service.NewService(r)
	c := controller.NewController(s)

	mux := http.NewServeMux()

	mux.HandleFunc("/clientes/{id}/transacoes", func(w http.ResponseWriter, r *http.Request) {
		c.CreateTransaction(w, r)
	})

	mux.HandleFunc("/clientes/{id}/extrato", func(w http.ResponseWriter, r *http.Request) {
		c.GetBankStatement(w, r)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
