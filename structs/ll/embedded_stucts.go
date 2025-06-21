package structs

type User struct {
	id             string
	Name           string
	Level          uint8
	*PrivateFields // embedded struct
}

type PrivateFields struct {
	email    string
	password string
	Age      uint8 // declared capital as to see whether we can access in main func or not
}

func CreateUser() User {
	return User{
		id:    "1234",
		Name:  "abc",
		Level: 10,
		PrivateFields: &PrivateFields{
			email:    "abc@gmail.com",
			password: "password",
			Age:      20,
		},
	}
}
