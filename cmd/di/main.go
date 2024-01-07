package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bjlee72/go-di-example/cmd/di/internal/adaptor"
	"github.com/bjlee72/go-di-example/component"
	"github.com/bjlee72/go-di-example/db"
)

func main() {
	cmp := component.New(
		adaptor.ToDataProvider(
			db.New()))

	fmt.Println(cmp.GetAll())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Hello: %v", cmp.GetAll()))
	})

	http.ListenAndServe(":8080", nil)
}
