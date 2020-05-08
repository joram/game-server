package main

import (
	"net/http"
	"strconv"
)

func getParam(r *http.Request, key string) int {
	vals, ok := r.URL.Query()[key]
	if !ok {
		panic("unable to get param")
	}
	if len(vals) != 1 {
		panic("didn't get the right number of vals")
	}
	val := vals[0]
	i1, _ := strconv.Atoi(val)
	return i1
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
