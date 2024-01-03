package feature

import "github.com/bjlee72/go-di-example/db"

// This file defines the core business logic of this component.

type Feature struct {
	dataProvider DataProvider
}

func New(dp DataProvider) *Feature {
	return &Feature{
		dataProvider: dp,
	}
}

func (c *Feature) Functionality(input int) (int, error) {
	result, err := c.dataProvider.Provide(input)
	/* do something complicated */
	return result, err
}

func Functionality(input int) (int, error) {
	feat := New(db.New())
	return feat.Functionality(input)
}
