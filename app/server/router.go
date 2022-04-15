package server

import (
	"log"
	"net/http"
	"randi_firmansyah/app/handler/loginHandler"
	"randi_firmansyah/app/handler/productHandler"
	"randi_firmansyah/app/handler/tokenHandler"
	"randi_firmansyah/app/handler/userHandler"
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/repository"
	"randi_firmansyah/app/repository/productRepository"
	"randi_firmansyah/app/repository/userRepository"
	"randi_firmansyah/app/service"
	"randi_firmansyah/app/service/productService"
	"randi_firmansyah/app/service/userService"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {
	// try connect to database
	log.Println("Connecting to Database...")
	db, err := gorm.Open(mysql.Open(GetConnectionString()), &gorm.Config{})
	helper.CheckFatal(err)

	// migrate model to database
	db.AutoMigrate(&productModel.Product{}, &userModel.User{})
	log.Println("Database Connected")

	// generate repository
	allRepositories := repository.Repository{
		IProductRepository: productRepository.NewRepository(db),
		IUserRepository:    userRepository.NewRepository(db),
	}

	// try connect to redis
	log.Println("Connecting to Redis in Background...")
	redis := ConnectToRedis()

	// generate service
	allServices := service.Service{
		IProductService: productService.NewService(allRepositories),
		IUserService:    userService.NewService(allRepositories),
	}

	// generate handler
	product := productHandler.NewProductHandler(allServices, redis)
	user := userHandler.NewUserHandler(allServices, redis)

	// router
	r := chi.NewRouter()

	// global token
	r.Group(func(g chi.Router) {
		g.Get("/globaltoken", loginHandler.GenerateTokens)
	})

	// // login
	// r.Group(func(l chi.Router) {
	// 	l.Post("/login", loginHandler.Login)
	// })

	// product
	r.Group(func(p chi.Router) {
		p.Use(tokenHandler.GetToken) // pelindung token
		p.Get("/product", product.GetSemuaProduct)
		p.Get("/product/{id}", product.GetProductByID)
		p.Post("/product", product.PostProduct)
		p.Put("/product/{id}", product.UpdateProduct)
		p.Delete("/product/{id}", product.DeleteProduct)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		u.Get("/user", user.GetSemuaUser)
		// u.Get("/user/{id}", userHandler.GetUserById)
		// u.Post("/user", userHandler.PostUser)
		// u.Put("/user/{id}", userHandler.UpdateUser)
		// u.Delete("/user/{id}", userHandler.DeleteUser)
	})

	log.Println("Running Service")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Println("Error Starting Service")
	}
}
