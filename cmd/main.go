package main

import "github.com/Guilherme-De-Marchi/ai-hub/api"

func main() {
	srv := api.NewServer(":8080")
	srv.Start()
}
