package http

import (
	"encoding/json"
	"log"
	"net/http"
	"orderAPI/service/internal/usecase"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	ucOrder	usecase.Order
}

func NewHandler(ucOrder	usecase.Order) *OrderHandler{
	return &OrderHandler{
		ucOrder: ucOrder,
	}
}

func (h *OrderHandler)InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/orders/{id}", h.getOrder).Methods("GET")
	return router
}

func (h *OrderHandler) getOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    orderID := string(params["id"])
	order, err := h.ucOrder.GetByUID(orderID)
	if err != nil {
		log.Println("error", err)
	}
	json.NewEncoder(w).Encode(order)
}

