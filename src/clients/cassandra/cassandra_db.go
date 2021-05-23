package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func GetSession() *gocql.Session {
	// Connect to CAssandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

	return session
}
