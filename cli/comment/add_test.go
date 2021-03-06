package comment

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
	. "github.com/z-t-y/flogo/cli/post"
	"github.com/z-t-y/flogo/utils"
)

func TestAddComment(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	content := "Flogo Comment Unit Test - " + strconv.FormatInt(time.Now().UnixNano(), 10)
	posts, err := GetPostsFrom("/api/v3/post/all", accessToken)
	if err != nil {
		t.Error(err)
	}
	var randomPost utils.Post
	for _, post := range posts {
		if !post.Private {
			randomPost = post
			break
		}
	}
	comment, err := AddComment(accessToken, content, randomPost.ID, 0)
	if err != nil {
		t.Error(err)
	}
	if comment.Body != content {
		t.Errorf("TestAddComment: expected comment body %v, actual %v", content, comment.Body)
	}
}
