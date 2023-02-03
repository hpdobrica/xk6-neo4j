package neo4j

import (
	"github.com/hpdobrica/xk6-neo4j/neo4j"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/neo4j", neo4j.New())
}
