package main

import (
	"flag"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	ApplicationServicePassword "github.com/sergeygardner/meal-planner-api/application/service/password"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	ApplicationMiddleware "github.com/sergeygardner/meal-planner-api/infrastructure/service/jwt"
	"github.com/sergeygardner/meal-planner-api/ui/graphql"
	"github.com/sergeygardner/meal-planner-api/ui/graphql/directive"
	RestHandler "github.com/sergeygardner/meal-planner-api/ui/rest/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	router                *chi.Mux
	jwtKey                = ApplicationMiddleware.GetJwtKey()
	jwtAuth               *jwtauth.JWTAuth
	flagJwtAuthentication bool
	flagCORS              bool
	flagContentTypeJSON   bool
	flagDev               bool
)

func init() {
	jwtAuth = jwtauth.New("HS256", jwtKey, nil)
	flag.BoolVar(&flagJwtAuthentication, "jwtAuthentication", false, "To use JwtAuthentication")
	flag.BoolVar(&flagCORS, "cors", false, "To use CORS")
	flag.BoolVar(&flagContentTypeJSON, "contentTypeJSON", false, "To use JSON")
	flag.BoolVar(&flagDev, "dev", false, "To use Dev")
}

func main() {
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	if flagDev {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()

	prepareGraphQLServer()
	prepareGraphQLServer()
	startGraphQLServer()
}

func prepareGraphQLServer() {
	router = chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	if flagCORS {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Methods", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
	}

	salt, errorPasswordSalt := os.LookupEnv("PASSWORD_SALT")

	if !errorPasswordSalt {
		panic("Can not find the password salt (PASSWORD_SALT)")
	}

	ApplicationServicePassword.SetPasswordSalt(salt)

	if flagContentTypeJSON {
		router.Use(render.SetContentType(render.ContentTypeJSON))
	}
	config := graphql.Config{Resolvers: &graphql.Resolver{}}

	if flagJwtAuthentication {
		config.Directives.Auth = directive.Auth
	}

	srv := handler.New(graphql.NewExecutableSchema(config))

	srv.AddTransport(transport.POST{})

	if flagDev {
		srv.Use(extension.Introspection{})
		router.Handle("/graphql/playground", playground.Handler("graphql Playground", "/graphql"))
	}

	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(jwtAuth))
		router.Handle("/graphql", srv)
	})
}

func startGraphQLServer() {
	projectPortInternal, status := os.LookupEnv("PROJECT_PORT_INTERNAL")

	if !status {
		panic("Can not find the internal port (PROJECT_PORT_INTERNAL)")
	}

	router.HandleFunc("/", RestHandler.Index)

	log.Info("The server is running")
	log.Fatal(http.ListenAndServe(projectPortInternal, router))
}
