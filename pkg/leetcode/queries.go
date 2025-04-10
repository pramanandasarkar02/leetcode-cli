package leetcode

import "fmt"



func(c *Client) GetProblemByNumber(number int) (*Question, error) {
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
	if resp.Data.QuestionList.Data != nil {
		return &resp.Data.QuestionList.Data[0], nil		

	}
	return nil, fmt.Errorf("no problem found with number %d", number)
}