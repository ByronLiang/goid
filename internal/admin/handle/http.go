package handle

import (
	"net/http"
	"strconv"

	"github.com/ByronLiang/goid/pkg/utils"

	"github.com/ByronLiang/goid/pkg/db"

	"github.com/ByronLiang/goid/pkg/model"
	"github.com/ByronLiang/goid/pkg/protocol/response"
	"github.com/gin-gonic/gin"
)

func AddLeaf(ctx *gin.Context) {
	var leaf model.Leaf
	err := ctx.BindJSON(&leaf)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(response.CodeInvalidRequestParams, err.Error()))
		return
	}
	if err := db.LeafDao.Create(&leaf); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(response.CodeDBError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(leaf))
	return
}

func GetLeaf(ctx *gin.Context) {
	filters := make([]db.FilterFunc, 0, 2)
	status, err := strconv.Atoi(ctx.DefaultQuery("status", "0"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(response.CodeInvalidRequestParams, err.Error()))
		return
	}
	if status > 0 {
		filters = append(filters, db.FilterStatus(status))
	}
	domainIdStr := ctx.Query("domain_id")
	if domainIdStr != "" {
		domainIds := utils.SplitParseInt64(domainIdStr)
		if len(domainIds) > 0 {
			filters = append(filters, db.FilterDomainId(domainIds...))
		}
	}
	leafs, err := db.LeafDao.GetLeaf(filters...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(response.CodeDBError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(leafs))
	return
}
