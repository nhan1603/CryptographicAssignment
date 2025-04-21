package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/iam"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/auth"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/menus"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/orders"
	"github.com/nhan1603/CryptographicAssignment/api/internal/handler/rest/authenticated/v1/operation"
	authHandler "github.com/nhan1603/CryptographicAssignment/api/internal/handler/rest/public/v1/auth"
)

type router struct {
	ctx       context.Context
	authCtrl  auth.Controller
	menuCtrl  menus.Controller
	orderCtrl orders.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.authenticated)
	r.Group(rtr.public)
}

func (rtr router) authenticated(r chi.Router) {
	prefix := "/api/authenticated"

	r.Group(func(r chi.Router) {
		r.Use(iam.AuthenticateUserMiddleware(rtr.ctx))
		prefix = prefix + "/v1"

		operationH := operation.New(rtr.menuCtrl, rtr.orderCtrl)
		r.Get(prefix+"/menu", operationH.GetMenuItems())
		r.Post(prefix+"/order", operationH.CreateOrder())
		r.Post(prefix+"/order/update_status", operationH.UpdateOrderStatus())
	})
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		r.Group(func(r chi.Router) {
			authH := authHandler.New(rtr.authCtrl)
			r.Post(prefix+"/login", authH.AuthenticateOperationUser())
			r.Post(prefix+"/user", authH.CreateUser())

			operationH := operation.New(rtr.menuCtrl, rtr.orderCtrl)
			r.Get(prefix+"/menu", operationH.GetMenuItems())
		})
	})
}
