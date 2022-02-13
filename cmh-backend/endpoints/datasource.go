package endpoints

import (
	"cmh-backend/cmhtypes"
	"cmh-backend/datastream"
	"cmh-backend/model"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

var allowedTypes = [...]string{
	"number",
	"log",
	"state",
}

func isValidType(sType string) bool {
	for _, supportedType := range allowedTypes {
		if sType == supportedType {
			return true
		}
	}
	return false
}

func AddSource(c *gin.Context) {
	var src cmhtypes.Source
	bindErr := c.BindJSON(&src)
	if bindErr != nil {
		c.JSON(400, gin.H{
			"error": "Required fields are missing",
		})
		return
	}

	srcTypeExists := model.CheckCollisionInSourceType("test", src.SourceTypeName)
	if srcTypeExists {
		srcExists := model.CheckCollisionInSource("test", src.Name)
		if srcExists {
			c.JSON(406, gin.H{
				"error": "Source \"" + src.Name + "\" already exists",
			})
		} else {
			added := model.AddSource("test", src)
			if added {
				srcType, err := model.FetchSpecificSourceType("test", src.Name)
				if err == nil {
					datastream.SourceMap[src.Name] = srcType
					datastream.SpawnStreamListener(src.Name)
					c.JSON(200, "")
				} else {
					log.Println("Error occured: ", err)
					c.JSON(500, gin.H{
						"error": "Internal Server Error. Please contact administrator.",
					})
				}
			} else {
				c.JSON(500, gin.H{
					"error": "Internal Server Error. Please contact administrator.",
				})
			}
		}
	} else {
		c.JSON(406, gin.H{
			"error": "Source type \"" + src.SourceTypeName + "\" does not exist. Please create one from Source Type tab.",
		})
	}
}

func EditSource(c *gin.Context) {

}

func DelSource(c *gin.Context) {

}

func FetchList(c *gin.Context) {
	var fetcher cmhtypes.Fetcher
	query := c.Request.URL.Query()
	shouldHaveKeys := []string{"id", "listName"}
	for _, key := range shouldHaveKeys {
		if !query.Has(key) {
			c.JSON(400, gin.H{
				"error": "Required fields are missing",
			})
			return
		}
	}
	fetcher.Id = query.Get("id")
	fetcher.ListName = query.Get("listName")

	//Fetch the list
	list := model.FetchList(fetcher.Id, strings.ToLower(fetcher.ListName))
	if len(list) > 0 && list[0]["error"] == true {
		c.JSON(400, gin.H{
			"error": true,
		})
		return
	}

	c.JSON(200, list)
}

func AddSourceType(c *gin.Context) {
	var srcType cmhtypes.SourceType
	bindErr := c.BindJSON(&srcType)
	if bindErr != nil {
		c.JSON(400, gin.H{
			"error": "Required fields are missing",
		})
		return
	}

	//Check for name collision
	if model.CheckCollisionInSourceType("test", srcType.Name) {
		c.JSON(406, gin.H{
			"error": srcType.Name + " already exists",
		})
		return
	}

	//Check for allowed types
	srcTypeStats := srcType.Stats
	var notSupportedKeys []string
	notSupportedKeyFound := false
	for key, value := range srcTypeStats {
		if !isValidType(value) {
			notSupportedKeys = append(notSupportedKeys, key)
			notSupportedKeyFound = true
		}
	}

	if notSupportedKeyFound {
		errorMsg := fmt.Sprintf("The following keys were supplied with unsupported types: %s", strings.Join(notSupportedKeys, ","))
		c.JSON(400, gin.H{
			"error": errorMsg,
		})
		return
	}

	//Insert type in database
	added := model.AddSourceType("test", srcType)
	if added {
		c.JSON(200, "")
	} else {
		c.JSON(500, gin.H{
			"error": "Internal Server Error. Please contact administrator",
		})
	}
}

func EditSourceType(c *gin.Context) {

}

func DelSourceType(c *gin.Context) {

}

func InitializeDataSourceEndPoints(engine *gin.Engine) {
	engine.POST("/addsource", AddSource)
	engine.POST("/addsourcetype", AddSourceType)
	engine.GET("/fetchlist", FetchList)
}
