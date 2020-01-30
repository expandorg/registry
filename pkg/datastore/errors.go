package datastore

type NoRowErr struct{}

func (err NoRowErr) Error() string {
	return "No Records"
}
