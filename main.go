package hoover

import (
	"github.com/docker/docker/client"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"log"
)

func main() {
	cleanupDockerImages()
}

func cleanupDockerImages() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print("Docker client couldn't be initialized.")
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Print("Not able to get List of Docker Images")
	}
	for _, image := range images {
		cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{Force: true})
		fmt.Println("Removed Image: " + image.ID)
	}

	fmt.Println("Removed all Docker Images")
}
