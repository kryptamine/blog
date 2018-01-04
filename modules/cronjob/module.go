package cronjob

import (
	"blog-api/modules/base"
)

type Module struct{}

func (m Module) GetRoutes() []base.Route {
	return []base.Route{
		{"GET", "/cron-job/instagram-parser", index},
	}
}

func (m Module) GetRepository() interface{} {
	return &Repository{}
}
