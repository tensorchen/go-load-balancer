package algorithm_test

import (
	"fmt"
	"github.com/tensorchen/go-load-balancer/algorithm"
	"github.com/tensorchen/go-load-balancer/model"
)

func Example_IpHash() {

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

	ips := []string{"192.0.0.1", "192.0.0.2", "192.0.0.3"}
	for _, ip := range ips {
		fmt.Println(algorithm.IpHash(s, ip).Address)
	}
	//Output:
	//node_3
	//node_3
	//node_2
}
