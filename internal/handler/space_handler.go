package handler

import (
	"strconv"

	"github.com/chenyl99x/toge-api/internal/domain"
	"github.com/chenyl99x/toge-api/internal/model"
	"github.com/chenyl99x/toge-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type SpaceHandler struct {
	spaceService domain.SpaceService
}

func NewSpaceHandler(spaceService domain.SpaceService) *SpaceHandler {
	return &SpaceHandler{
		spaceService: spaceService,
	}
}

// Create CreateSpace godoc
// @Summary 创建岛屿
// @Description 创建岛屿
// @Tags 岛屿
// @Accept json
// @Produce json
// @Param space body domain.CreateSpaceRequest true "创建岛屿请求"
// @Success 201 {object} response.Response{data=model.Space} "创建成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /space [post]
func (h *SpaceHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req domain.CreateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	space := &model.Space{
		Name:        req.Name,
		Description: req.Description,
		OwnerUserID: req.OwnerUserID,
		Type:        req.Type,
	}
	if err := h.spaceService.Create(ctx, space); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, space)
}

// GetByID GetSpaceByID godoc
// @Summary 获取岛屿详情
// @Description 获取岛屿详情
// @Tags 岛屿
// @Accept json
// @Produce json
// @Param id path uint true "岛屿ID"
// @Success 200 {object} response.Response{data=model.Space} "获取成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 404 {object} response.Response "空间不存在"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /space/{id} [get]
func (h *SpaceHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	space, err := h.spaceService.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, space)
}
