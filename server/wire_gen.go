// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post"
	handler7 "github.com/chenmingyong0423/fnote/server/internal/backup/handler"
	service9 "github.com/chenmingyong0423/fnote/server/internal/backup/service"
	handler2 "github.com/chenmingyong0423/fnote/server/internal/category/handler"
	repository2 "github.com/chenmingyong0423/fnote/server/internal/category/repository"
	dao2 "github.com/chenmingyong0423/fnote/server/internal/category/repository/dao"
	service3 "github.com/chenmingyong0423/fnote/server/internal/category/service"
	"github.com/chenmingyong0423/fnote/server/internal/comment"
	handler6 "github.com/chenmingyong0423/fnote/server/internal/count_stats/handler"
	repository3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository"
	dao3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository/dao"
	service2 "github.com/chenmingyong0423/fnote/server/internal/count_stats/service"
	"github.com/chenmingyong0423/fnote/server/internal/data_analysis"
	service4 "github.com/chenmingyong0423/fnote/server/internal/email/service"
	"github.com/chenmingyong0423/fnote/server/internal/file/handler"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/file/service"
	"github.com/chenmingyong0423/fnote/server/internal/friend"
	"github.com/chenmingyong0423/fnote/server/internal/global"
	"github.com/chenmingyong0423/fnote/server/internal/ioc"
	service6 "github.com/chenmingyong0423/fnote/server/internal/message/service"
	handler4 "github.com/chenmingyong0423/fnote/server/internal/message_template/handler"
	repository4 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository"
	dao4 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository/dao"
	service5 "github.com/chenmingyong0423/fnote/server/internal/message_template/service"
	"github.com/chenmingyong0423/fnote/server/internal/post"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	"github.com/chenmingyong0423/fnote/server/internal/post_visit"
	handler5 "github.com/chenmingyong0423/fnote/server/internal/tag/handler"
	repository6 "github.com/chenmingyong0423/fnote/server/internal/tag/repository"
	dao6 "github.com/chenmingyong0423/fnote/server/internal/tag/repository/dao"
	service8 "github.com/chenmingyong0423/fnote/server/internal/tag/service"
	handler3 "github.com/chenmingyong0423/fnote/server/internal/visit_log/handler"
	repository5 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository"
	dao5 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository/dao"
	service7 "github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp() (*gin.Engine, error) {
	database := ioc.NewMongoDB()
	fileDao := dao.NewFileDao(database)
	fileRepository := repository.NewFileRepository(fileDao)
	fileService := service.NewFileService(fileRepository)
	fileHandler := handler.NewFileHandler(fileService)
	categoryDao := dao2.NewCategoryDao(database)
	categoryRepository := repository2.NewCategoryRepository(categoryDao)
	countStatsDao := dao3.NewCountStatsDao(database)
	countStatsRepository := repository3.NewCountStatsRepository(countStatsDao)
	countStatsService := service2.NewCountStatsService(countStatsRepository)
	module := website_config.InitWebsiteConfigModule(database)
	iWebsiteConfigService := module.Svc
	categoryService := service3.NewCategoryService(categoryRepository, countStatsService, iWebsiteConfigService)
	categoryHandler := handler2.NewCategoryHandler(categoryService)
	emailService := service4.NewEmailService()
	msgTplDao := dao4.NewMsgTplDao(database)
	msgTplRepository := repository4.NewMsgTplRepository(msgTplDao)
	msgTplService := service5.NewMsgTplService(msgTplRepository)
	messageService := service6.NewMessageService(iWebsiteConfigService, emailService, msgTplService)
	post_likeModule := post_like.InitPostLikeModule(database)
	postModule := post.InitPostModule(database, module, countStatsService, fileService, post_likeModule)
	commentModule := comment.InitCommentModule(database, messageService, countStatsService, module, postModule)
	commentHandler := commentModule.Hdl
	websiteConfigHandler := module.Hdl
	friendModule := friend.InitFriendModule(database, messageService, module)
	friendHandler := friendModule.Hdl
	postHandler := postModule.Hdl
	visitLogDao := dao5.NewVisitLogDao(database)
	visitLogRepository := repository5.NewVisitLogRepository(visitLogDao)
	visitLogService := service7.NewVisitLogService(visitLogRepository)
	visitLogHandler := handler3.NewVisitLogHandler(visitLogService, countStatsService)
	msgTplHandler := handler4.NewMsgTplHandler(msgTplService)
	tagDao := dao6.NewTagDao(database)
	tagRepository := repository6.NewTagRepository(tagDao)
	tagService := service8.NewTagService(tagRepository, countStatsService)
	tagHandler := handler5.NewTagHandler(tagService)
	data_analysisModule := data_analysis.InitDataAnalysisModule(database, visitLogService, countStatsService, post_likeModule, commentModule)
	dataAnalysisHandler := data_analysisModule.Hdl
	countStatsHandler := handler6.NewCountStatsHandler(countStatsService)
	backupService := service9.NewBackupService(database)
	backupHandler := handler7.NewBackupHandler(backupService)
	writer := ioc.InitLogger()
	v, err := global.IsWebsiteInitializedFn(database)
	if err != nil {
		return nil, err
	}
	v2 := ioc.InitMiddlewares(writer, v)
	validators := ioc.InitGinValidators()
	post_indexModule := post_index.InitPostIndexModule(module)
	postIndexHandler := post_indexModule.Hdl
	post_draftModule := post_draft.InitPostDraftModule(database)
	postDraftHandler := post_draftModule.Hdl
	aggregate_postModule := aggregate_post.InitAggregatePostModule(postModule, post_draftModule)
	aggregatePostHandler := aggregate_postModule.Hdl
	postLikeHandler := post_likeModule.Hdl
	post_visitModule := post_visit.InitPostVisitModule(database)
	postVisitHandler := post_visitModule.Hdl
	engine, err := ioc.NewGinEngine(fileHandler, categoryHandler, commentHandler, websiteConfigHandler, friendHandler, postHandler, visitLogHandler, msgTplHandler, tagHandler, dataAnalysisHandler, countStatsHandler, backupHandler, v2, validators, postIndexHandler, postDraftHandler, aggregatePostHandler, postLikeHandler, postVisitHandler)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
