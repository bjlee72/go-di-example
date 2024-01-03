package feature

//go:generate mockgen -destination=mock_dataprovider_test.go -package=feature -source=interfaces.go DataProvider
type DataProvider interface {
	Provide(int) (int, error)
}
