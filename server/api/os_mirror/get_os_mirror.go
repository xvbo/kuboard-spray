package os_mirror

import (
	"net/http"

	"github.com/eip-work/kuboard-spray/common"
	"github.com/eip-work/kuboard-spray/constants"
	"github.com/gin-gonic/gin"
)

type GetMirrorRequest struct {
	Name string `uri:"cluster"`
}

func GetMirror(c *gin.Context) {
	var req GetMirrorRequest
	c.ShouldBindUri(&req)

	inventory, err := common.ParseYamlFile(MirrorInventoryPath(req.Name))
	if err != nil {
		common.HandleError(c, http.StatusInternalServerError, "cannnot parse file: "+MirrorInventoryPath(req.Name), err)
		return
	}

	statusFilePath := constants.GET_DATA_MIRROR_DIR() + "/" + req.Name + "/status.yaml"
	status, err := common.ParseYamlFile(statusFilePath)
	if err != nil {
		common.HandleError(c, http.StatusInternalServerError, "cannot pase file: "+statusFilePath, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"inventory": inventory,
			"status":    status,
		},
	})
}

func MirrorInventoryPath(name string) string {
	return constants.GET_DATA_MIRROR_DIR() + "/" + name + "/inventory.yaml"
}