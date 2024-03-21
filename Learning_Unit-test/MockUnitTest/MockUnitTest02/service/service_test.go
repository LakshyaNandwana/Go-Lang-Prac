package service

import (
	"fmt"
	"testing"

	"MockUnit/dao"

	"github.com/stretchr/testify/assert"
)


type UserDaoMockFoundUser struct{} 

func (u *UserDaoMockFoundUser) GetUser(ID int)(dao.User, error){
	return dao.User{
		ID: ID,
		Name: "Bruce",
	}, nil
}

type UserDaoMockNotFound struct{}

func (u *UserDaoMockNotFound)GetUser(ID int)(dao.User, error){
	return dao.User{}, fmt.Errorf("user not found")
}


func TestUserService_GetUser(t *testing.T){

	type fields struct{
		userDao dao.UserDaoInterface
	}

	type args struct{
		ID int
	}

	tests := []struct{
		name string
		fields fields
		args args
		want dao.User
		wantErr error
	}{
		{
			name:"should return user when found",
			fields: fields{
				userDao: &UserDaoMockFoundUser{},
			},
			args: args{10},
			want: dao.User{ID:10, Name:"Bruce"},
		},
		{
			name:"should return error when not found",
			fields: fields{
				userDao: &UserDaoMockNotFound{},
			},
			args: args{10},
			wantErr: fmt.Errorf("user not found"),
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			u := UserService{
				userDao: tt.fields.userDao,
			}
			got, err := u.GetUser(tt.args.ID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}