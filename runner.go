package main
import (
	"fmt"
	"os"
	"goof/plugs"
	"github.com/codegangsta/cli"
	"goof/util"
)

func main() {
	var (
		plug string
		debug bool
		page int
	)

	app := cli.NewApp()
	app.Name = "goof"
	app.Usage = "Extracts blogs from top web magazines"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name: "page",
			Value: 1,
			Usage: "Tells how many pages need to be extracted",
			Destination: &page,
		},
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

		db := util.NewDb()
		defer db.Close()

		Posts := db.Posts()

		switch plug {
		case "tech-crunch":
			t := plugs.NewTechCrunch()

			for p := 1; p <= page; p++ {
				posts := t.Next()


				for i := range posts {
					err := Posts.Insert(&posts[i])
					if err != nil {
						panic(err)
					} else {
						fmt.Printf("%s\n", posts[i].Json())
					}
				}
			}
			break
		default:
			panic("invalid plug name")
		}

	}
	app.Run(os.Args)
}