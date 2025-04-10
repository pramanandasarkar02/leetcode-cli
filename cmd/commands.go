package cmd

import (
	"os"

	"github.com/pramananadasarkar02/leetcode-local-cli/config"
	"github.com/pramananadasarkar02/leetcode-local-cli/pkg/leetcode"
)





func FetchProblemByNumber(client *leetcode.Client, cfg *config.Config, number int) error {
	question, err := client.GetProblemByNumber(number)
	if err != nil {
		return err
	}
	return saveQuestion(question, cfg)
}

func saveQuestion(question *leetcode.Question, cfg *config.Config) error {
	// save question to file
	content := question.Content
	

	return os.WriteFile(cfg.OutputFileName, []byte(content), 0644)
}