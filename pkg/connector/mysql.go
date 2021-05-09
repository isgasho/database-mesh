package connector

// MySQLConnector implemented Connector
type MySQLConnector struct {
	ConnectorCore
}

// Run MySQLConnector
func (c *MySQLConnector) Run() error {
	return c.ConnectorCore.Run()
}
