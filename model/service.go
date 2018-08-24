package model

type Service struct {
	Nodes []*Node `json:"nodes"`
	Index int
}
