package notes

import (
	"database/sql"
	"log"
	"time"

	"../../drivers/pgsqldrv"
	"../accounts"
	"./mdls"
)

func Create(token string, note mdls.NewNote) (bool, string) {
	flag, accountTinyInfo := accounts.GetTinyInfoByToken(token)
	if !flag || accountTinyInfo.AccountID < 1 {
		return false, "Error occurred while getting information about account by authentication token!"
	}

	_, err := pgsqldrv.All["developer_notes"].Exec(`
		INSERT INTO notes
		(author_account_id, created_at, title, body)
		VALUES($1, $2, $3, $4, $5)
	`, accountTinyInfo.AccountID, time.Now().Unix(), note.Title, note.Body)
	if err != nil {
		log.Println("[i] SQL query error! [ notes.Create() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: n.c #1"
	}

	return true, "New note created successfully!"
}

func GetAll(token string) (bool, string, []mdls.Note) {
	notes := []mdls.Note{}

	flag, accountTinyInfo := accounts.GetTinyInfoByToken(token)
	if !flag || accountTinyInfo.AccountID < 1 {
		return false, "Error occurred while getting information about account by authentication token!", notes
	}

	var rows *sql.Rows
	var err error
	if accountTinyInfo.AccountGroupID == 1 {
		rows, err = pgsqldrv.All["developer_notes"].Query(`
			SELECT
				id, author_account_id, created_at, title, body
			FROM notes
			INNER JOIN accounts ON accounts.id = notes.author_account_id
		`)
	} else {
		rows, err = pgsqldrv.All["developer_notes"].Query(`
			SELECT
				id, author_account_id, created_at, title, body
			FROM notes
			INNER JOIN accounts ON accounts.id = notes.author_account_id
			WHERE notes.author_account_id = $1
		`, accountTinyInfo.AccountID)
	}
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ notes.GetAll() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: n.ga #1", notes
	}
	for rows.Next() {
		var currentNote mdls.Note
		err := rows.Scan(
			&currentNote.AuthorAccountID,
			&currentNote.CreatedAt,
			&currentNote.Title,
			&currentNote.Body,
		)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ notes.GetAll() #2 ]\nMore info:")
			log.Println(err)
			return false, "Server error! Code: n.ga #2", notes
		}

		notes = append(notes, currentNote)
	}

	return true, "OK", notes
}

func GetByID(token string, id int64) (bool, string, mdls.Note) {
	var note mdls.Note

	flag, accountTinyInfo := accounts.GetTinyInfoByToken(token)
	if !flag || accountTinyInfo.AccountID < 1 {
		return false, "Error occurred while getting information about account by authentication token!", note
	}

	rows, err := pgsqldrv.All["developer_notes"].Query(`
		SELECT
			id, author_account_id, created_at, title, body
		FROM notes
		INNER JOIN accounts ON accounts.id = notes.author_account_id
		WHERE notes.id = $1
	`, accountTinyInfo.AccountID)
	defer rows.Close()
	if err != nil {
		log.Println("[i] SQL query error! [ notes.GetByID() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: n.gbi #1", note
	}
	for rows.Next() {
		err := rows.Scan(
			&note.AuthorAccountID,
			&note.CreatedAt,
			&note.Title,
			&note.Body,
		)
		if err != nil {
			log.Println("[i] SQL answer parse error! [ notes.GetByID() #2 ]\nMore info:")
			log.Println(err)
			return false, "Server error! Code: n.gbi #2", note
		}

		if accountTinyInfo.AccountID != note.AuthorAccountID {
			return false, "You do not have access to this note!", note
		}

		return true, "OK", note
	}

	return false, "Unknown error occurred while getting specific note by id!", note
}

func DeleteAll(token string) (bool, string) {
	flag, accountTinyInfo := accounts.GetTinyInfoByToken(token)
	if !flag || accountTinyInfo.AccountID < 1 {
		return false, "Error occurred while getting information about account by authentication token!"
	}

	_, err := pgsqldrv.All["developer_notes"].Exec(`
		DELETE FROM notes WHERE author_account_id = $1
	`, accountTinyInfo.AccountID)
	if err != nil {
		log.Println("[i] SQL query error! [ notes.DeleteByID() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: n.dbi #1"
	}

	return true, "Note removed successfully!"
}

func DeleteByID(token string, id int64) (bool, string) {
	flag, accountTinyInfo := accounts.GetTinyInfoByToken(token)
	if !flag || accountTinyInfo.AccountID < 1 {
		return false, "Error occurred while getting information about account by authentication token!"
	}

	flag, msg, noteInfo := GetByID(token, id)
	if !flag {
		return false, msg
	}

	if noteInfo.AuthorAccountID != accountTinyInfo.AccountID {
		return false, "You do not have access to this note!"
	}
	_, err := pgsqldrv.All["developer_notes"].Exec(`
		DELETE FROM notes WHERE id = $1
	`, accountTinyInfo.AccountID)
	if err != nil {
		log.Println("[i] SQL query error! [ notes.DeleteByID() #1 ]\nMore info:")
		log.Println(err)
		return false, "Server error! Code: n.dbi #1"
	}

	return true, "Note removed successfully!"
}
