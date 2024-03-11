package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/micael-ortega/crebito/internal/dto"
	"github.com/micael-ortega/crebito/internal/dto/request"
	"github.com/micael-ortega/crebito/internal/dto/response"
	"github.com/micael-ortega/crebito/internal/service"
	"github.com/micael-ortega/crebito/internal/utils"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod("post", w, r) {
		handleErrResponse(w, fmt.Errorf("method shall not pass"), http.StatusMethodNotAllowed)
		return
	}

	pathId := r.PathValue("id")

	id, err := strconv.Atoi(pathId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body request.TransactionRequest

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = utils.CheckValidBody(&body)
	if err != nil{
		http.Error(w, "body value not accepted", http.StatusUnprocessableEntity)
		return
	}

	var dto *dto.TransactionDTO
	var transactionErr error
	switch body.Kind {
	case "c":
		dto, transactionErr = c.service.CreditTransaction(body.Value, body.Description, id)
	case "d":
		dto, transactionErr = c.service.DebitTransaction(body.Value, body.Description, id)
	default:
		http.Error(w, "unknown transaction method", http.StatusUnprocessableEntity)
		return
	}

	if transactionErr != nil {
		http.Error(w, transactionErr.Error(), http.StatusInternalServerError)
		return
	}

	switch dto.Code {
	case 2:
		http.Error(w, "insuficient limit", http.StatusUnprocessableEntity)
		return 
	case 3:
		http.Error(w, "user not found", http.StatusNotFound)
		return 
	}

	res := response.TransactionResponse{
		Balance: dto.Balance,
		Limit: dto.Limit,
	}

	jsonResponse, err := writeJson(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, http.StatusOK, jsonResponse)

}

func (c *Controller) GetBankStatement(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		handleErrResponse(w, err, http.StatusBadRequest)
		return
	}

	response, err := c.service.GetLastTransactions(id)
	if err != nil {
		handleErrResponse(w, err, http.StatusNotFound)
		return
	}

	jsonResponse, err := writeJson(response)
	if err != nil {
		handleErrResponse(w, err, http.StatusInternalServerError)
		return
	}

	writeResponse(w, http.StatusOK, jsonResponse)

}

func writeJson(T interface{}) ([]byte, error) {
	jsonResponse, err := json.Marshal(T)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func handleErrResponse(w http.ResponseWriter, err error, code int) {
	w.Header().Add("error", err.Error())
	w.WriteHeader(code)
}

func writeResponse(w http.ResponseWriter, status int, jsonResponse []byte){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
