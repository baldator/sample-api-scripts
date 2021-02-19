package main

import (
	"context"
	"fmt"
	"os"

	"code.vegaprotocol.io/vega/proto/api"
	"google.golang.org/grpc"
)

func main() {
	nodeURLGrpc := os.Getenv("NODE_URL_GRPC")
	if len(nodeURLGrpc) == 0 {
		panic("NODE_URL_GRPC is null or empty")
	}

	conn, err := grpc.Dial(nodeURLGrpc, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dataClient := api.NewTradingDataServiceClient(conn)
	request := api.AssetsRequest{}
	assets, err := dataClient.Assets(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	// Find asset with name DAI
	assetFound := false
	var assetId string
	for _, asset := range assets.Assets {
		fmt.Printf("Assets: %s \n", asset.Name)
		fmt.Printf("Assets: %s \n", asset.Symbol)
		if asset.Symbol == "tDAI" {
			fmt.Println("Found an asset with name tDAI:")
			assetId = asset.Id
			assetFound = true
			break
		}
	}

	if !assetFound {
		panic("tDAI asset not found on specified Vega network, please propose and create the tDAI asset")
	}

	// Request a single asset by identifier on a Vega network
	requestAsset := api.AssetByIDRequest{Id: assetId}
	assetObject, err := dataClient.AssetByID(context.Background(), &requestAsset)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Asset by id: %s", assetObject.Asset)
}
