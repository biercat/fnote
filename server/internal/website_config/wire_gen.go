// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package website_config

import (
	"github.com/chenmingyong0423/fnote/server/internal/website_config/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/website_config/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/website_config/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/website_config/internal/web"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitWebsiteConfigModule(mongoDB *mongo.Database) Model {
	websiteConfigDao := dao.NewWebsiteConfigDao(mongoDB)
	websiteConfigRepository := repository.NewWebsiteConfigRepository(websiteConfigDao)
	websiteConfigService := service.NewWebsiteConfigService(websiteConfigRepository)
	websiteConfigHandler := web.NewWebsiteConfigHandler(websiteConfigService)
	model := Model{
		Svc: websiteConfigService,
		Hdl: websiteConfigHandler,
	}
	return model
}

// wire.go:

var ConfigProviders = wire.NewSet(web.NewWebsiteConfigHandler, service.NewWebsiteConfigService, repository.NewWebsiteConfigRepository, dao.NewWebsiteConfigDao, wire.Bind(new(service.IWebsiteConfigService), new(*service.WebsiteConfigService)), wire.Bind(new(repository.IWebsiteConfigRepository), new(*repository.WebsiteConfigRepository)), wire.Bind(new(dao.IWebsiteConfigDao), new(*dao.WebsiteConfigDao)))