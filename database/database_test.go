package database

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInbox(t *testing.T) {
	// test AddInbox
	//var AddInbox = Inbox{
	//	ArticleID:      0,
	//	UserID:         0,
	//	AuthorUserID:   0,
	//	ArticleTitle:   "",
	//	ArticleContent: "",
	//	ArticleDate:    0,
	//}
}

func TestArticle(t *testing.T) {
	testTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	testArticleName := "article" + testTime[(len(testTime)-5):]
	var a = Article{
		UserID:         1,
		ArticleTitle:   testArticleName,
		ArticleContent: "This is a testify test.",
		ArticleDate:    0,
	}
	DB, _ := SetupDefaultDatabase()
	// test Add
	//assert.Equal(t, nil, a.Add(DB), "should equal")
	//test Find
	var u = User{
		ID: 1,
	}
	res := Article{
		ArticleID:      1,
		UserID:         1,
		ArticleTitle:   "article67700",
		ArticleContent: "This is a testify test.",
		ArticleDate:    0,
	}
	testRes, _ := a.Find(DB, u)
	assert.Equal(
		t,
		res,
		testRes[0],
		"should equal",
	)
}
