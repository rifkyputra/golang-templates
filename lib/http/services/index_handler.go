package modularHTTP

import (
	. "ModularHTTPGo/types"
	"encoding/json"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ipAddr, _ := GetIpAddress()
	w.Header().Set("Content-type", "Application/json")

	json.NewEncoder(w).Encode(BaseResponse{
		Status: 200,
		Alert: map[string]interface{}{
			"type": 0,
		},
		Data: map[string]interface{}{
			"log": "your ip is : " + ipAddr.String(),
		},
	})

}
