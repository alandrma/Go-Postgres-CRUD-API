package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/buku", controller.AmbilSemuaBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")

	// soon
	// router.HandleFunc("/api/tmbhbuku", controller.TmbhBuku).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/hapusbuku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
