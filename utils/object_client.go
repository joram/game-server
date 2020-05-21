package utils
import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type ObjectClient struct {
	C      *websocket.Conn
	Player BaseMonsterInterface
	Mux    *sync.Mutex
}

var ObjectClients = []ObjectClient{}

func (cw *ObjectClient) ReadMessage() (map[string]interface{}, error) {
	//cw.Mux.Lock()
	//defer cw.Mux.Unlock()

	_, message, err := cw.C.ReadMessage()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(message), &result)
	return result, nil
}

func (cw *ObjectClient) RemoveObject(object ObjectInterface) {
	cw.Mux.Lock()
	defer cw.Mux.Unlock()

	type removeMessage struct {
		Action string `json:"action"`
		ObjectID int `json:"id"`
	}

	rm := removeMessage{"remove", object.GetID()}
	jsonString, err := json.Marshal(rm)
	if err != nil {
		log.Println("write:", err)
	}
	cw.C.WriteMessage(1, []byte(jsonString))
}

func (cw *ObjectClient) UpdateObject(object ObjectInterface) {
	cw.Mux.Lock()
	defer cw.Mux.Unlock()

	jsonString := object.AsString()
	cw.C.WriteMessage(1, []byte(jsonString))
}


func BroadcastLocationChange(object ObjectInterface, objectClients []ObjectClient){
	for _, client := range objectClients {
		client.UpdateObject(object)
	}
}
