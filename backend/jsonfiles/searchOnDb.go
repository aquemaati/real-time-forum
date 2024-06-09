package jsonfiles

import (
	"database/sql"
	"real-time-backend/backend/database"
	"real-time-backend/backend/modals"
)

func GetUserTable() ([]modals.UserRegistration, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var user modals.UserRegistration
		err := rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Date)
		return user, err
	}
	results, err := database.FetchDb("SELECT * FROM Users", executor)
	if err != nil {
		return nil, err
	}
	return database.ConvertResults[modals.UserRegistration](results)
}

func GetCategoriesTable() ([]modals.Categories, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var categorie modals.Categories
		err := rows.Scan(&categorie.Id, &categorie.Name, &categorie.Description)
		return categorie, err
	}
	results, err := database.FetchDb("SELECT * FROM Categories", executor)
	if err != nil {
		return nil, err
	}
	return database.ConvertResults[modals.Categories](results)
}

func GetPostTable() ([]modals.Post, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var post modals.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Creation, &post.Title, &post.Description)
		return post, err
	}
	results, err := database.FetchDb("SELECT * FROM Posts", executor)
	if err != nil {
		return nil, err
	}
	return database.ConvertResults[modals.Post](results)
}
