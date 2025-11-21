package main

import (
	"log"
	"os"
	"strings"
	"context"

	"github.com/joho/godotenv"
	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"

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

	if err := indexerClient.DownloadFragments(ctx, strings.Split(roots, ","), "copy-report.pdf", true); err != nil {
		log.Fatalf("Download file error: %v", err)
	}

	log.Printf("Download file successfully!\n")
	
}
