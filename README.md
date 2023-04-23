## Golang Templates

### http

reusable server boilerplate. using mux router and postgres.

consist of 3 main layers :

- Router
- Middleware
- Service

routers are consumed in `start_server.go` then attached to mux router, db and other adapters. utils are also reusable.

### how to use

import start_server.go to main.go

``` go
package main

import (
 server "ModularHTTPGo"
)
```

create main function
``` go

func main() {

 server.StartServer(":8080")

}

```


