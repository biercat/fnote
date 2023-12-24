// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/backend/internal/category/handler"
	"github.com/chenmingyong0423/fnote/backend/internal/category/repository"
	"github.com/chenmingyong0423/fnote/backend/internal/category/repository/dao"
	service2 "github.com/chenmingyong0423/fnote/backend/internal/category/service"
	"github.com/chenmingyong0423/fnote/backend/internal/comment/hanlder"
	repository3 "github.com/chenmingyong0423/fnote/backend/internal/comment/repository"
	dao3 "github.com/chenmingyong0423/fnote/backend/internal/comment/repository/dao"
	service3 "github.com/chenmingyong0423/fnote/backend/internal/comment/service"
	handler2 "github.com/chenmingyong0423/fnote/backend/internal/config/handler"
	repository4 "github.com/chenmingyong0423/fnote/backend/internal/config/repository"
	dao4 "github.com/chenmingyong0423/fnote/backend/internal/config/repository/dao"
	service4 "github.com/chenmingyong0423/fnote/backend/internal/config/service"
	repository2 "github.com/chenmingyong0423/fnote/backend/internal/count_stats/repository"
	dao2 "github.com/chenmingyong0423/fnote/backend/internal/count_stats/repository/dao"
	"github.com/chenmingyong0423/fnote/backend/internal/count_stats/service"
	service6 "github.com/chenmingyong0423/fnote/backend/internal/email/service"
	hanlder2 "github.com/chenmingyong0423/fnote/backend/internal/friend/hanlder"
	repository7 "github.com/chenmingyong0423/fnote/backend/internal/friend/repository"
	dao7 "github.com/chenmingyong0423/fnote/backend/internal/friend/repository/dao"
	service9 "github.com/chenmingyong0423/fnote/backend/internal/friend/service"
	"github.com/chenmingyong0423/fnote/backend/internal/ioc"
	service8 "github.com/chenmingyong0423/fnote/backend/internal/message/service"
	handler5 "github.com/chenmingyong0423/fnote/backend/internal/message_template/handler"
	repository6 "github.com/chenmingyong0423/fnote/backend/internal/message_template/repository"
	dao6 "github.com/chenmingyong0423/fnote/backend/internal/message_template/repository/dao"
	service7 "github.com/chenmingyong0423/fnote/backend/internal/message_template/service"
	handler3 "github.com/chenmingyong0423/fnote/backend/internal/post/handler"
	repository5 "github.com/chenmingyong0423/fnote/backend/internal/post/repository"
	dao5 "github.com/chenmingyong0423/fnote/backend/internal/post/repository/dao"
	service5 "github.com/chenmingyong0423/fnote/backend/internal/post/service"
	handler4 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/handler"
	repository8 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/repository"
	dao8 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/repository/dao"
	service10 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/service"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp(cfgPath string) (*gin.Engine, error) {
	config := ioc.InitConfig(cfgPath)
	database := ioc.NewMongoDB(config)
	categoryDao := dao.NewCategoryDao(database)
	categoryRepository := repository.NewCategoryRepository(categoryDao)
	countStatsDao := dao2.NewCountStatsDao(database)
	countStatsRepository := repository2.NewCountStatsRepository(countStatsDao)
	countStatsService := service.NewCountStatsService(countStatsRepository)
	categoryService := service2.NewCategoryService(categoryRepository, countStatsService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	commentDao := dao3.NewCommentDao(database)
	commentRepository := repository3.NewCommentRepository(commentDao)
	commentService := service3.NewCommentService(commentRepository)
	configDao := dao4.NewConfigDao(database)
	configRepository := repository4.NewConfigRepository(configDao)
	configService := service4.NewConfigService(configRepository)
	postDao := dao5.NewPostDao(database)
	postRepository := repository5.NewPostRepository(postDao)
	postService := service5.NewPostService(postRepository)
	emailService := service6.NewEmailService()
	msgTplDao := dao6.NewMsgTplDao(database)
	msgTplRepository := repository6.NewMsgTplRepository(msgTplDao)
	msgTplService := service7.NewMsgTplService(msgTplRepository)
	messageService := service8.NewMessageService(configService, emailService, msgTplService)
	commentHandler := hanlder.NewCommentHandler(commentService, configService, postService, messageService)
	configHandler := handler2.NewConfigHandler(configService)
	friendDao := dao7.NewFriendDao(database)
	friendRepository := repository7.NewFriendRepository(friendDao)
	friendService := service9.NewFriendService(friendRepository)
	friendHandler := hanlder2.NewFriendHandler(friendService, messageService, configService)
	postHandler := handler3.NewPostHandler(postService)
	visitLogDao := dao8.NewVisitLogDao(database)
	visitLogRepository := repository8.NewVisitLogRepository(visitLogDao)
	visitLogService := service10.NewVisitLogService(visitLogRepository)
	visitLogHandler := handler4.NewVisitLogHandler(visitLogService, configService)
	msgTplHandler := handler5.NewMsgTplHandler(msgTplService)
	writer := ioc.InitLogger(config)
	v := ioc.InitMiddlewares(config, writer)
	validators := ioc.InitGinValidators()
	engine, err := ioc.NewGinEngine(categoryHandler, commentHandler, configHandler, friendHandler, postHandler, visitLogHandler, msgTplHandler, v, validators)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
