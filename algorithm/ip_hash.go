package algorithm

import (
	"github.com/tensorchen/go-load-balancer/model"
	"hash/fnv"
)

func IpHash(s *model.Service, ip string) *model.Node {

	f := fnv.New32()
	f.Write([]byte(ip))

	hash := int(f.Sum32()) % len(s.Nodes)

	return s.Nodes[hash]
}
