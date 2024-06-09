package modals

// Json type
type UserRegistration struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Date      string `json:"date"`
}

type Categories struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Post struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userid"`
	Creation    string `json:"created"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	Message string `json:"message"`
}
