package postgresql

import (
    "fmt"
    "strings"
)

func (c *Connection) debugQuery(query string) {
    c.logger.Debug(fmt.Sprintf("SQL Query: %s", strings.Join(strings.Fields(query), " ")))
}
