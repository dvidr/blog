package actions

import (
	"github.com/dvidr/blog/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

func CommentsCreatePost(c buffalo.Context) error {
	comment := &models.Comment{}
	user := c.Value("current_user").(*models.User)
	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)
	comment.AuthorID = user.ID
	postID, err := uuid.FromString(c.Param("pid"))
	if err != nil {
		return errors.WithStack(err)
	}
	comment.PostID = postID
	verrs, err := tx.ValidateAndCreate(comment)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Flash().Add("danger", "There was an error adding your comment.")
		return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
	}
	c.Flash().Add("success", "Comment added successfully.")
	return c.Redirect(302, "/posts/detail/%s", c.Param("pid"))
}

// CommentsEdit default implementation.
func CommentsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("comments/edit.html"))
}

// CommentsDelete default implementation.
func CommentsDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("comments/delete.html"))
}
