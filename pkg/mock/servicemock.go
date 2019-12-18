package mock

type FakeDB struct{}

type FakeStore struct{}

func (db *FakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}
