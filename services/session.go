package services

import (
	"my-app-server/database"
	"my-app-server/models"
	"time"
)

func CreateSession(token string) error {
	bd, err := database.GetDB()
	if err != nil {
		return err
	}
	expiresIn := 3600
	timeNow := time.Now()
	_, err = bd.Exec("INSERT INTO sessions (token, expiresIn, lastUpdate) VALUES (?, ?, ?)", token, expiresIn, timeNow)
	return err
}

func DeleteSession(id int64) error {
	bd, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM sessions WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func UpdateSession(token string, expiresIn string, user models.User) error {
	bd, err := database.GetDB()
	if err != nil {
		return err
	}
	timeNow := time.Now()
	_, err = bd.Exec("UPDATE sessions SET token = ?, expiresIn = ?, lastUpdate = ? WHERE userId = ?", token, expiresIn, timeNow, user.Id)
	return err
}
func GetSessions() ([]models.Session, error) {
	//Declare an array because if there's error, we return it empty
	sessions := []models.Session{}
	bd, err := database.GetDB()
	if err != nil {
		return sessions, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, token, expiresIn, lastUpdate FROM sessions")
	if err != nil {
		return sessions, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var session models.Session
		err = rows.Scan(&session.Id, &session.Token, &session.ExpiresIn, &session.LastUpdate)
		if err != nil {
			return sessions, err
		}
		// and append it to the array
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func GetSessionById(id int64) (models.Session, error) {
	var session models.Session
	bd, err := database.GetDB()
	if err != nil {
		return session, err
	}
	row := bd.QueryRow("SELECT id, token, expiresIn, lastUpdate FROM sessions WHERE id = ?", id)
	err = row.Scan(&session.Id, &session.Token, &session.ExpiresIn, &session.LastUpdate)
	if err != nil {
		return session, err
	}
	// Success!
	return session, nil
}

func GetSessionByUserId(userId int64) (models.Session, error) {
	var session models.Session
	bd, err := database.GetDB()
	if err != nil {
		return session, err
	}
	row := bd.QueryRow("SELECT id, token, expiresIn, lastUpdate FROM sessions WHERE userId = ?", userId)
	err = row.Scan(&session.Id, &session.Token, &session.ExpiresIn, &session.LastUpdate)
	if err != nil {
		return session, err
	}
	// Success!
	return session, nil
}
