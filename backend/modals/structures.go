package modals

// Json type
type UserRegistration struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Date      string `json:"date"`
	Img       string `json:"image"`
}

type Categories struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Post struct {
	Id          string `json:"id"`
	UserId      string `json:"userid"`
	Creation    string `json:"created"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Comments struct {
	Id      string `json:"id"`
	PostId  string `json:"postid"`
	Userid  string `json:"userid"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
