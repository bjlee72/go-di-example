package component

//go:generate mockgen -destination=mock_dataprovider_test.go -package=component -source=interfaces.go DataProvider
type DataProvider interface {
	Query() []*Record // Not *db.Row type
}
