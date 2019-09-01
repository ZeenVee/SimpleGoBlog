package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func respondwithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func respondwithJSON(w http.ResponseWriter, code int, posts interface{}) {

	response, _ := json.Marshal(posts)
	//fmt.Println(posts)
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}
