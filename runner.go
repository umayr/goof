package main
import (
	"fmt"
	"goof/plugs"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	var (
		plug string
		debug bool
	)

	app := cli.NewApp()
	app.Name = "goof"
	app.Usage = "Extracts blogs from top web magazines"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "plug",
			Value: "",
			Usage: "Tells which plug needs to be invoked",
			Destination: &plug,
		},
		cli.BoolFlag{
			Name: "debug",
			Usage: "Turns on the debug mode",
			Destination: &debug,
		},
	}
	app.Action = func(c *cli.Context) {
		if debug {
			os.Setenv("DEBUG", "*")
		}

		switch plug {
		case "tech-crunch":
			t := plugs.NewTechCrunch()
			posts := t.Next()

			for i := range posts {
				fmt.Printf("%s\n", posts[i].Json())
			}
			break
		default:
			panic("invalid plug name")
		}

	}
	app.Run(os.Args)
}