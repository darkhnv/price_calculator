package iomanager

// IOManager interface defines methods for reading and writing data
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
