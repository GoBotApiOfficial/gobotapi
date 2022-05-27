package types

type MaskPosition struct {
	Point string `json:"point"`
	Scale float64 `json:"scale"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
}
