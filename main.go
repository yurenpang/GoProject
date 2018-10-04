package main

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

// Group contains information about a group of people
type Group struct {
  ID     int     `json:"id" binding:"required"`
  Response  int     `json:"response"`
  Info   string  `json:"info" binding:"required"`
}

// We'll create a list of jokes
var groups = []Group{
  Group{1, 0, "Tianyou; Looking for one roomate; More information should be stored in more variables."},
  Group{2, 0, "AAA; Looking for two;"},
  Group{3, 0, "BBB; Looking for four;"},
}

func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()

  // Serve frontend static files
  router.Use(static.Serve("/", static.LocalFile("./views", true)))

  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }

  api.GET("/groups", GroupHandler)
  api.POST("/groups/respond/:groupID", RespondGroup)

  // Start and run the server
  router.Run(":8080")
}

// GroupHandler retrieves a list of avaliable groups
func GroupHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, groups)
}

// RespondGroup sends a response to a particular group
func RespondGroup(c *gin.Context) {
  // confirm Group ID sent is valid
  // remember to import the `strconv` package
  if groupid, err := strconv.Atoi(c.Param("groupID")); err == nil {
    for i := 0; i < len(groups); i++ {
      if groups[i].ID == groupid {
        groups[i].Response += 1
      }
    }

    // return a pointer to the updated jokes list
    c.JSON(http.StatusOK, &groups)
  } else {
    // Joke ID is invalid
    c.AbortWithStatus(http.StatusNotFound)
  }
}