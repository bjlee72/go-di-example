package component

// This file defines the core business logic of this component.

type Component struct {
	dataProvider DataProvider
}

func NewComponent(dp DataProvider) *Component {
	return &Component{
		dataProvider: dp,
	}
}

func (c *Component) GetAll() []*Record {
	// Please note that db.Query returns []*db.Row --> thus mandates an Adaptor to be used with this logic
	return c.dataProvider.Query()
}
