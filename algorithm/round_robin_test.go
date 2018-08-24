package algorithm_test

import (
	"fmt"

	"github.com/tensorchen/go-load-balancer/algorithm"
	"github.com/tensorchen/go-load-balancer/model"
)

func Example_RoundRobin() {

	node1 := &model.Node{
		Address: "node_1",
	}
	node2 := &model.Node{
		Address: "node_2",
	}
	node3 := &model.Node{
		Address: "node_3",
	}
	s := &model.Service{
		Nodes: []*model.Node{node1, node2, node3},
	}

	for i := 0; i < 7; i++ {
		fmt.Println(algorithm.RoundRobin(s).Address)
	}
	//Output:
	//node_1
	//node_2
	//node_3
	//node_1
	//node_2
	//node_3
	//node_1
}
