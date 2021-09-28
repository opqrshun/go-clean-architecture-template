/*
 * Swagger MapSNS
 *
 * gobackend api
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://http-generator.tech)
 */

package http

import (
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gobackend/internal/config"
	"gobackend/pkg/conv"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// const UseAuthURLPattern = "^/v1/user$|^/v1/user/.*|.*/favorite$|.*/like$"
var UseAuthURLList = []string{
	"/v1/entities/:entity-id/favorite",
	"/v1/entities/:entity-id/like",
}

////UseSentryLocalHub
//func UseSentryLocalHub (ctx *gin.Context, tagKey string,tagValue string) {
//  if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
//		hub.WithScope(func(scope *sentry.Scope) {
//			scope.SetTag(tagKey, tagValue)
//		})
//	}
//}
// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	_config := config.GetAppConfig()
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: _config.SentryDSN,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// Set Sentry
	// tag main router
	router.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("app", "mainRouter")
		}
		ctx.Next()
	})

	//Add CORS midleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:6006", "*"}
	config.AllowHeaders = []string{"Origin", "content-type", "authorization"}
	config.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE"}
	router.Use(cors.New(config))

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			if ok := conv.Contains(UseAuthURLList, route.Pattern); ok {
				router.GET(route.Pattern, Auth(), route.HandlerFunc)
			} else {
				router.GET(route.Pattern, route.HandlerFunc)
			}

		case http.MethodPost:
			router.POST(route.Pattern, Auth(), route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, Auth(), route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, Auth(), route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, Auth(), route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/v1/",
		Index,
	},

	{
		"CreateAttribute",
		http.MethodPost,
		"/v1/entities/:entity-id/attributes",
		CreateAttribute,
	},

	{
		"DeleteAttribute",
		http.MethodDelete,
		"/v1/attributes/:attribute-id",
		DeleteAttribute,
	},

	{
		"GetAttribute",
		http.MethodGet,
		"/v1/attributes/:attribute-id",
		GetAttribute,
	},

	{
		"GetAttributesByEntity",
		http.MethodGet,
		"/v1/entities/:entity-id/attributes",
		GetAttributesByEntity,
	},

	{
		"UpdateAttribute",
		http.MethodPatch,
		"/v1/attributes/:attribute-id",
		UpdateAttribute,
	},

	{
		"CreateEntity",
		http.MethodPost,
		"/v1/entities",
		CreateEntity,
	},

	{
		"GetEntity",
		http.MethodGet,
		"/v1/entities/:entity-id",
		GetEntity,
	},

	{
		"GetEntities",
		http.MethodGet,
		"/v1/entities",
		GetEntities,
	},

	{
		"UpdateEntity",
		http.MethodPatch,
		"/v1/entities/:entity-id",
		UpdateEntity,
	},
}
