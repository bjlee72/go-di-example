package db

type Database struct {
	// DB internal
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Query() []*Row {
	// This is just for the explanation purpose. Not a real code.
	return []*Row{
		{ID: "A"},
		{ID: "B"},
		{ID: "C"},
	}
}
