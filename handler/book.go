package handler

// delivery, controller

import (
	"strconv"

	"github.com/ainmtsn1999/orm-book-api-deploy/helper"
	"github.com/ainmtsn1999/orm-book-api-deploy/model"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.GormBook{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllBook(c *gin.Context) {

	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetBookById(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.GormBook{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in.Id = idInt
	// call service
	res, err := h.app.UpdateBook(in)
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, gin.H{"message": "Book deleted successfully"})
}
