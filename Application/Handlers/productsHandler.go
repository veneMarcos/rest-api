package productsHandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Products: %\n", vars["products"])
}
