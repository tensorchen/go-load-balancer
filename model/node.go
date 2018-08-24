package model

type Node struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`

	CurrentWeight int
}
