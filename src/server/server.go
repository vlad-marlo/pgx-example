package main

import (
	"github.com/vlad-marlo/pgx-example/internal/controller/http"
)

func main() {
	ctrl := http.New()
	if err := ctrl.Start(); err != nil {
		panic(err)
	}
}
