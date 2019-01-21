package main

type Handler interface {
	Read([]byte) (int64, error)
	Write([]byte) (int64, error)
}

var Manager map[string]Handler

var (
	ErrAlreadyExist = errors.New("already exist it")
)

func Register(path string, h Handler) error {
	if v, ok := Manager[path]; ok {
		return ErrAdreadyExit
	}
	Manager[path] = h
	return nil
}
