package models

type Results struct {
	Vowels         int `json:"vowels,omitempty"`
	Spaces         int `json:"spaces,omitempty"`
	Capitalletters int `json:"capitalletters,omitempty"`
	Smallletters   int `json:"smallletters,omitempty"`
	Words          int `json:"words,omitempty"`
}
type DBResults struct {
	Id             int `db:"id"`
	Vowels         int `db:"vowels"`
	Spaces         int `db:"spaces"`
	Capitalletters int `db:"capitalletters"`
	Smallletters   int `db:"smallletters"`
	Words          int `db:"words"`
}
type UpdateField struct {
	Field string `json:"field"`
	Value int    `json:"value"`
	Id    int    `json:"id" db:"id"`
}
type Identity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Identify struct {
	Username string `json: username`
	Password string `json: password`
}
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

var RefreshData struct {
	RefreshToken string `json:"refresh_token"`
}

var JokeResponse struct {
	Joke string `json:"joke"`
}
