package main

import (
	"log"
	"net/http"
	"workspace/src/helper"

	"github.com/gin-gonic/gin"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//get album
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//save data into array
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//get data by id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found !"})
}

//delete data by id
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "User is Deleted !"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found !"})
}

// var client *mongo.Client

func main() {
	//Init Router
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/save", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/delete/:id", deleteAlbum)

	// router.Run("localhost:8000")

	// router.GET("/api/books", getBooks())
	// router.GET("/api/books/{id}", getBook)
	// router.POST("/api/books", createBook)
	// router.PUT("/api/books/{id}", updateBook)
	// router.DELETE("/api/books/{id}", deleteBook)

	// router.Run(":8000", router)

	// configuration.Port
	// config := helper.GetConfiguration()
	// log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", router))

}
