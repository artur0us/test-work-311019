package accounts

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/artur0us/test-work-311019/drivers/pgsqldrv"
	"github.com/artur0us/test-work-311019/entities/accounts/mdls"
)

func BasicAuth(username, password string) (bool, string, mdls.BasicAuthAnswer) {
	var basicAuthAnswer mdls.BasicAuthAnswer

	// Basic authentication logic
	// 0. Preparations
	if username == "" || password == "" {
		return false, "Emtpy username or password!", basicAuthAnswer
	}
	// 1. Check password
	rows, err := pgsqldrv.All["developer_notes"].Query(`
		SELECT
			accounts.id,
			accounts.username,
			accounts.created_at,
			accounts_info.account_group_id,
			accounts.password
		FROM accounts
		INNER JOIN accounts_info ON accounts_info.account_id = accounts.id
		WHERE username = $1
	`, username)
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ accounts.BasicAuth() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: a.ba #1", basicAuthAnswer
	}
	for rows.Next() {
		var passwordHashInDB string
		err := rows.Scan(
			&basicAuthAnswer.AccountTinyInfo.AccountID,
			&basicAuthAnswer.AccountTinyInfo.Username,
			&basicAuthAnswer.AccountTinyInfo.CreatedAt,
			&basicAuthAnswer.AccountTinyInfo.AccountGroupID,
			&passwordHashInDB,
		)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ accounts.BasicAuth() #2 ]\nMore info:")
			log.Println(err)
			return false, "Server error! Code: a.ba #2", basicAuthAnswer
		}

		specifiedPasswordHash := sha256.Sum256([]byte(password))
		if hex.EncodeToString(specifiedPasswordHash[:]) != passwordHashInDB {
			return false, "Wrong password! Input right password and try again!", basicAuthAnswer
		}

		// 2. Generate token and return it to client
		token := uuid.New().String()
		createdAt := time.Now().Unix()
		expiresAt := createdAt + ((((60 * 60) * 24) * 7) * 2) // Two weeks
		_, err = pgsqldrv.All["developer_notes"].Exec("INSERT INTO accounts_sessions(account_id, token, created_at, expires_at, user_agent_info) VALUES($1, $2, $3, $4, $5)", basicAuthAnswer.AccountTinyInfo.AccountID, token, createdAt, expiresAt, "Some info about client")
		if err != nil {
			log.Println("[i] SQL query error! [ accounts.BasicAuth() #3 ]\nMore info:")
			log.Println(err)
			return false, "Server error! Code: a.ba #3", basicAuthAnswer
		}
		basicAuthAnswer.Token = token

		return true, "Successfully authenticated!", basicAuthAnswer
	}

	return false, "Unknown authentication error!", basicAuthAnswer
}

func TokenAuth(token string) (bool, string, mdls.AccountTinyInfo) {
	var accountTinyInfo mdls.AccountTinyInfo

	// Token authentication logic
	rows, err := pgsqldrv.All["developer_notes"].Query(`
		SELECT
			accounts.id,
			accounts.username,
			accounts.created_at,
			accounts_info.account_group_id
		FROM accounts_sessions
		INNER JOIN accounts ON accounts.id = accounts_sessions.account_id
		INNER JOIN accounts_info ON accounts_info.account_id = accounts_sessions.account_id
		WHERE token = $1 AND $2 < expires_at
	`, token, time.Now().Unix())
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ accounts.TokenAuth() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: a.ta #1", accountTinyInfo
	}
	for rows.Next() {
		err := rows.Scan(
			&accountTinyInfo.AccountID,
			&accountTinyInfo.Username,
			&accountTinyInfo.CreatedAt,
			&accountTinyInfo.AccountGroupID,
		)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ accounts.TokenAuth() #2 ]\nMore info:")
			log.Println(err)
			return false, "Server error! Code: a.ta #2", accountTinyInfo
		}

		return true, "ok", accountTinyInfo
	}

	return false, "Nonexistent authentication token!", accountTinyInfo
}

func IsTokenAuthed(token string) bool {
	// Token authentication logic
	flag, _, _ := TokenAuth(token)
	return flag
}
