package game

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/utils"
	"net/http"
)


func ServePixels(w http.ResponseWriter, r *http.Request){
	utils.EnableCors(&w)

	x1 := utils.GetParam(r, "x1")
	y1 := utils.GetParam(r, "y1")
	x2 := utils.GetParam(r, "x2")
	y2 := utils.GetParam(r, "y2")
	fmt.Printf("(%d, %d), (%d, %d)\n", x1, y1, x2, y2)
	json.NewEncoder(w).Encode(utils.GetPixels(x1,y1,x2,y2))
}

