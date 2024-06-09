package jsonfiles

import (
	"forum/backend/database"
	"forum/backend/modals"
)

// Check on DB the Users table to send at handler
// Get data on my type struct
func GetUserTable() ([]modals.UserRegistration, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []modals.UserRegistration

	for rows.Next() {
		var user modals.UserRegistration
		if err != rows.Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Date) {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// Check on DB the Categories table to send at handler
func GetCategoriesTable() ([]modals.Categories, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []modals.Categories

	for rows.Next() {
		var categorie modals.Categories
		if err != rows.Scan(&categorie.Id, &categorie.Name, &categorie.Description) {
			return nil, err
		}
		categories = append(categories, categorie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// Check on DB the Posts table to send at handler
func GetPostTable() ([]modals.Post, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []modals.Post

	for rows.Next() {
		var post modals.Post
		if err != rows.Scan(&post.Id, &post.UserId, &post.Creation, &post.Title, &post.Description) {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
