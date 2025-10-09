package handler

import (
	"strconv"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"git.lulumia.fun/root/toge-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type NationHandler struct {
	nationService domain.NationService
}

func NewNationHandler(nationService domain.NationService) *NationHandler {
	return &NationHandler{nationService: nationService}
}

// Create CreateNation godoc
// @Summary      创建Nation
// @Description  创建新的Nation
// @Tags         nations
// @Accept       json
// @Produce      json
// @Param        nation body domain.CreateNationRequest true "Nation信息"
// @Success      201  {object}  response.Response{data=model.Nation}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nation/ [post]
func (h *NationHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateNationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	nation := &model.Nation{

		Name: req.Name,
	}

	if err := h.nationService.Create(ctx, nation); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Created(c, nation)
}

// GetByID GetNationByID godoc
// @Summary      获取Nation详情
// @Description  根据ID获取Nation详细信息
// @Tags         nations
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "NationID"
// @Success      200  {object}  response.Response{data=model.Nation}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /nation/{id} [get]
func (h *NationHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	nation, err := h.nationService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Nation not found")
		return
	}

	response.Success(c, nation)
}

// GetAll GetAllNations godoc
// @Summary      获取所有Nation
// @Description  获取所有Nation列表（支持分页、排序和搜索）
// @Tags         nations
// @Accept       json
// @Produce      json
// @Param        page       query     int    false  "页码，默认为1"  minimum(1)
// @Param        page_size  query     int    false  "每页大小，默认为10，最大100"  minimum(1) maximum(100)
// @Param        sort_by    query     string false  "排序字段：ID, Name, "
// @Param        sort_order query     string false  "排序方向：asc, desc，默认为desc"
// @Param        keyword    query     string false  "搜索关键词"
// @Param        search_by  query     string false  "搜索字段：Name, ，不指定则在所有字段中搜索"
// @Success      200  {object}  response.Response{data=pagination.PageResponse{data=[]model.Nation}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nation/ [get]
func (h *NationHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// 解析分页参数
	pageReq := pagination.ParsePageRequest(c)

	// 使用分页获取Nation列表
	pageResponse, err := h.nationService.GetAllWithPagination(ctx, pageReq)
	if err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, pageResponse)
}

// Update UpdateNation godoc
// @Summary      更新Nation
// @Description  根据ID更新Nation信息
// @Tags         nations
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "NationID"
// @Param        nation body domain.UpdateNationRequest true "Nation更新信息"
// @Success      200  {object}  response.Response{data=model.Nation}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nation/{id} [put]
func (h *NationHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req domain.UpdateNationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	nation, err := h.nationService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Nation not found")
		return
	}

	if req.Name != nil {
		nation.Name = *req.Name
	}

	if err := h.nationService.Update(ctx, nation); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, nation)
}

// Delete DeleteNation godoc
// @Summary      删除Nation
// @Description  根据ID删除Nation
// @Tags         nations
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "NationID"
// @Success      200  {object}  response.Response{data=map[string]interface{}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nation/{id} [delete]
func (h *NationHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.nationService.Delete(ctx, uint(id)); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Nation deleted successfully"})
}
