package main

import (
	"log"
	"swap.io-ledger/src/addressSyncer"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/auth"
	"swap.io-ledger/src/config"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/httpHandler"
	"swap.io-ledger/src/httpServer"
	"swap.io-ledger/src/managers/AddressSyncStatusManager"
	"swap.io-ledger/src/managers/CoinsManager"
	"swap.io-ledger/src/managers/TxsManager"
	"swap.io-ledger/src/managers/UsersAdressesManager"
	"swap.io-ledger/src/managers/UsersManager"
	"swap.io-ledger/src/managers/UsersSpendsManager"
	"swap.io-ledger/src/registrar"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/socketServer"
	"swap.io-ledger/src/txsHandler"
)

func main() {
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(privateKey)

	//data, err := hex.DecodeString("31363338363231323238323731")
	//hash := crypto.Keccak256Hash(data)
	//
	//signature, err := hex.DecodeString("5cfcfaf6a33d99daf66ba54d57506aef92d3bff0f6194b5556b5903207ee833e0fe3716acc43e037a02e2651ecb0ae9537bf52c36d95f12664121ec616031cbd")
	//signature = append(signature, 0)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//sigPublicKeyBytes, err := hex.DecodeString("049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
	//
	//fmt.Println(sigPublicKeyBytes)
	//fmt.Println(hex.EncodeToString(sigPublicKeyBytes))
	//
	//log.Println(crypto.VerifySignature(sigPublicKeyBytes, hash.Bytes(), signature[:64]))
	//
	//fmt.Println()

	config.InitializeConfig()

	registry := serviceRegistry.NewServiceRegistry()

	databaseInstance := database.InitialiseDatabase()
	err := registry.RegisterService(
		databaseInstance,
	)
	if err != nil {
		log.Panicln(err)
	}

	TxsManager.Register(registry)
	CoinsManager.Register(registry)
	UsersManager.Register(registry)
	UsersAdressesManager.Register(registry)
	UsersSpendsManager.Register(registry)
	AddressSyncStatusManager.Register(registry)

	txsHandler.Register(registry)

	hsd := config.AGENTS[0]
	err = AgentHandler.Register(
		registry,
		hsd.Network,
		hsd.BaseUrl,
		hsd.ApiKey,
	)
	if err != nil {
		log.Panicln(err)
	}

	AddressSyncer.Register(registry)
	registrar.Register(registry)
	auth.Register(registry)

	socketServer.Register(registry)
	HttpServer.Register(registry)

	err = registry.RegisterService(
		httpHandler.InitializeServer(),
	)
	if err != nil {
		log.Panicln(err)
	}

	registry.StartAll()

	<-make(chan struct{})
}
