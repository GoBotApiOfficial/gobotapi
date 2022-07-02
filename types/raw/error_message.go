package raw

type ErrorMessage struct {
	Ok          bool                `json:"ok"`
	Description string              `json:"description"`
	Code        int                 `json:"error_code"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}
