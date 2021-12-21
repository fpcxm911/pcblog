package pcblog

import (
	"errors"
	"pcblog/database"

	"gorm.io/gorm"
)

func (u *UserService) Login(request database.User, reply *LoginResponse) error {
	db, err := database.SetupDefaultDatabase()
	if err != nil {
		return err
	}

	rightAccount, err := request.FindOne(db)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(WrongUsername)
	}

	if rightAccount.Password != request.Password {
		return errors.New(WrongPassword)
	}

	// Search the Inbox for the user
	inbox := database.Inbox{}
	findInboxes, err := inbox.FindInbox(db, rightAccount)
	if err != nil {
		return err
	}

	*reply = LoginResponse{
		Db:          db,
		LoginStatus: LoginSuccess,
		Inbox:       findInboxes,
		UserInfo:    rightAccount,
	}
	return nil
}
