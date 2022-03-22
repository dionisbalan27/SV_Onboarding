package app

import (
	"backend-api/delivery/user_delivery"
	"backend-api/repository/user_repository"
	"backend-api/usecase/jwt_usecase"
	"backend-api/usecase/user_usecase"

	"backend-api/delivery/product_delivery"
	"backend-api/repository/product_repository"
	"backend-api/usecase/product_usecase"

	"backend-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(postgresConn *gorm.DB) *gin.Engine {

	userRepository := user_repository.GetUserRepository(postgresConn)
	jwtUsecase := jwt_usecase.GetJwtUsecase(userRepository)
	userUsecase := user_usecase.GetUserUsecase(jwtUsecase, userRepository)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)

	productRepository := product_repository.GetProductRepository(postgresConn)
	productUsecase := product_usecase.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)
	defaultCors := middleware.CORSMiddleware()

	router := gin.Default()
	router.Use(defaultCors)

	router.POST("/user", userDelivery.CreateNewUser)
	router.POST("/login", userDelivery.UserLogin)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.JWTAuth(jwtUsecase))
	{
		protectedRoutes.GET("/users", userDelivery.GetAllUsers)
		protectedRoutes.GET("/user/:id", userDelivery.GetUserById)
		protectedRoutes.GET("/products", productDelivery.GetAllProducts)
		protectedRoutes.GET("/products/:id", productDelivery.GetProductById)
		protectedRoutes.POST("/product", productDelivery.CreateNewProduct)
		protectedRoutes.PUT("/products/:id", productDelivery.UpdateProductData)
		protectedRoutes.DELETE("/product/:id", productDelivery.DeleteProductById)
	}

	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.JWTAuthAdmin(jwtUsecase, userRepository))
	{
		adminRoutes.PUT("/user/:id", userDelivery.UpdateUserData)
		adminRoutes.DELETE("/user/:id", userDelivery.DeleteUserById)
	}

	checkerRoutes := router.Group("/")
	checkerRoutes.Use(middleware.JWTAuthChecker(jwtUsecase, userRepository))
	{
		checkerRoutes.PUT("/products/:id/checked", productDelivery.UpdateCheckProduct)
	}

	signerRoutes := router.Group("/")
	signerRoutes.Use(middleware.JWTAuthSigner(jwtUsecase, userRepository))
	{
		signerRoutes.PUT("/products/:id/published", productDelivery.UpdatePublishProduct)
	}

	router.Run(":8001")

	return router
}
