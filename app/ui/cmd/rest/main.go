package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	ApplicationServicePassword "github.com/sergeygardner/meal-planner-api/application/service/password"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	ApplicationMiddleware "github.com/sergeygardner/meal-planner-api/infrastructure/service/jwt"
	RestHandler "github.com/sergeygardner/meal-planner-api/ui/rest/handler"
	RestAdminHandler "github.com/sergeygardner/meal-planner-api/ui/rest/handler/admin"
	"github.com/sergeygardner/meal-planner-api/ui/rest/service"
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
	flagRoutes            bool
)

func init() {
	jwtAuth = jwtauth.New("HS256", jwtKey, nil)
	flag.BoolVar(&flagJwtAuthentication, "jwtAuthentication", false, "To use JwtAuthentication")
	flag.BoolVar(&flagCORS, "cors", false, "To use CORS")
	flag.BoolVar(&flagContentTypeJSON, "contentTypeJSON", false, "To use JSON")
	flag.BoolVar(&flagDev, "dev", false, "To use Dev")
	flag.BoolVar(&flagRoutes, "flagRoutes", false, "Generate router documentation")
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

	prepareHTTPServer()
	startHTTPServer()
}

func middleWareJWT(router chi.Router) {
	if flagJwtAuthentication {
		router.Use(jwtauth.Verifier(jwtAuth))
		router.Use(jwtauth.Authenticator)
	}
	if flagContentTypeJSON {
		router.Use(render.SetContentType(render.ContentTypeJSON))
	}
}

func prepareHTTPServer() {
	router = chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)

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
}

