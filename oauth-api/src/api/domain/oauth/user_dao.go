package oauth

const (
	querySetByUsernameAndPassword = "SELECT id, username, FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User {
		"lindamanf": &User{
			ID: 123
			Username: "lindamanf",
		}
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}
