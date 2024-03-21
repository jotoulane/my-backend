package api

type apiPackage struct {
	UserApi
	ChatApi
}

var ApiPackageApp = new(apiPackage)
