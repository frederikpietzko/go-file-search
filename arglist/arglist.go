package arglist

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SearchStrategy = int

const (
	DFS SearchStrategy = iota
	BFS
)

type Args struct {
	Debug      bool
	Cwd        string
	Recursive  bool
	SearchTerm string
	Strategy   SearchStrategy
}

func isPossibleSearchTerm(index int, element string, lastArg string) bool {
	return index != 0 && !strings.HasPrefix(element, "-") && !strings.HasPrefix(lastArg, "--")
}

func New() *Args {
	args := Args{}
	var lastArg string
	for index, element := range os.Args {
		if element == "-d" {
			args.Debug = true
		} else if lastArg == "--cwd" {
			args.Cwd = element
		} else if element == "-r" {
			args.Recursive = true
		} else if isPossibleSearchTerm(index, element, lastArg) {
			args.SearchTerm = strings.ToLower(strings.Trim(element, " "))
		} else if element == "-dfs" {
			args.Strategy = DFS
		} else if element == "-bfs" {
			args.Strategy = BFS
		}
		lastArg = element
	}

	if args.Cwd == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		args.Cwd = cwd
	}

	return &args
}

func (args Args) PrintArgs() {
	fmt.Println(args)
}
