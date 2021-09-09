package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/matheustex/football-data-sdk"
)

func main() {
	httpClient := &http.Client{}
	ctx := context.Background()

	client := football.NewClient(httpClient)

	player, _ := client.Players.Find(ctx, "36")

	fmt.Println(player.ID)
	fmt.Println(player.Name)
}
