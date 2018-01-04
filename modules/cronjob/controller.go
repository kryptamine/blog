package cronjob

import (
	"blog-api/modules/post"
	"github.com/gin-gonic/gin"
	"github.com/yanatan16/golang-instagram/instagram"
	"net/url"
	"strconv"
	"strings"
)

func index(c *gin.Context) {
	api := instagram.New("72411783c7ee431db5b58c7a4f830689",
		"aca6880d1eff470d83666741b7efdf51",
		"6846801018.7241178.fa276f1bc27b4c0d9888f234d335a4b6", true)

	if ok, err := api.VerifyCredentials(); !ok {
		panic(err)
	}

	params := url.Values{}

	params.Set("count", "1")

	if resp, err := api.GetUserRecentMedia("self", params); err == nil {
		for _, media := range resp.Medias {
			if id, err := strconv.Atoi(strings.Split(media.Id, "_")[0]); err == nil {
				c.MustGet("Repository").(*Repository).addInstagramPost(&post.InstagramPost{
					"",
					"",
					"",
					post.InstagramData{
						Id:     id,
						Images: media.Images,
						Link:   media.Link,
					},
				})
			}
		}
	}
}
