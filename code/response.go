package code

type response struct {
	data    any
	code    int
	message string
	err     error
}

func (r *response) Data() any {
	return r.data
}

func (r *response) Code() int {
	return r.code
}

func (r *response) Message() string {
	return r.message
}

func (r *response) Error() error {
	return r.err
}
