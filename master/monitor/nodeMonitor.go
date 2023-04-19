package monitor

import (
	"k2edge/etcdutil"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
)

func nodeStatus(node types.Node, s *svc.ServiceContext) error {
	key := etcdutil.GenerateKey("node", etcdutil.SystemNamespace, node.Metadata.Name)
	node.Status.Working = false
	return etcdutil.DeleteOne(s.Etcd, s.Etcd.Ctx(), key)
}