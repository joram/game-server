package utils

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
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

func IsSolid(x,y int) bool {
	for _, town := range Towns {
		if town.Contains(x, y) {
			pixel := town.Pixel(x, y)
			return pixel.G > 180
		}
	}
	return GetPixel(x,y).G > 180
}


var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
