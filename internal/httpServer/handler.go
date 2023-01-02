package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/pkg/storage"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	userHTTP "github.com/rinatkh/test_2022/internal/users/delivery/http"
	usersRepository "github.com/rinatkh/test_2022/internal/users/repository"
	usersUsecase "github.com/rinatkh/test_2022/internal/users/usecase"

	productsHTTP "github.com/rinatkh/test_2022/internal/products/delivery/http"
	productsRepository "github.com/rinatkh/test_2022/internal/products/repository"
	productsUsecase "github.com/rinatkh/test_2022/internal/products/usecase"

	currencyRepository "github.com/rinatkh/test_2022/internal/currency/repository"
	currencyUsecase "github.com/rinatkh/test_2022/internal/currency/usecase"

	orderProductsRepository "github.com/rinatkh/test_2022/internal/OrderProducts/repository"
	orderProductsUsecase "github.com/rinatkh/test_2022/internal/OrderProducts/usecase"

	ordersHTTP "github.com/rinatkh/test_2022/internal/orders/delivery/http"
	ordersRepository "github.com/rinatkh/test_2022/internal/orders/repository"
	ordersUsecase "github.com/rinatkh/test_2022/internal/orders/usecase"

	friendsHTTP "github.com/rinatkh/test_2022/internal/friends/delivery/http"
	friendsRepository "github.com/rinatkh/test_2022/internal/friends/repository"
	friendsUsecase "github.com/rinatkh/test_2022/internal/friends/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	postgreConnection, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	userRepo := usersRepository.NewPostgresRepository(postgreConnection, s.log)
	productRepo := productsRepository.NewPostgresRepository(postgreConnection, s.log)
	currencyRepo := currencyRepository.NewPostgresRepository(postgreConnection, s.log)
	orderProductsRepo := orderProductsRepository.NewPostgresRepository(postgreConnection, s.log)
	ordersRepo := ordersRepository.NewPostgresRepository(postgreConnection, s.log)
	friendsRepo := friendsRepository.NewPostgresRepository(postgreConnection, s.log)

	userUC := usersUsecase.NewUserUC(s.cfg, s.log, userRepo)
	currencyUC := currencyUsecase.NewProductUC(s.cfg, s.log, currencyRepo)
	productUC := productsUsecase.NewProductUC(s.cfg, s.log, productRepo, currencyUC)
	orderProductsUC := orderProductsUsecase.NewOrderProductsUC(s.cfg, s.log, orderProductsRepo, productUC)
	ordersUC := ordersUsecase.NewOrderUC(s.cfg, s.log, ordersRepo, currencyUC, orderProductsUC, productUC)
	friendsUC := friendsUsecase.NewFriendsUC(s.cfg, s.log, friendsRepo, userUC)

	userHandler := userHTTP.NewUserHandler(userUC, s.log)
	productHandler := productsHTTP.NewProductHandler(productUC, s.log)
	ordersHandler := ordersHTTP.NewOrderHandler(ordersUC, s.log)
	friendsHandler := friendsHTTP.NewFriendHandler(friendsUC, s.log)

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	userHTTP.MapUserRoutes(app, userHandler)
	productsHTTP.MapProductRoutes(app, productHandler)
	ordersHTTP.MapOrderRoutes(app, ordersHandler)
	friendsHTTP.MapFriendsRoutes(app, friendsHandler)

	return nil
}
