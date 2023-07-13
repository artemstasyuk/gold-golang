package rpc

import (
	"net/http"

	"apisrv/pkg/vt"

	"apisrv/pkg/db"
	"apisrv/pkg/embedlog"

	zm "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
)

var (
	ErrNotImplemented = zenrpc.NewStringError(http.StatusInternalServerError, "Not implemented")
	ErrInternal       = zenrpc.NewStringError(http.StatusInternalServerError, "Internal error")
)

var allowDebugFn = func() zm.AllowDebugFunc {
	return func(req *http.Request) bool {
		return req != nil && req.FormValue("__level") == "5"
	}
}

//go:generate zenrpc

// New returns new zenrpc Server.
func New(dbo db.DB, logger embedlog.Logger, isDevel bool) zenrpc.Server {
	rpc := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD: true,
		AllowCORS: true,
	})

	rpc.Use(
		zm.WithDevel(isDevel),
		zm.WithHeaders(),
		zm.WithSentry(zm.DefaultServerName),
		zm.WithNoCancelContext(),
		zm.WithMetrics(zm.DefaultServerName),
		zm.WithTiming(isDevel, allowDebugFn()),
		zm.WithSQLLogger(dbo.DB, isDevel, allowDebugFn(), allowDebugFn()),
	)

	if errlog, stdlog := logger.Loggers(); errlog != nil && stdlog != nil {
		rpc.Use(
			zm.WithAPILogger(stdlog.Printf, zm.DefaultServerName),
			zm.WithErrorLogger(errlog.Printf, zm.DefaultServerName),
		)
	}

	// services
	rpc.RegisterAll(map[string]zenrpc.Invoker{
		"auth":     vt.NewAuthService(dbo, logger),
		"users":    vt.NewUserService(dbo, logger),
		"news":     vt.NewNewsService(dbo, logger),
		"category": vt.NewCategoryService(dbo, logger),
		"tags":     vt.NewTagService(dbo, logger),
	})

	return rpc
}

//nolint:unused
func internalError(err error) *zenrpc.Error {
	return zenrpc.NewError(http.StatusInternalServerError, err)
}
