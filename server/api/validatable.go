package api

type validatable interface {
	IsValid() *RequestError
}