func startHTTPServer() {
	projectPortInternal, status := os.LookupEnv("PROJECT_PORT_INTERNAL")

	if !status {
		panic("Can not find the internal port (PROJECT_PORT_INTERNAL)")
	}

	router.Route("/api", func(router chi.Router) {
		router.Get("/", RestHandler.API)
		router.Route("/{version}", func(router chi.Router) {
			router.Get("/", RestHandler.Version)
			router.Route("/auth", func(router chi.Router) {
				router.Group(func(router chi.Router) {
					middleWareJWT(router)
					router.Get("/", RestHandler.AuthCheck)
					router.Get("/refresh", RestHandler.AuthRefresh)
					router.Options("/refresh", RestHandler.AuthRefresh)
				})
				router.Post("/credentials", RestHandler.AuthCredentials)
				router.Options("/credentials", RestHandler.AuthCredentials)
				router.Post("/confirmation", RestHandler.AuthConfirmation)
				router.Options("/confirmation", RestHandler.AuthConfirmation)
				router.Post("/register", RestHandler.AuthRegister)
				router.Options("/register", RestHandler.AuthRegister)
			})
			router.Group(func(router chi.Router) {
				middleWareJWT(router)
				router.Route("/user", func(router chi.Router) {
					router.Get("/", RestHandler.UserInfo)
					router.Patch("/", RestHandler.UserUpdate)
					router.Delete("/", RestHandler.UserDelete)
				})
				router.Route("/recipes", func(router chi.Router) {
					router.Get("/", RestHandler.RecipesInfo)
					router.Post("/", RestHandler.RecipeCreate)
					router.Route("/{recipe_id}", func(router chi.Router) {
						router.Get("/", RestHandler.RecipeInfo)
						router.Patch("/", RestHandler.RecipeUpdate)
						router.Delete("/", RestHandler.RecipeDelete)
						router.Route("/categories", func(router chi.Router) {
							router.Get("/", RestHandler.RecipeCategoriesInfo)
							router.Post("/", RestHandler.RecipeCategoryCreate)
							router.Route("/{category_id}", func(router chi.Router) {
								router.Get("/", RestHandler.RecipeCategoryInfo)
								router.Patch("/", RestHandler.RecipeCategoryUpdate)
								router.Delete("/", RestHandler.RecipeCategoryDelete)
								setAltNameRouting(router)
							})
						})
						router.Route("/ingredients", func(router chi.Router) {
							router.Get("/", RestHandler.RecipeIngredientsInfo)
							router.Post("/", RestHandler.RecipeIngredientCreate)
							router.Route("/{ingredient_id}", func(router chi.Router) {
								router.Get("/", RestHandler.RecipeIngredientInfo)
								router.Patch("/", RestHandler.RecipeIngredientUpdate)
								router.Delete("/", RestHandler.RecipeIngredientDelete)
								router.Route("/measures", func(router chi.Router) {
									router.Get("/", RestHandler.RecipeMeasuresInfo)
									router.Post("/", RestHandler.RecipeMeasureCreate)
									router.Get("/{measure_id}", RestHandler.RecipeMeasureInfo)
									router.Patch("/{measure_id}", RestHandler.RecipeMeasureUpdate)
									router.Delete("/{measure_id}", RestHandler.RecipeMeasureDelete)
								})
								setPictureRouting(router)
								setAltNameRouting(router)
							})
						})
						router.Route("/processes", func(router chi.Router) {
							router.Get("/", RestHandler.RecipeProcessesInfo)
							router.Post("/", RestHandler.RecipeProcessCreate)
							router.Route("/{process_id}", func(router chi.Router) {
								router.Get("/", RestHandler.RecipeProcessInfo)
								router.Patch("/", RestHandler.RecipeProcessUpdate)
								router.Delete("/", RestHandler.RecipeProcessDelete)
								setPictureRouting(router)
								setAltNameRouting(router)
							})
						})
						setPictureRouting(router)
						setAltNameRouting(router)
					})
				})
				router.Route("/units", func(router chi.Router) {
					router.Get("/", RestHandler.UnitsInfo)
					router.Post("/", RestHandler.UnitCreate)
					router.Get("/{unit_id}", RestHandler.UnitInfo)
					router.Patch("/{unit_id}", RestHandler.UnitUpdate)
					router.Delete("/{unit_id}", RestHandler.UnitDelete)
					setAltNameRouting(router)
				})
				router.Route("/categories", func(router chi.Router) {
					router.Get("/", RestHandler.CategoriesInfo)
					router.Post("/", RestHandler.CategoryCreate)
					router.Route("/{category_id}", func(router chi.Router) {
						router.Get("/", RestHandler.CategoryInfo)
						router.Patch("/", RestHandler.CategoryUpdate)
						router.Delete("/", RestHandler.CategoryDelete)
						setPictureRouting(router)
						setAltNameRouting(router)
					})
				})
				router.Route("/ingredients", func(router chi.Router) {
					router.Get("/", RestHandler.IngredientsInfo)
					router.Post("/", RestHandler.IngredientCreate)
					router.Route("/{ingredient_id}", func(router chi.Router) {
						router.Get("/", RestHandler.IngredientInfo)
						router.Patch("/", RestHandler.IngredientUpdate)
						router.Delete("/", RestHandler.IngredientDelete)
						setPictureRouting(router)
						setAltNameRouting(router)
					})
				})
				router.Route("/planners", func(router chi.Router) {
					router.Get("/", RestHandler.PlannersInfo)
					router.Post("/", RestHandler.PlannerCreate)
					router.Route("/{planner_id}", func(router chi.Router) {
						router.Get("/", RestHandler.PlannerInfo)
						router.Patch("/", RestHandler.PlannerUpdate)
						router.Delete("/", RestHandler.PlannerDelete)
						router.Get("/calculate", RestHandler.PlannerCalculateInfo)
						router.Route("/intervals", func(router chi.Router) {
							router.Get("/", RestHandler.PlannerIntervalsInfo)
							router.Post("/", RestHandler.PlannerIntervalCreate)
							router.Route("/{interval_id}", func(router chi.Router) {
								router.Get("/", RestHandler.PlannerIntervalInfo)
								router.Patch("/", RestHandler.PlannerIntervalUpdate)
								router.Delete("/", RestHandler.PlannerIntervalDelete)
								router.Route("/recipes", func(router chi.Router) {
									router.Get("/", RestHandler.PlannerRecipesInfo)
									router.Post("/", RestHandler.PlannerRecipeCreate)
									router.Route("/{recipe_id}", func(router chi.Router) {
										router.Get("/", RestHandler.PlannerRecipeInfo)
										router.Patch("/", RestHandler.PlannerRecipeUpdate)
										router.Delete("/", RestHandler.PlannerRecipeDelete)
									})
								})
							})
						})
					})
				})
			})
			router.Group(func(router chi.Router) {
				router.Route("/admin", func(router chi.Router) {
					router.Route("/users", func(router chi.Router) {
						middleWareJWT(router)
						router.Use(service.EnsureAdmin)
						router.Get("/", RestAdminHandler.UserInfo)
						router.Patch("/", RestAdminHandler.UserUpdate)
						router.Delete("/", RestAdminHandler.UserDelete)
					})
				})
			})
		})
	})
	router.HandleFunc("/", RestHandler.Index)

	if flagRoutes {
		err := os.WriteFile("rest/routers.json", []byte(docgen.JSONRoutesDoc(router)), 0644)
		if err != nil {
			return
		}
		return
	}

	log.Info("The server is running")
	log.Fatal(http.ListenAndServe(projectPortInternal, router))
}

func setAltNameRouting(router chi.Router) {
	router.Route("/alt-names", func(router chi.Router) {
		router.Get("/", RestHandler.AltNamesInfo)
		router.Post("/", RestHandler.AltNameCreate)
		router.Get("/{alt_name_id}", RestHandler.AltNameInfo)
		router.Patch("/{alt_name_id}", RestHandler.AltNameUpdate)
		router.Delete("/{alt_name_id}", RestHandler.AltNameDelete)
	})
}

func setPictureRouting(router chi.Router) {
	router.Route("/pictures", func(router chi.Router) {
		router.Get("/", RestHandler.PicturesInfo)
		router.Post("/", RestHandler.PictureCreate)
		router.Route("/{picture_id}", func(router chi.Router) {
			router.Get("/", RestHandler.PictureInfo)
			router.Patch("/", RestHandler.PictureUpdate)
			router.Delete("/", RestHandler.PictureDelete)
			setAltNameRouting(router)
		})
	})
}
