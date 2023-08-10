package handlers

import (
	"book-library/api/types"
	"book-library/api/utils"
	"net/http"

	"github.com/labstack/echo"
)

var book =  &[]types.Book{
	{Id:"1",Title: "Learning Go",Author: "Muma",Qauntity: 2},
	{Id:"2",Title: "Learning PHP",Author: "Musa",Qauntity: 2},
}

//GetALLBooks ... Get all books
// @Summary Get all books
// @Description get all books
// @Tags Book
// @Success 200 {array} model.Book
// @Failure 404 {object} object
// @Router / [get]
func GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK,&book)
}

//GetBookById ... Get book by id
// @Summary Get single book by id
// @Description get single book
// @Tags Book
// @Success 200 {object} model.Book
// @Failure 404 {object} object
// @Router / [get]
func GetBookById(c echo.Context) error{
	id := c.Param("id")
	book,err := utils.GetById(id,*book)

	if err != nil {
		return c.String(http.StatusBadRequest,"Bad Request")
	}

	return c.JSON(http.StatusCreated,book)
}

//StoreBooks ... inser book 
// @Summary insert single book 
// @Description store book 
// @Tags Book
// @Success 200 {object} model.Book
// @Failure 404 {object} object
// @Router / [post]
func StoreBooks(c echo.Context) error {
	var newBook types.Book
	err := c.Bind(&newBook)

	if err != nil {
		return c.String(http.StatusBadRequest,"Bad Request")
	}

	*book = append(*book,newBook)

	return c.JSON(http.StatusCreated,newBook)
}

//CheckoutBook ... checkout a book
// @Summary checkout book from library
// @Description checkout book  
// @Tags Book
// @Success 200 {object} model.Book
// @Failure 404 {object} object
// @Router / [patch]
func CheckoutBook (c echo.Context) error {
	id := c.QueryParams().Get("id")

	book,err := utils.GetById(id,*book)

	if err != nil {
		return c.String(http.StatusBadRequest,"Bad Request")
	}

	if book.Qauntity <= 0{
		return c.String(http.StatusBadRequest,"Book not available")
	}

	book.Qauntity -= 1

	return c.JSON(http.StatusCreated,book)
}

//CheckinBook ... checkin a book
// @Summary checkin book to library
// @Description checkin book  
// @Tags Book
// @Success 200 {object} model.Book
// @Failure 404 {object} object
// @Router / [patch]
func CheckinBook (c echo.Context) error{
	id := c.QueryParams().Get("id")

	book,err := utils.GetById(id,*book)

	if err != nil {
		return c.String(http.StatusBadRequest,"Bad Request")
	}

	book.Qauntity += 1

	return c.JSON(http.StatusCreated,book)
}

//DeleteBook ... delete a book
// @Summary delete from library
// @Description delete book  
// @Tags Book
// @Success 200 {object} object
// @Failure 404 {object} object
// @Router / [delete]
func DeleteBook (c echo.Context) error {
	id := c.Param("id")

	index,err := utils.GetItemIndex(id,*book)

	if err != nil {
		return c.String(http.StatusBadRequest,"Bad Request")
	}
	newBookList := *book
	newBookList  = append(newBookList[:int(index)],newBookList[index+1:]...)
	*book = newBookList

	return c.JSON(http.StatusOK,"Deleted Succesfully")

}