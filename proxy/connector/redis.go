package connector

// RedisConnector implemented Connector
type RedisConnector struct {
	ConnectorCore
}

// Run RedisConnector
func (c *RedisConnector) Run() error {
	return c.ConnectorCore.Run()
}
