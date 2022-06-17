package server

import (
	"log"
	"net/http"
	"os"
	"randi_firmansyah/app/handler/cartHandler"
	"randi_firmansyah/app/handler/categoryHandler"
	"randi_firmansyah/app/handler/orderHandler"
	"randi_firmansyah/app/handler/productHandler"
	"randi_firmansyah/app/handler/tokenHandler"
	"randi_firmansyah/app/handler/userHandler"
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/repository"
	"randi_firmansyah/app/repository/cartRepository"
	"randi_firmansyah/app/repository/categoryRepository"
	"randi_firmansyah/app/repository/orderRepository"
	"randi_firmansyah/app/repository/productRepository"
	"randi_firmansyah/app/repository/userRepository"
	"randi_firmansyah/app/service"
	"randi_firmansyah/app/service/cartService"
	"randi_firmansyah/app/service/categoryService"
	"randi_firmansyah/app/service/orderService"
	"randi_firmansyah/app/service/productService"
	"randi_firmansyah/app/service/userService"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {
	// try connect to database
	log.Println("Connecting to Database...")
	db, err := gorm.Open(mysql.Open(getConnectionString()), &gorm.Config{})
	helper.CheckFatal(err)

	// migrate model to database
	// db.AutoMigrate(&productModel.Product{}, &userModel.User{}, &cartModel.Cart{}, &categoryModel.Category{}, &orderModel.Order{})
	log.Println("Database Connected")

	// generate repository
	allRepositories := repository.Repository{
		IProductRepository:  productRepository.NewRepository(db),
		IUserRepository:     userRepository.NewRepository(db),
		ICartRepository:     cartRepository.NewRepository(db),
		ICategoryRepository: categoryRepository.NewRepository(db),
		IOrderRepository:    orderRepository.NewRepository(db),
	}

	// try connect to redis
	log.Println("Connecting to Redis in Background...")
	redis := connectToRedis()

	// generate service
	allServices := service.Service{
		IProductService:  productService.NewService(allRepositories),
		IUserService:     userService.NewService(allRepositories),
		ICartService:     cartService.NewService(allRepositories),
		ICategoryService: categoryService.NewService(allRepositories),
		IOrderService:    orderService.NewService(allRepositories),
	}

	// generate handler
	product := productHandler.NewProductHandler(allServices, redis)
	order := orderHandler.NewOrderHandler(allServices, redis)
	user := userHandler.NewUserHandler(allServices, redis)
	// login := loginHandler.NewLoginHandler(allServices)
	cart := cartHandler.NewCartHandler(allServices, redis)
	category := categoryHandler.NewCategoryHandler(allServices, redis)

	// router
	r := chi.NewRouter()

	// check service
	r.Group(func(g chi.Router) {
		g.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response.ResponseRunningService(w)
		})
	})

	// // global token
	// r.Group(func(g chi.Router) {
	// 	g.Get("/globaltoken", login.GenerateToken)
	// })

	// // login
	// r.Group(func(l chi.Router) {
	// 	l.Post("/login", login.Login)
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

	// order
	r.Group(func(p chi.Router) {
		p.Use(tokenHandler.GetToken) // pelindung token
		p.Get("/order", order.GetSemuaOrder)
		p.Get("/order/id/{id}", order.GetOrderByID)
		p.Put("/order/{saldo}/{id}", order.PayOrder)
		p.Get("/order/{username}", order.GetOrderByUsername)
		p.Post("/order", order.PostOrder)
		// p.Put("/order/{id}", order.UpdateProduct)
		// p.Delete("/order/{id}", order.DeleteProduct)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		// u.Get("/user", user.GetSemuaUser)
		// u.Get("/user/{id}", user.GetUserByID)
		// u.Get("/user/{username}", user.GetUserByUsername)
		u.Post("/user", user.PostUser)
		u.Put("/user/{username}", user.UpdateUser)
		u.Delete("/user/{id}", user.DeleteUser)
	})

	// cart
	r.Group(func(c chi.Router) {
		c.Use(tokenHandler.GetToken) // pelindung token
		c.Get("/cart", cart.GetSemuaCart)
		c.Get("/cart/{username}", cart.GetCartByUsername)
		c.Post("/cart", cart.PostCart)
		c.Put("/cart/{id}/{qty}", cart.UpdateCart)
		c.Delete("/cart/{id}", cart.DeleteCart)
	})

	// category
	r.Group(func(c chi.Router) {
		c.Use(tokenHandler.GetToken) // pelindung token
		c.Get("/category", category.GetSemuaCategory)
		c.Get("/category/{id}", category.GetCategoryByID)
		c.Post("/category", category.PostCategory)
		c.Put("/category/{id}", category.UpdateCategory)
		c.Delete("/category/{id}", category.DeleteCategory)
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Println("Service running on " + host + ":" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Println("Error Starting Service")
	}
}
