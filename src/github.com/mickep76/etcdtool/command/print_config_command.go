package command

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/mickep76/iodatafmt"
)

func NewPrintConfigCommand() cli.Command {
	return cli.Command{
		Name:  "print-config",
		Usage: "Print configuration",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "sort", Usage: "returns result in sorted order"},
			cli.StringFlag{Name: "format, f", EnvVar: "ETCDTOOL_FORMAT", Value: "JSON", Usage: "Data serialization format YAML, TOML or JSON"},
		},
		Action: func(c *cli.Context) {
			printConfigCommandFunc(c)
		},
	}
}

// printConfigCommandFunc
func printConfigCommandFunc(c *cli.Context) {
	// Load configuration file.
	e := LoadConfig(c)

	// Get data format.
	f, err := iodatafmt.Format(c.String("format"))
	if err != nil {
		log.Fatal(err.Error())
	}

	iodatafmt.Print(e, f)
}
