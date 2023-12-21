package main

import (
	"fmt"

	"github.com/bjlee72/go-di-example/cmd/di/internal/adaptor"
	"github.com/bjlee72/go-di-example/component"
	"github.com/bjlee72/go-di-example/db"
)

func main() {
	cmp := component.NewComponent(
		adaptor.ToDataProvider(
			db.NewDatabase()))
	fmt.Println(cmp.GetAll())
}
