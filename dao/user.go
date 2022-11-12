package dao

var User map[string]map[string]string

func InitMap() {
	User = make(map[string]map[string]string)
}
func Adduser(username, password, question, answer string) {
	User[username]["username"] = username
	User[username]["password"] = password
	User[username]["question"] = question
	User[username]["answer"] = answer

}

func SearchUser(username string) bool {
	if User[username][username] == "" {
		return false
	}
	return true
}

func Searchpassword(username string) string {
	return User[username]["password"]
}
func Change(username, password string) {
	User[username]["password"] = password
}

func Check(username, answer string) bool {
	if User[username]["answer"] == answer {
		return true
	}
	return false
}

func Searchquestion(username string) string {
	return User[username]["question"]
}

func Delate(username string) {
	delete(User, User[username]["username"])
	delete(User, User[username]["password"])
	delete(User, User[username]["question"])
	delete(User, User[username]["answer"])
}
