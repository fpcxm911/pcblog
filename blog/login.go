package pcblog

import (
	"errors"
	"log"
	"pcblog/database"

	"gorm.io/gorm"
)

func (u *UserService) Login(request database.User, reply *LoginResponse) error {
	db, err := database.SetupDefaultDatabase()
	if err != nil {
		log.Println("database connection error:", err)
		return err
	}

	rightAccount, err := request.FindOne(db)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Wrong username received")
		return errors.New(WrongUsername)
	}

	if rightAccount.Password != request.Password {
		log.Println("Wrong password received")
		return errors.New(WrongPassword)
	}

	// Search the Inbox for the user
	inbox := database.Inbox{}
	findInboxes, err := inbox.FindInbox(db, rightAccount)
	if err != nil {
		log.Println("Inbox search failed:", err)
		return err
	}

	log.Println("Preparing login response...")
	*reply = LoginResponse{
		Db:          db,
		LoginStatus: LoginSuccess,
		Inbox:       findInboxes,
		UserInfo:    rightAccount,
	}
	log.Println("Login response completed.")
	return nil
}
