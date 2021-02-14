package cassandra

import (
	"github.com/gocql/gocql"
)

// var (
// // cluster *gocql.ClusterConfig
// )

// func init() {
// 	// Connect to cassandra cluster

// }

func GetSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth2"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}
