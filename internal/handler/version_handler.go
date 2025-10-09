package handler

import (
	"strconv"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"git.lulumia.fun/root/toge-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type VersionHandler struct {
	versionService domain.VersionService
}

func NewVersionHandler(versionService domain.VersionService) *VersionHandler {
	return &VersionHandler{versionService: versionService}
}

// Create CreateVersion godoc
// @Summary      创建版本号
// @Description  创建新的版本号
// @Tags         版本号
// @Accept       json
// @Produce      json
// @Param        version body domain.CreateVersionRequest true "Version信息"
// @Success      201  {object}  response.Response{data=model.Version}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /version/ [post]
func (h *VersionHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	version := &model.Version{

		Name: req.Name,
	}

	if err := h.versionService.Create(ctx, version); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Created(c, version)
}

// GetByID GetVersionByID godoc
// @Summary      获取版本号详情
// @Description  根据ID获取版本号详细信息
// @Tags         版本号
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "VersionID"
// @Success      200  {object}  response.Response{data=model.Version}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /version/{id} [get]
func (h *VersionHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	version, err := h.versionService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Version not found")
		return
	}

	response.Success(c, version)
}

// GetAll GetAllVersions godoc
// @Summary      获取所有版本号
// @Description  获取所有版本号列表（支持分页、排序和搜索）
// @Tags         版本号
// @Accept       json
// @Produce      json
// @Param        page       query     int    false  "页码，默认为1"  minimum(1)
// @Param        page_size  query     int    false  "每页大小，默认为10，最大100"  minimum(1) maximum(100)
// @Param        sort_by    query     string false  "排序字段：ID, Name, "
// @Param        sort_order query     string false  "排序方向：asc, desc，默认为desc"
// @Param        keyword    query     string false  "搜索关键词"
// @Param        search_by  query     string false  "搜索字段：Name, ，不指定则在所有字段中搜索"
// @Success      200  {object}  response.Response{data=pagination.PageResponse{data=[]model.Version}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /version/ [get]
func (h *VersionHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// 解析分页参数
	pageReq := pagination.ParsePageRequest(c)

	// 使用分页获取Version列表
	pageResponse, err := h.versionService.GetAllWithPagination(ctx, pageReq)
	if err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, pageResponse)
}

// Update UpdateVersion godoc
// @Summary      更新版本号
// @Description  根据ID更新版本号信息
// @Tags         versions
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "VersionID"
// @Param        version body domain.UpdateVersionRequest true "Version更新信息"
// @Success      200  {object}  response.Response{data=model.Version}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /version/{id} [put]
func (h *VersionHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req domain.UpdateVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	version, err := h.versionService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Version not found")
		return
	}

	if req.Name != nil {
		version.Name = *req.Name
	}

	if err := h.versionService.Update(ctx, version); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, version)
}

// Delete DeleteVersion godoc
// @Summary      删除版本号
// @Description  根据ID删除版本号
// @Tags         版本号
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "VersionID"
// @Success      200  {object}  response.Response{data=map[string]interface{}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /version/{id} [delete]
func (h *VersionHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.versionService.Delete(ctx, uint(id)); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Version deleted successfully"})
}
