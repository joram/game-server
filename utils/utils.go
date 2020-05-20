package utils

import (
	"net/http"
	"strconv"
	"github.com/gorilla/websocket"
)

func GetParam(r *http.Request, key string) int {
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

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
