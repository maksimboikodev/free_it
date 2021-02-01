package gorilla

import (
	"fmt"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	//response := fmt.Sprintf("Product %s" /*, id*/)
	response := fmt.Sprintf("TEST")
	fmt.Fprint(w, response)
}
