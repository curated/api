package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/glog"

	"github.com/curated/elastic/config"
	"github.com/labstack/echo"
	"github.com/olivere/elastic"
)

const (
	batchSize  = 10
	issueIndex = "issue"
)

var issueSortOptions = map[string]bool{
	"thumbsUp":       true,
	"thumbsDown":     true,
	"laugh":          true,
	"hooray":         true,
	"confused":       true,
	"heart":          true,
	"repoForks":      true,
	"repoStargazers": true,
	"createdAt":      true,
	"updatedAt":      true,
}

// Server serves HTTP requests
type Server struct {
	Context context.Context
	Client  *elastic.Client
	Echo    *echo.Echo
	Config  *config.Config
}

// IssuesRequest structure
type IssuesRequest struct {
	Sort string
	Asc  bool
	From int
}

// IssuesResponse structure
type IssuesResponse struct {
	Total  int64         `json:"total"`
	Issues []interface{} `json:"issues"`
}

// New creates the server
func New() *Server {
	cfg := config.New()

	cli, err := elastic.NewClient(
		elastic.SetURL(cfg.Elastic.URL),
		elastic.SetBasicAuth(cfg.Elastic.Username, cfg.Elastic.Password),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)

	if err != nil {
		glog.Fatalf("Failed creating server: %v", err)
	}

	s := &Server{
		Context: context.Background(),
		Client:  cli,
		Echo:    echo.New(),
		Config:  cfg,
	}

	s.Echo.GET("/issues", s.searchIssues)
	return s
}

// Start initializes the server
func (s *Server) Start() {
	glog.Fatal(s.Echo.Start(":1323"))
}

func (s *Server) searchIssues(c echo.Context) error {
	req, err := s.parseIssueRequest(c)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sr, err := s.Client.Search().
		Index(issueIndex).
		Sort(req.Sort, req.Asc).
		From(req.From).
		Size(batchSize).
		Do(s.Context)

	if err != nil {
		glog.Errorf("Failed searching issues: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	res := IssuesResponse{
		Total: sr.Hits.TotalHits,
	}

	for _, hit := range sr.Hits.Hits {
		var issue map[string]interface{}
		err := json.Unmarshal(*hit.Source, &issue)

		if err != nil {
			glog.Errorf("Failed parsing issue: %v\n%s", err, string(*hit.Source))
			return c.NoContent(http.StatusInternalServerError)
		}

		res.Issues = append(res.Issues, issue)
	}

	return c.JSON(http.StatusOK, res)
}

func (s *Server) parseIssueRequest(c echo.Context) (*IssuesRequest, error) {
	sort := c.QueryParam("sort")
	if !issueSortOptions[sort] {
		return nil, fmt.Errorf("Cannot sort by '%s'", sort)
	}

	asc, err := strconv.ParseBool(c.QueryParam("asc"))
	if err != nil {
		return nil, fmt.Errorf("Invalid boolean value for 'asc' param")
	}

	from, err := strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return nil, fmt.Errorf("Invalid integer value for 'from' param")
	}

	return &IssuesRequest{
		Sort: sort,
		Asc:  asc,
		From: from,
	}, nil
}
