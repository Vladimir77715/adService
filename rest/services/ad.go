package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Vladimir77715/adService/core/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	DESC = "DESC"
)

var (
	AdListQuery = func(key string, order string) string {
		var orderBY strings.Builder
		if key != "" {
			orderBY.WriteString("ORDER BY ")
			orderBY.WriteString("a." + key)
			if order == "" {
				order = DESC
			}
			orderBY.WriteString(" " + order)
		}

		return fmt.Sprintf("\tWITH ad_list AS (SELECT a.*, coalesce (ail.link,'') link,"+
			"ROW_NUMBER() OVER (PARTITION BY ail.ad_id order by ail.create_ts desc) rn "+
			"FROM ad a "+
			"LEFT JOIN ad_images_links ail  ON ail.ad_id = a.id "+
			"%s) "+
			"SELECT ad_list.\"name\",ad_list.\"link\", ad_list.price  FROM ad_list WHERE rn = 1", orderBY.String())
	}
	AdInsertQuery = func(ad database.Ad) string {
		return fmt.Sprintf("INSERT INTO public.ad (\"name\", description, price, create_ts)"+
			"VALUES('%s', '%s', '%d', now()) "+
			"RETURNING id", ad.Name, ad.Description, ad.Price)
	}
)

func RegisterAdService(apiGroup *gin.RouterGroup) {
	rg := apiGroup.Group("/ad")
	rg.GET("/adList", adList)
	rg.GET("/:id", adSingle)
	rg.POST("/", adAdd)
}

func adAdd(c *gin.Context) {
	var ad database.Ad
	if err := c.BindJSON(&ad); err != nil {
		writeError(c, err)
		return
	}
	var id int
	if err := database.Client.ExecuteQuery(AdInsertQuery(ad), func(row *sql.Rows) error {
		if err := row.Scan(&id); err != nil {
			return err
		}
		return nil
	}); err != nil {
		writeError(c, err)
		return
	}
	writeJson(c, struct {
		Id int `json:"id"`
	}{Id: id})
}

func adList(c *gin.Context) {
	list := make([]database.Ad, 0)
	if err := database.Client.ExecuteQuery(AdListQuery(c.Query("orderField"), c.Query("order")), func(row *sql.Rows) error {
		var imageLink string
		var ad database.Ad
		if e := row.Scan(&ad.Name, &imageLink, &ad.Price); e != nil {
			return e
		}
		if imageLink != "" {
			ad.ImageLinks = append(ad.ImageLinks, imageLink)
		}
		list = append(list, ad)
		return nil
	}); err != nil {
		writeError(c, err)
		return
	}
	if err := writeJson(c, &list); err != nil {
		writeError(c, err)
		return
	}

}

func adSingle(c *gin.Context) {
	var fields []string
	c.BindJSON(&fields)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		writeError(c, err)
		return
	}
	var args strings.Builder
	args.WriteString("a.name, a.price")
	var ad database.Ad
	var addDescription = false
	var uploadImages = false
	if fields != nil {
		for _, f := range fields {
			uploadImages = f == "images"
			if f == "description" {
				addDescription = true
				args.WriteString(", a.description")
			}
		}
	}
	query := fmt.Sprintf("SELECT %s FROM public.ad a WHERE a.id = %d", args.String(), id)
	if err = database.Client.ExecuteQuery(query, func(row *sql.Rows) error {
		if addDescription {
			if e := row.Scan(&ad.Name, &ad.Price, &ad.Description); e != nil {
				return e
			}
		} else {
			if e := row.Scan(&ad.Name, &ad.Price); e != nil {
				return e
			}
		}
		return nil
	}); err != nil {
		writeError(c, err)
		return
	}

	if uploadImages {
		query = fmt.Sprintf("SELECT al.link FROM ad_images_links al WHERE al.id = %d", id)
		if err = database.Client.ExecuteQuery(query, func(row *sql.Rows) error {
			var imageLink string
			if e := row.Scan(&imageLink); e != nil {
				return e
			}
			ad.ImageLinks = append(ad.ImageLinks, imageLink)
			return nil
		}); err != nil {
			writeError(c, err)
			return
		}
	}
	if err = writeJson(c, &ad); err != nil {
		writeError(c, err)
		return
	}
}

func writeJson(c *gin.Context, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.Writer.Write(b)
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return nil

}

func writeError(c *gin.Context, err error) {
	c.Status(http.StatusInternalServerError)
	c.Header("Content-Type", "application/json")
	c.Writer.WriteString(err.Error())

}
