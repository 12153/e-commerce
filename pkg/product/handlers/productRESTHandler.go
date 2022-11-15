package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/12153/e-commerce/pkg/common/recordKeeper"
	"github.com/12153/e-commerce/pkg/common/recordKeeper/mongoKeeper"
	"github.com/12153/e-commerce/pkg/product"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewProductHandler(r *mux.Router) {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	keeper := mongoKeeper.NewRecordKeeper(
		*client.Database("e-commerce"),
		"products",
	)

	sr := r.Path("/products").Subrouter()
	sr.HandleFunc("/{id}", GetProduct(keeper)).Methods(http.MethodGet)
	sr.HandleFunc("", CreateProduct(keeper)).Methods(http.MethodPost)

}

func GetProduct(keeper recordKeeper.RecordKeeper) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Ctx(r.Context()).Err(err).Msg("Error parsing id")
		}

		w.Write([]byte(fmt.Sprintf("{id: %d}", id)))
	}
}

func CreateProduct(keeper recordKeeper.RecordKeeper) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("cool")
		var p product.Product
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("cool: %s", err)
		}
		err = json.Unmarshal(b, &p)
		if err != nil {
			fmt.Printf("cool: %s", err)
		}

		_, err = keeper.CreateEntity(r.Context(), p)
		if err != nil {
			fmt.Printf("cool: %s", err)
		}
	}
}
