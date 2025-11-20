package main

import (
	"context"

	"github.com/joho/godotenv"


)

func prepare() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main()  {
	prepare()

	log.Printf("begin to init w3client")
	w3client := blockchain.MustNewWeb3(os.Getenv("EVM_RPC"), os.Getenv("PRIVATE_KEY"))
	defer w3client.Close()

	indexerClient, err := indexer.NewClient(os.Getenv("INDEXER_RPC"))
	if err != nil {
		log.Fatalf("create indexer client error: %v", err)
	}

	ctx := context.Background()

	roots := os.Getenv("ROOTS")

	downloader, err := transfer.newDownloader(ctx, )
	if err != nil {
		log.Fatal("Failed to initialize downloader: %v", err)
	}


	log.Printf("Begin to download file ...\n")
	clients := node.MustNewZgsClients(args.nodes, providerOption)
	closer := func() {
		for _, client := range clients {
			client.Close()
		}
	}

	downloader, err := transfer.NewDownloader(clients, common.LogOption{Logger: logrus.StandardLogger()})
	if err != nil {
		closer()
		return nil, nil, err
	}
	downloader.WithRoutines(downloadArgs.routines)
	
}