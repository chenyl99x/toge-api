package handler

import (
	"strconv"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"git.lulumia.fun/root/toge-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ArtifactSetHandler struct {
	artifactSetService domain.ArtifactSetService
}

func NewArtifactSetHandler(artifactSetService domain.ArtifactSetService) *ArtifactSetHandler {
	return &ArtifactSetHandler{artifactSetService: artifactSetService}
}

// Create CreateArtifactSet godoc
// @Summary      创建圣遗物套装
// @Description  创建新的圣遗物套装
// @Tags         圣遗物套装
// @Accept       json
// @Produce      json
// @Param        artifactset body domain.CreateArtifactSetRequest true "ArtifactSet信息"
// @Success      201  {object}  response.Response{data=model.ArtifactSet}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact-set/ [post]
func (h *ArtifactSetHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateArtifactSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	artifactSet := &model.ArtifactSet{

		Name: req.Name,
	}

	if err := h.artifactSetService.Create(ctx, artifactSet); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Created(c, artifactSet)
}

// GetByID GetArtifactSetByID godoc
// @Summary      获取ArtifactSet详情
// @Description  根据ID获取ArtifactSet详细信息
// @Tags         圣遗物套装
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactSetID"
// @Success      200  {object}  response.Response{data=model.ArtifactSet}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /artifact-set/{id} [get]
func (h *ArtifactSetHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	artifactset, err := h.artifactSetService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "ArtifactSet not found")
		return
	}

	response.Success(c, artifactset)
}

// GetAll GetAllArtifactSets godoc
// @Summary      获取所有ArtifactSet
// @Description  获取所有ArtifactSet列表（支持分页、排序和搜索）
// @Tags         artifactsets
// @Accept       json
// @Produce      json
// @Param        page       query     int    false  "页码，默认为1"  minimum(1)
// @Param        page_size  query     int    false  "每页大小，默认为10，最大100"  minimum(1) maximum(100)
// @Param        sort_by    query     string false  "排序字段：ID, Name, "
// @Param        sort_order query     string false  "排序方向：asc, desc，默认为desc"
// @Param        keyword    query     string false  "搜索关键词"
// @Param        search_by  query     string false  "搜索字段：Name, ，不指定则在所有字段中搜索"
// @Success      200  {object}  response.Response{data=pagination.PageResponse{data=[]model.ArtifactSet}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact-set/ [get]
func (h *ArtifactSetHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// 解析分页参数
	pageReq := pagination.ParsePageRequest(c)

	// 使用分页获取ArtifactSet列表
	pageResponse, err := h.artifactSetService.GetAllWithPagination(ctx, pageReq)
	if err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, pageResponse)
}

// Update UpdateArtifactSet godoc
// @Summary      更新圣遗物套装
// @Description  根据ID更新圣遗物套装信息
// @Tags         圣遗物套装
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactSetID"
// @Param        artifactset body domain.UpdateArtifactSetRequest true "ArtifactSet更新信息"
// @Success      200  {object}  response.Response{data=model.ArtifactSet}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact-set/{id} [put]
func (h *ArtifactSetHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req domain.UpdateArtifactSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	artifactset, err := h.artifactSetService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "ArtifactSet not found")
		return
	}

	if req.Name != nil {
		artifactset.Name = *req.Name
	}

	if err := h.artifactSetService.Update(ctx, artifactset); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, artifactset)
}

// Delete DeleteArtifactSet godoc
// @Summary      删除圣遗物套装
// @Description  根据ID删除圣遗物套装
// @Tags         圣遗物套装
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactSetID"
// @Success      200  {object}  response.Response{data=map[string]interface{}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact-set/{id} [delete]
func (h *ArtifactSetHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.artifactSetService.Delete(ctx, uint(id)); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "ArtifactSet deleted successfully"})
}
