package database

import (
	"database/sql"
	"real-time-backend/backend/modals"
)

func GetUserTable() ([]modals.UserRegistration, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var user modals.UserRegistration
		err := rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Date, &user.Img)
		return user, err
	}
	results, err := FetchDb("SELECT * FROM Users", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.UserRegistration](results)
}

func GetCategoriesTable() ([]modals.Categories, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var categorie modals.Categories
		err := rows.Scan(&categorie.Id, &categorie.Name, &categorie.Description)
		return categorie, err
	}
	results, err := FetchDb("SELECT * FROM Categories", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Categories](results)
}

func GetPostTable() ([]modals.Post, error) {
	executor := func(rows *sql.Rows) (interface{}, error) {
		var post modals.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Creation, &post.Title, &post.Description)
		return post, err
	}
	results, err := FetchDb("SELECT * FROM Posts", executor)
	if err != nil {
		return nil, err
	}
	return ConvertResults[modals.Post](results)
}
