package main

import (
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/application/usecase"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/infrastructure/persistence/datastore"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/handler"
	pe "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/entry"
	ps "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/shop"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	// err = handler.InitGCS()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	DB := handler.OpenDBConnection()
	defer DB.Close()

	logger := logrus.New()
	// [Initialize]
	// User
	userRepo := datastore.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	pu.RegisterUsersServer(server, userUseCase)
	// Shop
	shopRepo := datastore.NewShopRepository(DB)
	shopUseCase := usecase.NewShopUseCase(shopRepo)
	ps.RegisterShopsServer(server, shopUseCase)
	// Vegetable
	vegetableRepo := datastore.NewVegetableRepository(DB, logger)
	vegetableUseCase := usecase.NewVegetableUseCase(vegetableRepo, logger)
	pv.RegisterVegetablesServer(server, vegetableUseCase)
	// Entry
	entryRepo := datastore.NewEntryRepository(DB)
	entryUseCase := usecase.NewEntryUseCase(entryRepo)
	pe.RegisterEntrysServer(server, entryUseCase)
	// Start Server
	server.Serve(listenPort)
}
