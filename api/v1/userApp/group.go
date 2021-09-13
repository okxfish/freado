package userApp

type ApiGroup struct {
	AccountApi // 结构内嵌
	AuthApi
}

type RouterGroup struct {
}

func (receiver *RouterGroup) InitAccountApi() {

}
