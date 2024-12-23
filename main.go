package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scoop/components"
	"scoop/semantics"
)

var HadError bool = false
var HadRuntimeError bool = false

type Scoop struct {
}

func (s *Scoop) runPrompt() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n>")
		reader.Scan()
		line := reader.Text()
		s.run(line)
		HadError = false
	}
}

func (s *Scoop) runFile(path string) {
	// log.Println("running file...")
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(65)
	}
	s.run(string(bytes))
	if HadError {
		log.Fatal("Errors occured while running...")
		os.Exit(65)
	}

	if HadRuntimeError {
		log.Fatal("Runtime Errors occured while running...")
		os.Exit(70)
	}

}

func (s *Scoop) run(source string) {
	log.Printf("Scanning Args: [ %v ]", source)

	scanner := components.InitScanner(source)
	tokens := scanner.ScanTokens()
	parser := components.InitParser(tokens)

	expression, err := parser.Parse()

	if err != nil {
		fmt.Printf("Error from Parser  : %v", err.Error())
	}

	interpreter := semantics.InitInterpreter()

	interpreter.Interprete(expression)

	printer := semantics.InitAbstractSyntaxTreePrinter()

	if HadError {
		return
	}

	fmt.Printf("Expression Tree : %v", printer.Print(expression))
}

func Report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v : %v", line, where, message)
	HadError = true
}

func Error(line int, message string) {
	Report(line, "", message)
}

func PrintError(token semantics.Token, message string) {
	if token.TokenType == semantics.EOF {
		Report(token.Line, " at end", message)
	} else {
		Report(token.Line, " at "+token.Lexeme+"'", message)
	}
}

func main() {
	runner := Scoop{}
	log.Println("Starting Scoop Interpreter...")
	args := os.Args[1:]
	if len(args) > 1 {
		log.Println("Usage : Scoop [script]...")
	} else if len(args) == 1 {
		runner.runFile(args[0])
	} else {
		runner.runPrompt()
	}
}
