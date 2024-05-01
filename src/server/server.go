package main

import (
	"github.com/vlad-marlo/pgx-example/internal/controller/http"
	"github.com/vlad-marlo/pgx-example/internal/repository/memo"
)

func main() {
	store := memo.New()
	ctrl := http.New(store)
	if err := ctrl.Start(); err != nil {
		panic(err)
	}
}
