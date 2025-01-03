package rest

import (
	"gosql/modules/students/payload"
	util "gosql/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e endpoint) GetAllStudent(c *gin.Context) {

	// get query params
	cp, err := strconv.Atoi(c.Query("current_page"))
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}
	ps, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	// Bind
	payload := payload.ReqGetAllStudents{
		CurrentPage: cp,
		PageSize:    ps,
	}

	// service
	notvalid := util.Validates(c, payload)
	if notvalid {
		return
	}

	// call repository
	res, err := e.useCaseStudent.GetAllStudent(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}

func (e endpoint) GetStudent(c *gin.Context) {

	// Filtering URL
	id := c.Param("id")
	is, err := strconv.Atoi(id)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	// Fill Struct
	payload := payload.ReqGetDetail{
		ID: is,
	}

	// call repository
	res, err := e.useCaseStudent.GetDetail(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		panic(err)
	}

	util.ResponseOK(c, 200, res)

}
