package nux

type Response interface {
	Data() any
	Code() int
	Message() string
	Error() error
}

const Success = 0

var MapSuccessMsg = map[int]string{
	Success: "success",
}

type Map map[string]any

func (m Map) Data() any {
	return m
}

func (m Map) Code() int {
	return Success
}

func (m Map) Message() string {
	return MapSuccessMsg[Success]
}

func (m Map) Error() error {
	return nil
}
