package algorithm_test

import (
	"fmt"
	"testing"

	"github.com/tensorchen/go-load-balancer/algorithm"
	"github.com/tensorchen/go-load-balancer/model"
)

func Test_SmoothWeightRoundRobin(t *testing.T) {

}

func Example_SmoothWeightRoundRobin() {
	node1 := &model.Node{
		Address: "node_1",
		Weight:  5,
	}
	node2 := &model.Node{
		Address: "node_2",
		Weight:  1,
	}
	node3 := &model.Node{
		Address: "node_3",
		Weight:  1,
	}
	s := &model.Service{
		Nodes: []*model.Node{node1, node2, node3},
	}

	for i := 0; i < 7; i++ {
		fmt.Println(algorithm.SmoothWeightRoundRobin(s).Address)
	}
	//Output:
	//node_1
	//node_1
	//node_2
	//node_1
	//node_3
	//node_1
	//node_1
}
