package fetchclient

type FetchError interface {
	SetError(code string, message string)
	IsError() bool
}
