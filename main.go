package main

import (
	"fmt"
	"net/http"
	"os"

	driver "./driver"

	ph "./handler/http"
	jwtService "./service/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "admin", dbPass, dbName)
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	jwtServiceObj := jwtService.Init(tokenAuth)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	organizerHandler := ph.InitOrganizerHandler(connection)
	userHandler := ph.InitUserHandler(connection)
	accountHandler := ph.InitAccountHandler(connection)
	authHandler := ph.InitAuthHandler(connection, jwtServiceObj)
	rundownHandler := ph.InitRundownHandler(connection)
	rundownItemHandler := ph.InitRundownItemHandler(connection)
	concregationHandler := ph.InitConcregationHandler(connection)
	deviceInventoryHandler := ph.InitDeviceInventoryHandler(connection)
	serviceScheduleHandler := ph.InitServiceScheduleHandler(connection)
	sectorCoordinator := ph.InitSectorCoordinatorHandler(connection)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Route("/public", func(rt chi.Router) {
			rt.Route("/organizer", func(route chi.Router) {
				route.Get("/{name}", organizerHandler.GetByName)
				route.Get("/city/{name}", organizerHandler.GetByCity)
				route.Get("/province/{name}", organizerHandler.GetByProvince)
				route.Get("/province/name/{province}/{name}", organizerHandler.GetByProvinceAndName)
			})

			rt.Route("/rundown", func(route chi.Router) {
				route.Get("/{organizerId:[0-9]+}", rundownHandler.GetByOrganizerId)
				route.Get("/{organizerId:[0-9]+}/{startDate:[0-100]+}/{endDate:[0-100]+}", rundownHandler.GetByOrganizerIdAndDate)
			})

			rt.Route("/rundown_item", func(route chi.Router) {
				route.Get("/{rundownId:[0-9]+}", rundownItemHandler.GetByRundownId)
			})

			rt.Route("/auth", func(route chi.Router) {
				route.Post("/login", authHandler.Login)
				route.Post("/register", authHandler.Register)
			})
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtServiceObj.Verifier())
		r.Use(jwtServiceObj.Authenticator())

		r.Route("/admin", func(rt chi.Router) {
			rt.Route("/auth", func(route chi.Router) {
				route.Post("/update", authHandler.Update)
			})

			rt.Route("/organizer", func(route chi.Router) {
				route.Get("/getById/{id:[0-9]+}", organizerHandler.GetByID)
			})

			rt.Route("/user", func(route chi.Router) {
				route.Get("/{id:[0-9]+}", userHandler.GetByID)
			})

			rt.Route("/account", func(route chi.Router) {
				route.Get("/{id:[0-9]+}", accountHandler.GetByID)
			})

			rt.Route("/rundown", func(route chi.Router) {
				route.Get("/{organizerId:[0-9]+}", rundownHandler.GetByOrganizerId)
				route.Post("/", rundownHandler.Create)
				route.Put("/", rundownHandler.Update)
				route.Delete("/{id:[0-9]+}", rundownHandler.Delete)
			})

			rt.Route("/rundown_item", func(route chi.Router) {
				route.Get("/{rundownId:[0-9]+}", rundownItemHandler.GetByRundownId)
				route.Post("/", rundownItemHandler.Create)
				route.Put("/", rundownItemHandler.Update)
				route.Delete("/{id:[0-9]+}", rundownItemHandler.Delete)
			})

			rt.Route("/concregation", func(route chi.Router) {
				route.Post("/", concregationHandler.Create)
				route.Put("/", concregationHandler.Update)
				route.Delete("/{id:[0-9]+}", concregationHandler.Delete)
				route.Get("/{organizerId:[0-9]+}", concregationHandler.GetByOrganizerId)
			})

			rt.Route("/device_inventory", func(route chi.Router) {
				route.Post("/", deviceInventoryHandler.Create)
				route.Put("/", deviceInventoryHandler.Update)
				route.Delete("/{id:[0-9]+}", deviceInventoryHandler.Delete)
				route.Get("/{organizerId:[0-9]+}", deviceInventoryHandler.GetByOrganizerId)
			})

			rt.Route("/service_schedule", func(route chi.Router) {
				route.Post("/", serviceScheduleHandler.Create)
				route.Put("/", serviceScheduleHandler.Update)
				route.Delete("/{id:[0-9]+}", serviceScheduleHandler.Delete)
				route.Get("/{organizerId:[0-9]+}", serviceScheduleHandler.GetByOrganizerId)
			})

			rt.Route("/sector_coordinator", func(route chi.Router) {
				route.Post("/", sectorCoordinator.Create)
				route.Put("/", sectorCoordinator.Update)
				route.Delete("/{id:[0-9]+}", sectorCoordinator.Delete)
				route.Get("/{organizerId:[0-9]+}", sectorCoordinator.GetByOrganizerId)
			})
		})
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":3000", r)
}
