package api

import (
	"fmt"
	"github.com/Nivelian/codete-webscraping/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		h.ServeHTTP(w, r)
	})
}

func websiteInfoHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	info, err := GetWebsiteInfo(url)
	if err != nil {
		InternalServerErr(err, w)
		return
	}

	if err := SendStruct(info, w); err != nil {
		InternalServerErr(err, w)
	}
}

func StartServer(config *model.Config) {
	router := mux.NewRouter()
	router.Use(corsHandler)

	// api
	apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	apiRouter.HandleFunc("/scrap", websiteInfoHandler).Queries("url", "{url}").Methods(http.MethodGet)

	// file server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.StaticPath))).Methods(http.MethodGet)

	logrus.Info(fmt.Sprintf("Starting server on port %v", config.Port))
	logrus.Fatal(http.ListenAndServe(config.Port, router))
}
