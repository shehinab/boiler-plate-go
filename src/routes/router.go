package route

import (
	"github.com/gin-gonic/gin"

	user "gitlab.com/abhishek.k8/crud/src/controller/user"
)

type (
	//IRouter default router
	IRouter interface {
		GetRoutes(appGroup *gin.RouterGroup)
		// InitSubRoutes(appGroup *gin.RouterGroup)
	}
	//MainRouter default MainRouter
	MainRouter struct{}
)

//GetRoutes -
func (mc *MainRouter) GetRoutes(appGroup *gin.RouterGroup) {
	mc.InitSubRoutes(appGroup)
}

//InitSubRoutes -
func (mc *MainRouter) InitSubRoutes(appGroup *gin.RouterGroup) {
	controllerMap := makeControllerMap()
	for key, cntrlr := range controllerMap {
		cntrlr.GetRoutes(appGroup.Group(key))
	}
}

func makeControllerMap() map[string]IRouter {
	controllerMap := make(map[string]IRouter)

	controllerMap["user"] = &user.UserController{}
	return controllerMap
}
