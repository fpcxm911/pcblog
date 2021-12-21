package pcblog

import (
	"errors"
	"pcblog/database"

	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) Register(request database.User, reply *LoginResponse) error {

	db, err := database.SetupDefaultDatabase()
	if err != nil {
		return err
	}

	// check if the username has been used
	//var repeatedUser = database.User{}
	_, err = request.FindOne(db)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(RegisterUsernameUsed)
	}

	// create new account in user table
	userinfo, err := request.CreateNewUser(db)
	if err != nil {
		return err
	}

	// Search the Inbox for the new user
	inbox := database.Inbox{}
	findInboxes, err := inbox.FindInbox(db, userinfo)
	if err != nil {
		return err
	}

	*reply = LoginResponse{
		Db:          db,
		LoginStatus: LoginSuccess,
		Inbox:       findInboxes,
		UserInfo:    userinfo,
	}

	return nil
}
