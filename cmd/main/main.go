package main

import (
	"fmt"
	"os"

	"github.com/darragh-downey/secure-c/pkg/codegen"
	"github.com/darragh-downey/secure-c/pkg/lexer"
	"github.com/darragh-downey/secure-c/pkg/optimizer"
	"github.com/darragh-downey/secure-c/pkg/parser"
	"github.com/darragh-downey/secure-c/pkg/semantic"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: secure-c-compiler <source-file>")
		return
	}

	sourceFile := os.Args[1]
	source, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Lexical Analysis
	l := lexer.New(string(source))
	tokens := l.IterateTokens()

	// Parsing
	p := parser.NewParser(tokens)
	ast, parseErr := p.Parse()
	if parseErr != nil {
		fmt.Printf("Parse error: %v\n", parseErr)
		return
	}

	// Semantic Analysis
	semanticAnalyzer := semantic.NewAnalyzer()
	semanticErr := semanticAnalyzer.Analyze(ast)
	if semanticErr != nil {
		fmt.Printf("Semantic error: %v\n", semanticErr)
		return
	}

	// Optimization
	opt := optimizer.NewOptimizer()
	optimizeErr := opt.Optimize(ast)
	if optimizeErr != nil {
		fmt.Printf("Optimization error: %v\n", optimizeErr)
		return
	}

	// Code Generation
	codeGen := codegen.NewCodeGenerator()
	code, codegenErr := codeGen.Generate(ast)
	if codegenErr != nil {
		fmt.Printf("Code generation error: %v\n", codegenErr)
		return
	}

	fmt.Println("Generated Code:")
	fmt.Println(code)
}
