package main

import (
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/application/usecase"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/infrastructure/persistence/datastore"
	"github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/handler"
	pe "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/entry"
	ps "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/shop"
	pu "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/user"
	pv "github.com/2018-miraikeitai-org/Rakusale-Another-Server/interfaces/server/rpc/vegetable"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listenPortOnAuth, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	listenPortNoAuth, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalln(err)
	}
	serverOnAuth := grpc.NewServer()
	serverNoAuth := grpc.NewServer()
	err := handler.InitGCS()
	if err != nil {
		log.Fatalln(err)
	}
	// [Initialize]
	// User
	userRepo := datastore.NewUserRepository(handler.OpenDBConnection())
	userUseCase := usecase.NewUserUseCase(userRepo)
	pu.RegisterUsersServer(serverOnAuth, userUseCase)
	// Shop
	shopRepo := datastore.NewShopRepository(handler.OpenDBConnection())
	shopUseCase := usecase.NewShopUseCase(shopRepo)
	ps.RegisterShopsServer(serverOnAuth, shopUseCase)
	// Vegetable
	vegetableRepo := datastore.NewVegetableRepository(handler.OpenDBConnection())
	vegetableUseCase := usecase.NewVegetableUseCase(vegetableRepo)
	pv.RegisterVegetablesServer(serverOnAuth, vegetableUseCase)
	// Entry
	entryRepo := datastore.NewEntryRepository(handler.OpenDBConnection())
	entryUseCase := usecase.NewEntryUseCase(entryRepo)
	pe.RegisterEntrysServer(serverNoAuth, entryUseCase)
	// Start Server
	serverOnAuth.Serve(listenPortOnAuth)
	serverNoAuth.Serve(listenPortNoAuth)
}
