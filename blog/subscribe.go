package pcblog

import (
	"errors"
	"pcblog/database"
)

func (u *UserService) Subscribe(request [2]database.User, reply *string) error {
	db, err := database.SetupDefaultDatabase()
	if err != nil {
		return err
	}
	user, err := request[0].FindOne(db)
	subTarget, err := request[1].FindOne(db)
	if err != nil {
		return errors.New("the user that you want to subscribe is not found")
	}

	var follower database.Follower
	var subscriber database.Subscriber
	follower = database.Follower{
		UserID:       subTarget.ID,
		UserFollower: user.ID,
	}
	subscriber = database.Subscriber{
		UserID:         user.ID,
		UserSubscriber: subTarget.ID,
	}

	err = follower.Update(db)
	if err != nil {
		return errors.New("user_follower table update fail")
	}
	err = subscriber.Update(db)
	if err != nil {
		return errors.New("user_subscriber table update fail")
	}

	*reply = "Subscribe successfully."
	return nil
}

func (u *UserService) Unsubscribe(request [2]database.User, reply *string) error {
	db, err := database.SetupDefaultDatabase()
	if err != nil {
		return err
	}
	user, err := request[0].FindOne(db)
	unsubTarget, err := request[1].FindOne(db)
	if err != nil {
		return errors.New("the user that you want to unsubscribe is not found")
	}

	var unsubed database.Follower
	var unsuber database.Subscriber
	unsubed = database.Follower{
		UserID:       unsubTarget.ID,
		UserFollower: user.ID,
	}
	unsuber = database.Subscriber{
		UserID:         user.ID,
		UserSubscriber: unsubTarget.ID,
	}

	err = unsuber.Update(db)
	if err != nil {
		return errors.New("user_follower table update failed")
	}
	err = unsubed.Update(db)
	if err != nil {
		return errors.New("user_subscriber table update failed")
	}

	*reply = "Unsubscribe successfully."
	return nil
}
