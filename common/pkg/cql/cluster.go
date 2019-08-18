package cql

import (
	"time"

	"github.com/gocql/gocql"
)

func NewCluster(localDC, keyspace string, hosts ...string) *gocql.ClusterConfig {
	c := gocql.NewCluster(hosts...)
	c.Keyspace = keyspace
	//c.Authenticator = gocql.PasswordAuthenticator{}
	if len(hosts) > 1 {
		fallback := gocql.RoundRobinHostPolicy()
		if localDC != "" {
			fallback = gocql.DCAwareRoundRobinPolicy(localDC)
		}
		c.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(fallback)
	}
	c.Timeout = time.Second * 30
	c.SocketKeepalive = time.Second * 30
	c.NumConns = 10
	c.Compressor = &gocql.SnappyCompressor{}
	c.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{NumRetries: 3, Max: c.Timeout / 2}
	c.Consistency = gocql.Quorum
	if localDC != "" {
		c.Consistency = gocql.LocalQuorum
	}

	return c
}
