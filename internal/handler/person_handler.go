package handler

import (
	"strconv"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"git.lulumia.fun/root/toge-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	personService domain.PersonService
}

func NewPersonHandler(personService domain.PersonService) *PersonHandler {
	return &PersonHandler{personService: personService}
}

// Create CreatePerson godoc
// @Summary      创建Person
// @Description  创建新的Person
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        person body domain.CreatePersonRequest true "Person信息"
// @Success      201  {object}  response.Response{data=model.Person}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /persons/ [post]
func (h *PersonHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	person := &model.Person{

		Name: req.Name,

		Age: req.Age,

		Gender: req.Gender,

		Email: req.Email,

		Phone: req.Phone,

		Address: req.Address,

		Company: req.Company,

		Position: req.Position,

		Status: req.Status,
	}

	if err := h.personService.Create(ctx, person); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Created(c, person)
}

// GetByID GetPersonByID godoc
// @Summary      获取Person详情
// @Description  根据ID获取Person详细信息
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "PersonID"
// @Success      200  {object}  response.Response{data=model.Person}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /persons/{id} [get]
func (h *PersonHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	person, err := h.personService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Person not found")
		return
	}

	response.Success(c, person)
}

// GetAll GetAllPersons godoc
// @Summary      获取所有Person
// @Description  获取所有Person列表（支持分页、排序和搜索）
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        page       query     int    false  "页码，默认为1"  minimum(1)
// @Param        page_size  query     int    false  "每页大小，默认为10，最大100"  minimum(1) maximum(100)
// @Param        sort_by    query     string false  "排序字段：ID, Name, Status, "
// @Param        sort_order query     string false  "排序方向：asc, desc，默认为desc"
// @Param        keyword    query     string false  "搜索关键词"
// @Param        search_by  query     string false  "搜索字段：Name, Email, ，不指定则在所有字段中搜索"
// @Success      200  {object}  response.Response{data=pagination.PageResponse{data=[]model.Person}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /persons/ [get]
func (h *PersonHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// 解析分页参数
	pageReq := pagination.ParsePageRequest(c)

	// 使用分页获取Person列表
	pageResponse, err := h.personService.GetAllWithPagination(ctx, pageReq)
	if err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, pageResponse)
}

// Update UpdatePerson godoc
// @Summary      更新Person
// @Description  根据ID更新Person信息
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "PersonID"
// @Param        person body domain.UpdatePersonRequest true "Person更新信息"
// @Success      200  {object}  response.Response{data=model.Person}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /persons/{id} [put]
func (h *PersonHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req domain.UpdatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	person, err := h.personService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Person not found")
		return
	}

	if req.Name != nil {
		person.Name = *req.Name
	}

	if req.Age != nil {
		person.Age = *req.Age
	}

	if req.Gender != nil {
		person.Gender = *req.Gender
	}

	if req.Email != nil {
		person.Email = *req.Email
	}

	if req.Phone != nil {
		person.Phone = *req.Phone
	}

	if req.Address != nil {
		person.Address = *req.Address
	}

	if req.Company != nil {
		person.Company = *req.Company
	}

	if req.Position != nil {
		person.Position = *req.Position
	}

	if req.Status != nil {
		person.Status = *req.Status
	}

	if err := h.personService.Update(ctx, person); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, person)
}

// Delete DeletePerson godoc
// @Summary      删除Person
// @Description  根据ID删除Person
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "PersonID"
// @Success      200  {object}  response.Response{data=map[string]interface{}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /persons/{id} [delete]
func (h *PersonHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.personService.Delete(ctx, uint(id)); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Person deleted successfully"})
}
