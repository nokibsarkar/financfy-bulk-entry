package models

type ResponseSingle[T interface{}] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}
type ResponseMultiple[T interface{}] struct {
	Data  []T    `json:"data"`
	Error string `json:"error"`
}
