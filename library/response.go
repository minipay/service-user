package library

// Error Response ...
type ErroResponse struct {
	Message string `json:"message"`
}

type Data = interface{}
type Meta = interface{}

// Error Response ...
type SuccessResponse struct {
	Message string `json:"message"`
	Result Data `json:"result"`
	Meta Meta `json:"meta"`
}