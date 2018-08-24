package algorithm

import "github.com/tensorchen/go-load-balancer/model"

func RoundRobin(s *model.Service) *model.Node {
	if s.Index >= len(s.Nodes) {
		s.Index = 0
	}

	s.Index++
	return s.Nodes[s.Index-1]
}
