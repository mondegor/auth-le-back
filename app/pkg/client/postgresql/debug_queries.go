package postgresql

import (
    "fmt"
    "strings"
)

func (c *Connection) debugQuery(query string) {
    query = strings.ReplaceAll(strings.ReplaceAll(query, "\t", ""), "\n", " ")
    c.logger.Debug(fmt.Sprintf("SQL Query: %s", query))
}
