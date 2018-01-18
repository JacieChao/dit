package imp

import (
	"fmt"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)



func Run(image string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	historydata, err := cli.ImageHistory(context.Background(), image)
	for _, his := range historydata {
		fmt.Printf("history is %+v \n", his)
	}
}
