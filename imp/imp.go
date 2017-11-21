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
	fmt.Printf("show history : %+v", historydata)
	//fmt.Printf("My name is %s, My age is %d", name, age)
	//fmt.Printf("Image id is %s", image)
}
