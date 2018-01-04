package post

import (
	"blog-api/modules/base"
)

type Module struct{}

func (m Module) GetRoutes() []base.Route {
	return []base.Route{
		{"GET", "/posts", index},
		{"GET", "/posts/:id", show},
		{"POST", "/posts", create},
		{"DELETE", "/posts/:id", remove},
	}
}

func (m Module) GetRepository() interface{} {
	return &Repository{
		Collection: base.DB.C(CollectionName),
	}
}
