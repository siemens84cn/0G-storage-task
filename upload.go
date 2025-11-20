package main

import (
	"log"
	"os"
	"strings"
	"context"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/0gfoundation/0g-storage-client/core"
	"github.com/0gfoundation/0g-storage-client/transfer"
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
	nodes, err := indexerClient.SelectNodes(ctx, 1, []string{
		"http://34.174.223.105:5678",
		"http://104.196.238.89:5678",
		"http://34.57.99.219:5678",
		"http://34.55.197.204:5678",
		"http://34.133.200.179.5678",
	}, "max", true)
	if err != nil {
		log.Fatalf("select nodes error: %v", err)
	}

	uploader, err := transfer.NewUploader(ctx, w3client, nodes)
	if err != nil {
		log.Fatalf("create uploader error: %v", err)
	}

	file, err := core.Open(os.Getenv("FILE_NAME"))
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	log.Printf("Begin to upload file ...\n")
	fragmentSizeStr := os.Getenv("FRAGMENT_SIZE")
	fragmentSize, err := strconv.ParseInt(fragmentSizeStr, 10, 64)
	if err != nil {
		log.Fatalf("Error Fragment size: %v", err)
	}

	_, roots, err := uploader.SplitableUpload(ctx, file, fragmentSize, transfer.UploadOption{
		SkipTx:				true,
		FinalityRequired:	transfer.FileFinalized,
		FullTrusted:		false,
		NRetries:			10,
		TaskSize:			10,
		Method:				"10",
	})
	if err != nil {
		log.Fatalf("Upload file error: %v", err)
	}

	log.Printf("Upload file successfully!\n")

	log.Printf("Roots size: %d\n", len(roots))
	s := make([]string, len(roots))
	for i, root := range roots {
		s[i] = root.String()
	}
	log.Printf("File uploaded in %v fragments, roots = %v", len(roots), strings.Join(s, ","))
}