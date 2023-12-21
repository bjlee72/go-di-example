package adaptor

import (
	"reflect"
	"testing"

	"github.com/bjlee72/go-di-example/component"
	"github.com/bjlee72/go-di-example/db"
)

func TestDBDataProvider_toRecords(t *testing.T) {
	tests := []struct {
		name string
		db   *db.Database
		rows []*db.Row
		want []*component.Record
	}{
		{
			name: "Test the basic conversion works fine",
			db:   &db.Database{},
			rows: []*db.Row{{ID: "A"}, {ID: "B"}, {ID: "C"}},
			want: []*component.Record{{ID: "A"}, {ID: "B"}, {ID: "C"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddp := &DBDataProvider{
				Database: tt.db,
			}
			if got := ddp.toRecords(tt.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
