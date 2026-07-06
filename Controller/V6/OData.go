package V6

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type ODataPagedResponse[T any] struct {
	Context  string `json:"ODataContext"`
	Count    int64  `json:"ODataCount"`
	Value    []T    `json:"Value"`
	NextLink string `json:"ODataNextLink,omitempty"`
}

func serviceRoot(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/api/v2", scheme, c.Request.Host)
}

func ContextURL(c *gin.Context, entitySet string) string {
	return serviceRoot(c) + "/$metadata#" + entitySet
}

func ToODataResponse[T any](c *gin.Context, result *Repository.PaginationResult[T], entitySet string) ODataPagedResponse[T] {
	var nextLink string
	if result.Skip+result.Top < int(result.Total) {
		q := url.Values{}
		for k, v := range c.Request.URL.Query() {
			q[k] = v
		}
		q.Set("$skip", strconv.Itoa(result.Skip+result.Top))
		nextLink = fmt.Sprintf("%s/%s?%s", serviceRoot(c), entitySet, q.Encode())
	}

	return ODataPagedResponse[T]{
		Context:  ContextURL(c, entitySet),
		Count:    result.Total,
		Value:    result.Data,
		NextLink: nextLink,
	}
}

func ToODataEntityResponse(entity interface{}, contextURL string) (map[string]interface{}, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	m["ODataContext"] = contextURL
	return m, nil
}

var (
	odataOps = map[string]string{
		"eq": "=",
		"ne": "<>",
		"gt": ">",
		"lt": "<",
		"ge": ">=",
		"le": "<=",
	}
	simpleExpr   = regexp.MustCompile(`^(\w+)\s+(eq|ne|gt|lt|ge|le)\s+(.+)$`)
	containsExpr = regexp.MustCompile(`(?i)^contains\((\w+),\s*'([^']*)'\)$`)
	startsExpr   = regexp.MustCompile(`(?i)^startswith\((\w+),\s*'([^']*)'\)$`)
	endsExpr     = regexp.MustCompile(`(?i)^endswith\((\w+),\s*'([^']*)'\)$`)
	andSplit     = regexp.MustCompile(`(?i)\s+and\s+`)
)

func ParseODataFilter(filter string) ([]Repository.Condition, error) {
	if filter == "" {
		return nil, nil
	}

	var conditions []Repository.Condition

	for _, part := range andSplit.Split(filter, -1) {
		part = strings.TrimSpace(part)

		if m := containsExpr.FindStringSubmatch(part); m != nil {
			conditions = append(conditions, Repository.Condition{Field: m[1], Operator: "LIKE", Value: "%" + m[2] + "%"})
			continue
		}
		if m := startsExpr.FindStringSubmatch(part); m != nil {
			conditions = append(conditions, Repository.Condition{Field: m[1], Operator: "LIKE", Value: m[2] + "%"})
			continue
		}
		if m := endsExpr.FindStringSubmatch(part); m != nil {
			conditions = append(conditions, Repository.Condition{Field: m[1], Operator: "LIKE", Value: "%" + m[2]})
			continue
		}
		if m := simpleExpr.FindStringSubmatch(part); m != nil {
			conditions = append(conditions, Repository.Condition{
				Field:    m[1],
				Operator: odataOps[m[2]],
				Value:    strings.Trim(m[3], "'"),
			})
			continue
		}

		return nil, fmt.Errorf("unsupported filter expression: %s", part)
	}

	return conditions, nil
}

func ParseODataOrderBy(orderby string) (sortBy, order string) {
	if orderby == "" {
		return "", ""
	}
	parts := strings.Fields(orderby)
	sortBy = parts[0]
	order = "asc"
	if len(parts) > 1 && strings.EqualFold(parts[1], "desc") {
		order = "desc"
	}
	return
}
