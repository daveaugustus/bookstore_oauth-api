package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {

}

func GetSession() (*gocql.Session, error) {
	// Connect to CAssandra cluster:
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	return cluster.CreateSession()
}
