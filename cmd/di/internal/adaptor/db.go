package adaptor

import (
	"github.com/bjlee72/go-di-example/component"
	"github.com/bjlee72/go-di-example/db"
)

type DBDataProvider struct {
	*db.Database
}

func (ddp *DBDataProvider) Query() []*component.Record {
	return ddp.toRecords(ddp.Database.Query())
}

func (ddp *DBDataProvider) toRecords(rows []*db.Row) []*component.Record {
	result := make([]*component.Record, 0)
	for _, row := range rows {
		result = append(result, &component.Record{
			ID: row.ID,
		})
	}
	return result
}

// ToDataProvider converts db.Database instance to component.DataProvider instance.
func ToDataProvider(db *db.Database) component.DataProvider {
	return &DBDataProvider{db}
}
