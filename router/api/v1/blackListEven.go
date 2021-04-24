package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"qq_black_user_list/models"
	"qq_black_user_list/pkg/e"
)

//func CreateEven(c *gin.Context) {}
func GetEvenForQq(c *gin.Context) {
	//fmt.Println(com.StrTo(c.Param("qq_num")).String())
	even := models.Even{
		QqNum: c.Param("qq_num"),
	}
	valid := validation.Validation{}
	valid.Required(even.QqNum, "qq_num")
	if valid.HasErrors() {
		e.ResponseForJson(c, nil, e.ErrInput)
	}
	if err, isExist := even.QqNumIsExist(); err != nil {
		fmt.Println(err.Error())
		e.ResponseForJson(c, nil, e.ErrDB)
		return
	} else if isExist {
		if err := even.GetEvenAllDataForQq(); err != nil {
			e.ResponseForJson(c, nil, e.ErrDB)
			return
		}
	} else {
		e.ResponseForJson(c, nil, e.ErrNoQqInList)
		return
	}

	//执行e.ResponseForJson 正确返回
	e.ResponseForJson(c, gin.H{
		"qq_num":                 even.QqNum,
		"even_content":           even.EvenContent,
		"even_content_more_info": even.EvenContentMoreInfo,
		"verify_status":          even.VerifyStatus,
	}, nil)
	return
}

//func GetEvenList(c *gin.Context) {}
//func DeleteEven(c *gin.Context) {}
//func SetEvenAuditStatus(c *gin.Context) {}
