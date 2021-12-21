package pcblog

import (
	"errors"
	"pcblog/database"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	var UserService UserService
	var LoginResponse LoginResponse
	var u database.User
	testTime := time.Now().UnixNano()
	// test register new account when repeated username
	u = database.User{Username: "admin"}
	assert.Equal(
		t, errors.New("this username has been taken. Please try a new one"),
		UserService.Register(u, &LoginResponse),
		"should equal")
	// test
	testNewName := "test" + strconv.FormatInt(testTime, 10)
	u = database.User{Username: testNewName[(len(testNewName) - 15):]}
	assert.Equal(
		t, nil,
		UserService.Register(u, &LoginResponse),
		"should equal")
}
