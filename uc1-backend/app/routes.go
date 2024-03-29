package app

import (
	"backend-api/delivery/role_delivery"
	"backend-api/delivery/user_delivery"
	"backend-api/repository/role_repository"
	"backend-api/repository/user_repository"
	"backend-api/usecase/jwt_usecase"
	"backend-api/usecase/role_usecase"
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

	roleRepository := role_repository.GetRoleRepository(postgresConn)
	roleUsecase := role_usecase.GetRoleUsecase(roleRepository)
	roleDelivery := role_delivery.GetRoleDelivery(roleUsecase)

	productRepository := product_repository.GetProductRepository(postgresConn)
	productUsecase := product_usecase.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)
	defaultCors := middleware.CORSMiddleware()

	router := gin.Default()
	router.Use(defaultCors)

	router.POST("/user", userDelivery.CreateNewUser)
	router.POST("/login", userDelivery.UserLogin)
	router.GET("/roles", roleDelivery.GetAllRole)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.JWTAuth(jwtUsecase))
	{
		protectedRoutes.GET("/users", userDelivery.GetAllUsers)
		protectedRoutes.GET("/user/:id", userDelivery.GetUserById)
		protectedRoutes.GET("/products", productDelivery.GetAllProducts)
		protectedRoutes.GET("/products/:id", productDelivery.GetProductById)
	}

	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.JWTAuthAdmin(jwtUsecase, userRepository))
	{
		adminRoutes.PUT("/user/:id", userDelivery.UpdateUserData)
		adminRoutes.DELETE("/user/:id", userDelivery.DeleteUserById)
		adminRoutes.POST("/product", productDelivery.CreateNewProduct)
		adminRoutes.PUT("/products/:id", productDelivery.UpdateProductData)
		adminRoutes.DELETE("/product/:id", productDelivery.DeleteProductById)
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
