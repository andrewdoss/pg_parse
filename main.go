package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	pg_query "github.com/pganalyze/pg_query_go"
)

// Provisioning for future options
func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	var queryStr string
	switch len(args) {
	case 0:
		queryBytes, err := ioutil.ReadAll(os.Stdin)
		queryStr = string(queryBytes)
		if err != nil {
			fmt.Printf("expected query string from stdin: %v", err)
			os.Exit(1)
		}
	case 1:
		queryStr = args[0]
	default:
		fmt.Println("expected at most one argument containing a query string")
		os.Exit(1)
	}

	parse(queryStr)
}

func parse(queryStr string) {
	tree, err := pg_query.ParseToJSON(queryStr)
	if err != nil {
		fmt.Printf("unable to parse provided query to json: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", tree)
}
