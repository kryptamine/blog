package post

import (
	"blog-api/modules/base"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math"
)

const CollectionName = "posts"

type Repository struct {
	Collection *mgo.Collection
}

func (r Repository) all(curPage int) PostsResult {
	var transforms []StrategyX

	find := r.Collection.Find(nil)

	skip := (curPage - 1) * base.PerPage

	count, err := find.Count()

	if err != nil {
		count = 0
	}

	if err := find.Limit(base.PerPage).Skip(skip).All(&transforms); err != nil {
		fmt.Println("Failed to get results:", err)
	}

	return PostsResult{
		transforms,
		base.Pagination{
			TotalCount: count,
			CurPage:    curPage,
			Page:       math.Ceil(float64(count) / float64(base.PerPage)),
			PageCount:  len(transforms),
		},
	}
}

func (r Repository) get(id string) (error, BlogPost) {
	result := BlogPost{}

	if !bson.IsObjectIdHex(id) {
		return errors.New("the object is not hexId"), result
	}

	if err := r.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		fmt.Println("Failed to write results:", err)

		return errors.New("item is not found"), result
	}

	return nil, result
}

func (r Repository) add(post *BlogPost) bool {
	post.Id = bson.NewObjectId()

	err := r.Collection.Insert(post)

	return err == nil
}

func (r Repository) delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("the object is not hexId")
	}

	oid := bson.ObjectIdHex(id)

	if err := r.Collection.RemoveId(oid); err != nil {
		return errors.New("object is not exist")
	}

	return nil
}
