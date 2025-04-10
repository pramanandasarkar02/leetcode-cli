package main

import (
	"fmt"

	"github.com/pramananadasarkar02/leetcode-local-cli/cmd"
	"github.com/pramananadasarkar02/leetcode-local-cli/config"
	"github.com/pramananadasarkar02/leetcode-local-cli/pkg/leetcode"
)


func main(){
	
	

	//  set command line flags 

	//  Initialize configuration

	OutputFileName := "problems/problem.md"

	cfg := config.NewConfig(OutputFileName)

	


	//  Initialize clients
	client := leetcode.NewClient()


	//  Execute command 

	err := cmd.FetchProblemByNumber(client, cfg, 12)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Problem saved to %s\n", cfg.OutputFileName)







}