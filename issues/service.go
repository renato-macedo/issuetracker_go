package issues

import (
	"github.com/renato-macedo/issuetracker_go/database"
)

// Service has the usecase methods
type Service struct {
	DB *database.DB
}

// FindIssues return a slice of open issues, set arg to true to include closed
func (s *Service) FindIssues(includeClosed bool) ([]*Issue, error) {
	q := `SELECT idIssues, title, description, status, createdAt, name
				FROM ISSUES 
				INNER JOIN USERS
				ON USERS.idUsers = ISSUES.author`

	if includeClosed == false {
		q += ` WHERE status = 'OPEN'`
	}

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

// CreateIssue inserts into database a issue with status "OPEN" and fill createdAt with the Now() function
func (s *Service) CreateIssue(issue *IssueDTO) error {
	q := `
	INSERT INTO ISSUES (title, description, author, status, createdAt) values (?, ?, ?, "OPEN", Now())
	`
	_, err := s.DB.Conn.Exec(q, issue.Title, issue.Description, issue.AuthorID)
	return err
}
