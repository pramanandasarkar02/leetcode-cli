package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pramananadasarkar02/leetcode-local-cli/config"
	"github.com/pramananadasarkar02/leetcode-local-cli/pkg/leetcode"
)





func FetchProblemBySlug(client *leetcode.Client, cfg *config.Config, slug string) error {
	question, err := client.GetProblemBySlug(slug)
	if err != nil {
		return err
	}
	return saveQuestion(question, cfg)
}

func FetchProblemByNumber(client *leetcode.Client, cfg *config.Config, number int) error {
	question, err := client.GetProblemByNumber(number)
	if err != nil {
		return err
	}
	return saveQuestion(question, cfg)
}

func FetchProblemByDifficulty(client *leetcode.Client, cfg *config.Config, difficulty string) error {
	difficulty = strings.ToUpper(difficulty)
	if difficulty != "EASY" && difficulty != "MEDIUM" && difficulty != "HARD" {
		return fmt.Errorf("invalid difficulty: must be easy, medium, or hard")
	}
	question, err := client.GetRandomProblem(difficulty)
	if err != nil {
		return err
	}
	return saveQuestion(question, cfg)
}

func FetchRandomProblem(client *leetcode.Client, cfg *config.Config) error {
	question, err := client.GetRandomProblem("")
	if err != nil {
		return err
	}
	return saveQuestion(question, cfg)
}

func saveQuestion(q *leetcode.Question, cfg *config.Config) error {
	content := cleanHTML(q.Content)
	finalContent := fmt.Sprintf("Title: %s\nDifficulty: %s\n\nDescription:\n%s",
		q.Title,
		q.Difficulty,
		content,
	)
	return os.WriteFile(cfg.OutputFileName, []byte(finalContent), 0644)
}

func cleanHTML(content string) string {
	content = strings.ReplaceAll(content, "<p>", "\n")
	content = strings.ReplaceAll(content, "</p>", "\n")
	content = strings.ReplaceAll(content, "<br>", "\n")
	content = strings.ReplaceAll(content, "<strong>", "")
	content = strings.ReplaceAll(content, "</strong>", "")
	
	for strings.Contains(content, "<") && strings.Contains(content, ">") {
		start := strings.Index(content, "<")
		end := strings.Index(content, ">")
		if start < end {
			content = content[:start] + content[end+1:]
		}
	}
	
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n\n\n", "\n\n")
	return content
}