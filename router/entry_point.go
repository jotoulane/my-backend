package router

type routerPackage struct {
	UserRouter userRouter
	ChatRouter chatRouter
}

var RouterPackageApp = new(routerPackage)
