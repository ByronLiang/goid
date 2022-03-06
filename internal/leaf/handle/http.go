package handle

import (
	"net/http"
	"strconv"

	"github.com/ByronLiang/goid/pkg/protocol/response"
	"github.com/ByronLiang/goid/pkg/service"
	"github.com/gin-gonic/gin"
)

func QueryLeaf(ctx *gin.Context) {
	domainId, _ := strconv.Atoi(ctx.Param("domain"))
	number, err := service.Leaf.Get(int64(domainId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(response.CodeInternalError, err.Error()))
	}
	ctx.JSON(http.StatusOK, response.Success(number))
	return
}
