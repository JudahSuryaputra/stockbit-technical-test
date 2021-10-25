package main

import (
	"stockbit-backend/cfg"
	"stockbit-backend/http"
)

func init() {
	cfg.Init()
}

func main() {
	conn, _ := repository.Conn()

	http.RunServer(conn)
}
