package controller

import (
	"fmt"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/middleware"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"

	"github.com/acm-uiuc/core/controller/auth"
	"github.com/acm-uiuc/core/controller/docs"
	"github.com/acm-uiuc/core/controller/group"
	"github.com/acm-uiuc/core/controller/user"
)

type Controller struct {
	*echo.Echo
	svc *service.Service
}

func New(svc *service.Service) (*Controller, error) {
	controller := &Controller{
		Echo: echo.New(),
		svc:  svc,
	}

	docsController := docs.New(controller.svc)
	authController := auth.New(controller.svc)
	userController := user.New(controller.svc)
	groupController := group.New(controller.svc)

	controller.Use(echoMiddleware.Logger())
	controller.Use(echoMiddleware.Recover())
	controller.Use(middleware.Context(controller.svc))

	controller.GET(
		"/api",
		Chain(docsController.Documentation),
	)

	controller.GET(
		"/api/auth/:provider",
		Chain(authController.GetOAuthRedirect),
	)
	controller.GET(
		"/api/auth/:provider/redirect",
		Chain(authController.GetOAuthRedirectLanding),
	)
	controller.POST(
		"/api/auth/:provider",
		Chain(authController.GetToken),
	)

	controller.GET(
		"/api/user",
		Chain(userController.GetUser, middleware.AuthorizeMatchAny(
			controller.svc, middleware.AuthorizeMatchParameters{
				Marks: model.UserValidMarks,
			},
		)),
	)
	controller.POST(
		"/api/user",
		Chain(userController.CreateUser, middleware.AuthorizeMatchAny(
			controller.svc, middleware.AuthorizeMatchParameters{
				Marks: model.UserValidMarks,
			},
		)),
	)
	controller.GET(
		"/api/user/filter",
		Chain(userController.GetUsers, middleware.AuthorizeMatchAny(
			controller.svc, middleware.AuthorizeMatchParameters{
				Marks:      []string{model.UserMarkRecruiter},
				Committees: []string{model.GroupTop4},
			},
		)),
	)
	controller.POST(
		"/api/user/mark",
		Chain(userController.MarkUser, middleware.AuthorizeMatchAny(
			controller.svc, middleware.AuthorizeMatchParameters{
				Committees: []string{model.GroupTop4},
			},
		)),
	)

	controller.GET(
		"/api/group",
		Chain(groupController.GetGroups),
	)
	controller.POST(
		"/api/group/verify",
		Chain(groupController.VerifyMembership),
	)

	return controller, nil
}

func ContextConverter(handler func(*context.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		ctx, ok := c.(*context.Context)
		if !ok {
			return fmt.Errorf("failed to convert context")
		}
		return handler(ctx)
	}
}

func Chain(handler func(*context.Context) error, middlewares ...(func(next echo.HandlerFunc) echo.HandlerFunc)) func(echo.Context) error {
	aggregated := ContextConverter(handler)
	for i := range middlewares {
		aggregated = middlewares[len(middlewares)-i-1](aggregated)
	}

	return aggregated
}