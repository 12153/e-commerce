package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewProductHandler(r *mux.Router) {
	sr := r.Path("/products").Subrouter()
	sr.HandleFunc("/{id}", GetProduct).Methods(http.MethodGet)

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Ctx(r.Context()).Err(err).Msg("Error parsing id")
	}

	w.Write([]byte(fmt.Sprintf("{id: %d}", id)))
}
