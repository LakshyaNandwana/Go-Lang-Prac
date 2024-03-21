package service

import(
	"MockUnit/dao"
)



type UserService struct{

	userDao dao.UserDaoInterface //dependency to the user dao interface
}

func (u UserService) GetUser(ID int) (dao.User, error){
	return u.userDao.GetUser(ID) //calling user dao getUser
}