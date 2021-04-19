package database

import (
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"
)

func (db *CyberDb) SaveCyberlink(
	cyberlinks []graphtypes.Link,
	agent string,
	timestamp string,
	height int64,
	txHash string,
	) error {
	query := `
		INSERT INTO cyberlinks (object_from, object_to, subject, timestamp, height, transaction_hash)
		VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING
	`
	for i, _ := range cyberlinks {
		_, err := db.Sql.Exec(query,
			cyberlinks[i].From,
			cyberlinks[i].To,
			agent,
			timestamp,
			height,
			txHash,
		)
		if err != nil {
			return err
		}
	}
	return nil
}