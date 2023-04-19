package main

import (
	"gatorrater/router"
)

// test with curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"username\": \"diminikov\",\"password\": \"1234\"}"
func main() {
	r := router.SetupRouter()
	_ = r.Run(":8080")
}
