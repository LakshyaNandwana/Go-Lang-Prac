package dao


type User struct{
	ID int
	Name string
}


type UserDaoInterface interface{
	GetUser(ID int)(User, error)
}