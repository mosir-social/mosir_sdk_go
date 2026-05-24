package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mosir-social/mosir_sdk_go"
)

func main() {
	endpoint := "https://beta.mosir.app/api/v1"
	client := mosir_sdk_go.NewClient(endpoint, "", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	post, err := client.GetPost(ctx, "VLO8u7UXqclQ7byjfMEX0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting Discord bot example...")
	fmt.Printf("Fetched post from %s\n", endpoint)
	fmt.Println(post.GetPost.Author.Username)
	fmt.Println(post.GetPost.Content)
}
