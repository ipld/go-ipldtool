package schema

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	schemadmt "github.com/ipld/go-ipld-prime/schema/dmt"
	schemadsl "github.com/ipld/go-ipld-prime/schema/dsl"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"

	"github.com/ipld/go-ipldtool/app/shared"
)

var Cmd_Schema = &cli.Command{
	Name:     "schema",
	Category: "Advanced",
	Usage:    "Manipulate schemas -- parsing, compiling, transforming, and storing.",
	Subcommands: []*cli.Command{{
		Name:  "parse",
		Usage: "Parse a schema DSL document, and produce the DMT form, emitted in JSON by default.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "no-compile",
				Usage: `Skip the compilation phase, and just emit the DMT (regardless of whether it's logically valid).`,
			},
			&cli.BoolFlag{
				Name:  "save",
				Usage: `Put the parsed schema into storage, and return a CID pointing to it.  (Roughly equivalent to piping the schema parse command into a put command.)`,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       `Defines what format the DMT should be produced in.  Valid arguments are codecs, specified as the word "codec:" followed by a multicodec name, or "codec:0x" followed by a multicodec indicator number in hexidecimal.`,
				DefaultText: "codec:json",
			},
		},
		Action: Action_SchemaParse,
	}, {
		Name:  "compile",
		Usage: "Compile a schema DMT document, exiting nonzero and reporting errors if anything is logically invalid.",
	}, {
		Name:  "codegen",
		Usage: "Generate code for working with IPLD schemas",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "generator",
				Usage:    "Generator to be used for creating the code. Currently supports (go-gengo)",
				Required: true,
			},
			&cli.PathFlag{
				Name:  "output",
				Usage: "Directory where the codegen files should be output to",
				Value: "ipldsch",
			},
			&cli.StringFlag{
				Name:  "package",
				Usage: "Package name for generated files",
				Value: "ipldsch",
			},
		},
		Action: Action_GoCodegen,
	}},
	// Someday: it may be neat to have a handful of well-known transforms, like: strip all rename directives, or make all representations default, etc.
}

// Action_SchemaParse is the function that implements the `ipld schema parse` subcommand's behaviors.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- for incomprehensible or invalid arguments.
//   - schema-dsl-parse-failed -- if the DSL document didn't parse.
//   - schema-compile-failed -- if the schema was parsed, but was logically invalid.
//
func Action_SchemaParse(args *cli.Context) error {
	// Parse positional args.
	var sourceArg string
	switch args.Args().Len() {
	case 1:
		sourceArg = args.Args().Get(0)
	default:
		return shared.ErrInvalidArgs("schema parse command needs exactly one positional argument", nil)
	}

	// Let's get some data!
	inputReader, _, err := shared.ParseDataSourceArg(sourceArg)
	if err != nil {
		return err
	}

	// Parse!
	dmt, err := DSLParse(sourceArg, inputReader)
	if err != nil {
		return err
	}

	// Compile!  Maybe.  Just to make sure we can.
	if !args.Bool("no-compile") {
		_, err = SchemaCompile(dmt)
		if err != nil {
			return err
		}
	}

	// Regard the DMT as a node (which we'll need for either printout or for saving it).
	dmtn := bindnode.Wrap(dmt, schemadmt.Type.Schema.Type())

	// Figure out the output format.
	encoder, err := shared.ParseEncoderArg(args.String("output"), "codec:json", "output")
	if err != nil {
		return err
	}

	// Print out the DMT.
	// TODO: or do something else if the "save" flag is set.
	return ipld.EncodeStreaming(args.App.Writer, dmtn, encoder)
}

func Action_GoCodegen(args *cli.Context) error {
	if args.NArg() != 1 {
		return fmt.Errorf("invalid number of arguments")
	}

	s, err := schemadsl.ParseFile(args.Args().First())
	if err != nil {
		return err
	}

	var ts schema.TypeSystem
	ts.Init()
	if err := schemadmt.Compile(&ts, s); err != nil {
		return err
	}

	generator := args.Path("generator")
	outputDir := args.Path("output")
	pkgName := args.String("package")

	if generator != "go-gengo" {
		return fmt.Errorf("unsupported generator: %s", generator)
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	a := gengo.AdjunctCfg{}
	gengo.Generate(outputDir, pkgName, ts, &a)

	return nil
}
