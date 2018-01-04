package cronjob

import (
	"blog-api/modules/base"
	"blog-api/modules/post"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func (r Repository) addInstagramPost(iPost *post.InstagramPost) bool {
	iPost.Id = bson.NewObjectId()
	iPost.Type = "instagram"

	err := base.DB.C(post.CollectionName).Insert(iPost)

	return err == nil
}
