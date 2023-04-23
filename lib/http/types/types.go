package modularHTTP

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type MiddleHandler func(w http.ResponseWriter, r *http.Request) Handler

type Router struct {
	Name     string
	MuxRoute []*mux.Route
}

type BaseResponse struct {
	Status int                    `json:"status"`
	Alert  map[string]interface{} `json:"alert"`
	Data   interface{}            `json:"body"`
}

func GetIpAddress() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP, nil
		}
	}

	return nil, err

}
