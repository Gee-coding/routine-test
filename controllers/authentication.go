package controllers

func Authenticate(username string, password string) bool {
	if username == "test" && password == "1234" {
		return true
	}
	return false
}