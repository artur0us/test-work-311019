package accounts

import (
	"log"
	"time"

	"github.com/artur0us/test-work-311019/drivers/pgsqldrv"
	"github.com/artur0us/test-work-311019/entities/accounts/mdls"
)

func GetTinyInfoByToken(token string) (bool, mdls.AccountTinyInfo) {
	var accountTinyInfo mdls.AccountTinyInfo

	rows, err := pgsqldrv.All["developer_notes"].Query(`
		SELECT
			accounts.id,
			accounts.username,
			accounts.created_at,
			accounts_info.account_group_id
		FROM accounts
		INNER JOIN accounts_info ON accounts_info.account_id = accounts.id
		INNER JOIN accounts_sessions ON accounts_sessions.account_id = accounts.id
		WHERE accounts_sessions.token = $1 AND $2 < expires_at
	`, token, time.Now().Unix())
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ accounts.GetTinyInfoByToken() #1 ]\nMore info:")
		log.Println(err)
		return false, accountTinyInfo
	}
	for rows.Next() {
		err := rows.Scan(
			&accountTinyInfo.AccountID,
			&accountTinyInfo.Username,
			&accountTinyInfo.CreatedAt,
			&accountTinyInfo.AccountGroupID,
		)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ accounts.GetTinyInfoByToken() #2 ]\nMore info:")
			log.Println(err)
			return false, accountTinyInfo
		}

		return true, accountTinyInfo
	}

	return false, accountTinyInfo
}

func GetIDByToken(token string) int64 {
	if token == "" {
		return -1
	}

	rows, err := pgsqldrv.All["developer_notes"].Query(`
		SELECT
			accounts_sessions.account_id
		FROM accounts_sessions
		WHERE token = $1 AND $2 < expires_at
	`, token, time.Now().Unix())
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ accounts.GetIDByToken() #1 ]\nMore info:")
		log.Println(err)
		return -1
	}
	for rows.Next() {
		var accountId int64
		err := rows.Scan(&accountId)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ accounts.GetIDByToken() #2 ]\nMore info:")
			log.Println(err)
			return -1
		}

		return accountId
	}

	return -1
}
