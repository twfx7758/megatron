package routers

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

//User 用户model类型
type User struct {
	Name   string
	Mobile string
}

//UserController 用户restapi
type UserController struct {
}

//Register 注册用户的路由
func (u *UserController) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.POST("add").To(u.createUser))

	container.Add(ws)
}

func (u *UserController) findUser(request *restful.Request, response *restful.Response) {
	//id := request.PathParameter("user-id")
	//usr, ok := nil, false
	var usr User
	ok := false

	if !ok {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

func (u *UserController) createUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(usr)
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}
