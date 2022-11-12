package dao

var user = map[string]string{
	"lsx":    "123456",
	"hiahia": "654321",
}
var question = map[string]string{
	"lsx":    "输入1",
	"hiahia": "aa",
}
var answer = map[string]string{
	"lsx":    "1",
	"hiahia": "aa",
}

func Adduser(username, password, question1, answer1 string) {
	user[username] = password
	question[username] = question1
	answer[username] = answer1

}

func SearchUser(username string) bool {
	if user[username] == "" {
		return false
	}
	return true
}

func Searchpassword(username string) string {
	return user[username]
}
func Change(username, password string) {
	user[username] = password
}

func Check(username, answer1 string) bool {
	if answer[username] == answer1 {
		return true
	}
	return false
}

func Searchquestion(username string) string {
	return question[username]
}

func Delate(username string) {
	delete(user, user[username])
	delete(question, question[username])
	delete(answer, answer[username])
}
