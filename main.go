package apiproject

import (
	"net/http"
)

func main() {
	routes()
	http.ListenAndServe(":8090", nil)
}
