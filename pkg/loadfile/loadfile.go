package loadfile

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ServeStaticFile(staticName string) {
	r := mux.NewRouter()
	r.HandleFunc(fmt.Sprintf("/static/%s/", staticName), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("static/%s/index.html", staticName))
		w.WriteHeader(http.StatusOK)
	})
	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:7000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("ListenServe: ", err)
	}
}
