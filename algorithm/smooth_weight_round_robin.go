package algorithm

import "github.com/tensorchen/go-load-balancer/model"

//平滑加权轮询策略
func SmoothWeightRoundRobin(s *model.Service) *model.Node {
	totalWeight := 0

	//遍历所有节点
	//累计权重到totalWeight
	//每个节点CurrentWeight自增Weight
	for _, node := range s.Nodes {
		node.CurrentWeight += node.Weight
		totalWeight += node.Weight
	}


	//计算CurrentWeight最大节点
	cw := s.Nodes[0].CurrentWeight
	maxI := 0

	for i, node := range s.Nodes {
		if node.CurrentWeight > cw {
			cw = node.CurrentWeight
			maxI = i
		}
	}

	//CurrentWeight最大节点减去总权重
	s.Nodes[maxI].CurrentWeight -= totalWeight

	//节点列表中CurrentWeight最大的为本次选中节点
	return s.Nodes[maxI]
}