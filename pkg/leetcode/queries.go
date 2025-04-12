package leetcode

import "fmt"


func (c *Client) GetProblemBySlug(slug string) (*Question, error) {
	query := `
		query getQuestionDetail($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				content
				title
				titleSlug
				difficulty
			}
		}
	`
	variables := map[string]string{"titleSlug": slug}
	resp, err := c.executeQuery(query, variables)
	if err != nil {
		return nil, err
	}
	return &resp.Data.Question, nil
}

func (c *Client) GetProblemByNumber(number int) (*Question, error) {
	query := `
		query problemList($categorySlug: String, $filters: QuestionListFilterInput) {
			questionList(categorySlug: $categorySlug, filters: $filters) {
				data {
					content
					title
					titleSlug
					difficulty
				}
			}
		}
	`
	variables := map[string]interface{}{
		"categorySlug": "all-code-essentials",
		"filters": map[string]interface{}{
			"searchKeywords": fmt.Sprintf("%d", number),
		},
	}
	resp, err := c.executeQuery(query, variables)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.QuestionList.Data) == 0 {
		return nil, fmt.Errorf("no problem found with number %d", number)
	}
	return &resp.Data.QuestionList.Data[0], nil
}

func (c *Client) GetRandomProblem(difficulty string) (*Question, error) {
	query := `
		query problemList($categorySlug: String, $filters: QuestionListFilterInput) {
			questionList(categorySlug: $categorySlug, filters: $filters) {
				data {
					content
					title
					titleSlug
					difficulty
				}
			}
		}
	`
	filters := map[string]interface{}{}
	if difficulty != "" {
		filters["difficulty"] = difficulty
	}
	variables := map[string]interface{}{
		"categorySlug": "all-code-essentials",
		"filters":      filters,
	}
	resp, err := c.executeQuery(query, variables)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.QuestionList.Data) == 0 {
		return nil, fmt.Errorf("no problems found")
	}
	// Simple random selection (could be improved with proper randomization)
	return &resp.Data.QuestionList.Data[0], nil
}