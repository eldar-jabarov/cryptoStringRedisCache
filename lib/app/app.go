package app

import (
	"cryptoStringRedisCache/internal/infrastructure/controller"
	"cryptoStringRedisCache/internal/infrastructure/service"
	"cryptoStringRedisCache/lib/cache"
	"cryptoStringRedisCache/lib/responder"
	"cryptoStringRedisCache/lib/swagger"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"sync"
)

type App struct {
	wait          sync.WaitGroup
	cryptoService service.Crypter
}

func New() *App {
	return &App{}
}

func (a *App) Run() *App {
	serviceCrypter := generateService()

	controllerCrypter := generateController(serviceCrypter)

	route := generateRoute(controllerCrypter)

	go a.listenAndServe(route)
	a.wait.Add(1)
	return a
}

func (a *App) Wait() {
	a.wait.Wait()
}

func generateService() service.Crypter {
	cacher := cache.NewCache(fmt.Sprintf("redis:%s", getEnv("REDIS_PORT")))
	crypter := cache.NewServiceCrypto(service.NewCrypter(), cacher)

	proxy := cache.NewServiceCrypto(crypter, cacher)
	return proxy
}

func generateController(service service.Crypter) controller.Crypter {
	return controller.NewCrypto(responder.New(), service)
}

func generateRoute(controller controller.Crypter) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.DefaultLogger)
	generateRouteSwagger(router)

	router.Post("/api/crypto", controller.Crypto)
	return router
}

func generateRouteSwagger(router chi.Router) {
	router.Get("/swagger/", swagger.SwaggerUI)
	router.Get("/static/swagger.json", swagger.SwaggerJson)
}

func getEnv(name string) string {
	res := os.Getenv(name)
	if res == "" {
		log.Printf("env not found name: %s", name)
	}
	return res
}

func (a *App) listenAndServe(router chi.Router) {
	log.Println(http.ListenAndServe(":8080", router))
	a.wait.Done()
}
