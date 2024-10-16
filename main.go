package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scoop/components"
)

type Scoop struct {
	hadError        bool
	hadRuntimeError bool
}

func (s *Scoop) runPrompt() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		reader.Scan()
		line := reader.Text()
		s.run(line)
		s.hadError = false
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
	if s.hadError {
		log.Fatal("Errors occured while running...")
		os.Exit(65)
	}

	if s.hadRuntimeError {
		log.Fatal("Runtime Errors occured while running...")
		os.Exit(70)
	}

}

func (s *Scoop) run(source string) {
	log.Printf("Scanning Args: [ %v ]", source)

	scanner := components.InitScanner(source)
	tokens := scanner.ScanTokens()
	parser := components.InitParser(tokens)

	if s.hadError {
		return
	}
}

func (s *Scoop) report(line int, where string, message string) {
	fmt.Printf("[line %v] Error %v : %v", line, where, message)
	s.hadError = true
}

func (s *Scoop) error(line int, message string) {
	s.report(line, "", message)
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
