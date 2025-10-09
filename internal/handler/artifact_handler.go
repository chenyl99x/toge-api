package handler

import (
	"strconv"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"git.lulumia.fun/root/toge-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ArtifactHandler struct {
	artifactService domain.ArtifactService
}

func NewArtifactHandler(artifactService domain.ArtifactService) *ArtifactHandler {
	return &ArtifactHandler{artifactService: artifactService}
}

// Create CreateArtifact godoc
// @Summary      创建圣遗物
// @Description  创建新的圣遗物
// @Tags         圣遗物
// @Accept       json
// @Produce      json
// @Param        artifact body domain.CreateArtifactRequest true "圣遗物信息"
// @Success      201  {object}  response.Response{data=model.Artifact}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact/ [post]
func (h *ArtifactHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateArtifactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	artifact := &model.Artifact{

		ArtifactSetID: req.ArtifactSetID,

		Name: req.Name,

		Type: req.Type,

		Description: req.Description,

		Story: req.Story,
	}

	if err := h.artifactService.Create(ctx, artifact); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Created(c, artifact)
}

// GetByID GetArtifactByID godoc
// @Summary      获取圣遗物详情
// @Description  根据ID获取圣遗物详细信息
// @Tags         圣遗物
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactID"
// @Success      200  {object}  response.Response{data=model.Artifact}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /artifact/{id} [get]
func (h *ArtifactHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	artifact, err := h.artifactService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Artifact not found")
		return
	}

	response.Success(c, artifact)
}

// GetAll GetAllArtifacts godoc
// @Summary      获取所有圣遗物
// @Description  获取所有圣遗物列表（支持分页、排序和搜索）
// @Tags         圣遗物
// @Accept       json
// @Produce      json
// @Param        page       query     int    false  "页码，默认为1"  minimum(1)
// @Param        page_size  query     int    false  "每页大小，默认为10，最大100"  minimum(1) maximum(100)
// @Param        sort_by    query     string false  "排序字段：ID, ArtifactSetID, Name, "
// @Param        sort_order query     string false  "排序方向：asc, desc，默认为desc"
// @Param        keyword    query     string false  "搜索关键词"
// @Param        search_by  query     string false  "搜索字段：Name, Description, ，不指定则在所有字段中搜索"
// @Success      200  {object}  response.Response{data=pagination.PageResponse{data=[]model.Artifact}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact/ [get]
func (h *ArtifactHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// 解析分页参数
	pageReq := pagination.ParsePageRequest(c)

	// 使用分页获取Artifact列表
	pageResponse, err := h.artifactService.GetAllWithPagination(ctx, pageReq)
	if err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, pageResponse)
}

// Update UpdateArtifact godoc
// @Summary      更新圣遗物
// @Description  根据ID更新圣遗物信息
// @Tags         圣遗物
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactID"
// @Param        artifact body domain.UpdateArtifactRequest true "Artifact更新信息"
// @Success      200  {object}  response.Response{data=model.Artifact}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact/{id} [put]
func (h *ArtifactHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req domain.UpdateArtifactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	artifact, err := h.artifactService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, "Artifact not found")
		return
	}

	if req.ArtifactSetID != nil {
		artifact.ArtifactSetID = *req.ArtifactSetID
	}

	if req.Name != nil {
		artifact.Name = *req.Name
	}

	if req.Type != nil {
		artifact.Type = *req.Type
	}

	if req.Description != nil {
		artifact.Description = *req.Description
	}

	if req.Story != nil {
		artifact.Story = *req.Story
	}

	if err := h.artifactService.Update(ctx, artifact); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, artifact)
}

// Delete DeleteArtifact godoc
// @Summary      删除圣遗物
// @Description  根据ID删除圣遗物
// @Tags         圣遗物
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ArtifactID"
// @Success      200  {object}  response.Response{data=map[string]interface{}}
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /artifact/{id} [delete]
func (h *ArtifactHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.artifactService.Delete(ctx, uint(id)); err != nil {
		response.DatabaseError(c, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Artifact deleted successfully"})
}
