package post

import (
	"blog-api/modules/base"
	"github.com/yanatan16/golang-instagram/instagram"
	"gopkg.in/mgo.v2/bson"
)

type BlogPost struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `json:"title" binding:"required,min=1,max=20"`
	Type  string        `json:"type"  binding:"required,min=1,max=20"`
}

type InstagramData struct {
	Id     int               `json:"id,omitempty"`
	Images *instagram.Images `json:"images,omitempty"`
	Link   string            `json:"link,omitempty"`
}

type InstagramPost struct {
	Id    bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title string        `json:"title" binding:"required,min=1,max=20"`
	Type  string        `json:"type"`
	Data  InstagramData `json:"data,omitempty"`
}

type Strategy interface {
	GetType() string
}

type StrategyX struct {
	Strategy `json:"post"`
}

type PostsResult struct {
	Posts      []StrategyX
	Pagination base.Pagination
}

func (c InstagramPost) GetType() string {
	return c.Type
}

func (c BlogPost) GetType() string {
	return c.Type
}

func getStructure(postType string) Strategy {
	switch postType {
	case "instagram":
		return &InstagramPost{}
		break
	case "blog":
		return &BlogPost{}
		break
	}

	return &BlogPost{}
}

func (s *StrategyX) SetBSON(raw bson.Raw) error {
	post := &BlogPost{}

	err := raw.Unmarshal(post)

	impl := getStructure(post.GetType())

	err = raw.Unmarshal(impl)

	if err != nil {
		return err
	}

	s.Strategy = impl

	return nil
}
