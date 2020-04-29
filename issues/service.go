package issues

import (
	"github.com/renato-macedo/issuetracker_go/database"
)

// Service has the usecase methods
type Service struct {
	DB *database.DB
}

// FindIssues return a slice of issues
func (s *Service) FindIssues() ([]*Issue, error) {
	q := `SELECT idIssues, title, description, status, createdAt, name
				FROM ISSUES 
				INNER JOIN USERS
				ON USERS.idUsers = ISSUES.author
	`
	rows, err := s.DB.Conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	issues := make([]*Issue, 0)
	for rows.Next() {

		issue := new(Issue)

		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Status, &issue.CreatedAt, &issue.Author)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, nil
}
