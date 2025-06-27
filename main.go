package main


// Import Packages
import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

// Create the prototype of Book
type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

// Book Data
var books=[]book{
	{
		ID:"1",
		Title:"In Search of Lost Time",
		Author:"Marcel Prostun",
		Quantity: 58,
	},
	{
		ID:"2",
		Title:"The Amazing Spiderman",
		Author:"Marvel Studio",
		Quantity: 35,
	},
	{
		ID:"3",
		Title:"The Dark Knight",
		Author:"Christopher Nolan",
		Quantity: 49,
	},
}

// Get Books
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

// Get books by ID
func bookById(c *gin.Context){
	id:=c.Param("id")
	book,err:=getBookById(id)

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK,book)
}


// Checkout the book
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id in query parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string)(*book,error){
	for i,b:=range books{
		if b.ID==id{
			return &books[i],nil
		}
	}
	return nil,errors.New("book not found")
}

// Created the Book
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


// Delete the book
func deleteBookById(c *gin.Context){
	id:=c.Param("id")
	index:=-1

	for i,b:=range books{
		if b.ID==id{
			index=i
			break
		}
	}
	if index==-1{
		c.JSON(http.StatusNotFound,gin.H{"message":"Book Not Found"})
	}

	books=append(books[:index],books[index+1:]...)
	c.JSON(http.StatusOK,gin.H{"message":"Book Deleted Successfully"})
}

func main(){
	router:=gin.Default()
	router.GET("/books",getBooks)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkout", checkoutBook)
	router.POST("/books",createBook)
	router.DELETE("/books/:id", deleteBookById)
	router.Run("localhost:8080")
}