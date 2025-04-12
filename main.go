package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pramananadasarkar02/leetcode-local-cli/cmd"
	"github.com/pramananadasarkar02/leetcode-local-cli/config"
	"github.com/pramananadasarkar02/leetcode-local-cli/pkg/leetcode"
)


func main(){
	
	

	//  set command line flags 
	slug := flag.String("slug", "", "Leetcode problem slug(e.g. two-sum)")
	number := flag.Int("number", 0, "Leetcode problem number")
	difficulty := flag.String("difficulty", "", "Filter by difficulty (easy, medium, hard)")
	random := flag.Bool("random", false, "Get a random problem")
	output := flag.String("output", "problem.txt", "Output file name")
	flag.Parse()

	//  Initialize configuration
	

	cfg := config.NewConfig(*output)

	


	//  Initialize clients
	client := leetcode.NewClient()


	//  Execute command 
	var err error
	switch {
	case *random:
		err = cmd.FetchRandomProblem(client, cfg)
	case *number > 0:
		err = cmd.FetchProblemByNumber(client, cfg, *number)
	case *difficulty != "":
		err = cmd.FetchProblemByDifficulty(client, cfg, *difficulty)
	case *slug != "":
		err = cmd.FetchProblemBySlug(client, cfg, *slug)
	default:
		fmt.Println("Error: Please specify a fetch option")
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	


	

	






}