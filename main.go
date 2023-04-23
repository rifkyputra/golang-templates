package main

import (
	"fmt"

	server "ModularHTTPGo"

	lib "github.com/rifkyputra/formatters"
)

func main() {

	json := `{"name":"John","age":30,"city":"New York","pets":[{"name":"Fluffy","type":"dog"},{"name":"Whiskers","type":"cat"}]}`

	fmt.Println(lib.BeautifyJSON(json))

	server.StartServer(":8080")

}
