package customerHandler

import (
	"bookingtogo-test-be/modules/customers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	CustomerUC customers.CustomerUsecase
	r          *mux.Router
}

func NewCustomerHandler(r *mux.Router, CustomerUC customers.CustomerUsecase) {
	handler := customerHandler{r: r, CustomerUC: CustomerUC}

	handler.r.HandleFunc("/customer/{id}", handler.getCustomer).Methods("GET")

}

func (h *customerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Param Query id cant empty", http.StatusBadRequest)
		return
	}
	myID, errConvertion := strconv.Atoi(id)

	if errConvertion != nil {
		http.Error(w, errConvertion.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.CustomerUC.FindCustomerById(uint64(myID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var jsonData, errMarshal = json.Marshal(result)
	if errMarshal != nil {
		fmt.Println(errMarshal.Error())
		return
	}

	w.Write(jsonData)
	return
}
