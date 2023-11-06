package funcs

import "fmt"

func Help() {
	fmt.Print(
		`Usage of Interpreter:
  
  Subcommands:
   *  [--input] <inputPath>:
      This is alias to run command.

   *  2ast [--input] <inputPath>:
      Parse the input file and output the AST json to stdout.

   *  ast [--input] <inputPath> [--output] <outputPath>:
      Parse the input file and output the AST json to the output file.

   *  compile [--input] <inputPath> [--output <outputPath>]:
      Parse the input file and output the compressed ast to the output file.
      If no output file is specified, the output file will be the same name
      as the input file, but with the extension changed to .ast.

   *  run [--input] path:
      Parse the input file and run it.

   *  repl [--input <path>]:
      Enter the REPL mode.

   *  help:
      Show this help message.
`)
}
