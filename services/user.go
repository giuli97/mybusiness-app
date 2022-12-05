package services

import (
	"my-app-server/helpers"
	"my-app-server/types"
)

func CreateUser(user types.User) error {
	bd, err := helpers.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO users (name, userName, password) VALUES (?, ?, ?)", user.Name, user.UserName, user.Password)
	return err
}

func DeleteUser(id int64) error {
	bd, err := helpers.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func UpdateUser(user types.User) error {
	bd, err := helpers.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE users SET name = ?, userName = ?, password = ? WHERE id = ?", user.Name, user.UserName, user.Password, user.Id)
	return err
}
func GetUsers() ([]types.User, error) {
	//Declare an array because if there's error, we return it empty
	users := []types.User{}
	bd, err := helpers.GetDB()
	if err != nil {
		return users, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, name, userName, password FROM users")
	if err != nil {
		return users, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var user types.User
		err = rows.Scan(&user.Id, &user.Name, &user.UserName, &user.Password)
		if err != nil {
			return users, err
		}
		// and append it to the array
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (types.User, error) {
	var user types.User
	bd, err := helpers.GetDB()
	if err != nil {
		return user, err
	}
	row := bd.QueryRow("SELECT id, name, userName, password FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.UserName, &user.Password)
	if err != nil {
		return user, err
	}
	// Success!
	return user, nil
}

func GetUserByUsername(userName string) (types.User, error) {
	var user types.User
	bd, err := helpers.GetDB()
	if err != nil {
		return user, err
	}
	row := bd.QueryRow("SELECT id, name, userName, password FROM users WHERE userName = ?", userName)
	err = row.Scan(&user.Id, &user.Name, &user.UserName, &user.Password)
	if err != nil {
		return user, err
	}
	// Success!
	return user, nil
}
