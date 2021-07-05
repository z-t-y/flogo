package comment

import (
	"fmt"
	"github.com/z-t-y/flogo/cli/auth"
	. "github.com/z-t-y/flogo/cli/post"
	"strconv"
	"testing"
	"time"
)

func TestGetComment(t *testing.T) {
	t.Parallel()
	accessToken, err := auth.GetAccessToken(username, password)
	title := "Flog Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := UploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	content = "Flog Comment Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	comment, err := addComment(accessToken, content, post.ID, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(comment.ID)
	commentReturned, err := getComment(accessToken, comment.ID)
	if err != nil {
		t.Error(err)
	}
	if commentReturned.ID != comment.ID {
		t.Error("TestGetComment: expect comment to equal commentReturned")
	}
}