package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// AllTutorials returns all tutorials, with optional query filters.
//
// Supported query parameters:
//   - category : filter tutorials by category (e.g., "go", "javascript")
//   - author   : filter tutorials by author name
//   - level    : filter by difficulty level (e.g., "beginner", "advanced")
//   - search   : full-text keyword search in title/description
//   - limit    : number of tutorials to return (pagination)
//   - sort     : sorting order (e.g., "date", "popularity")
//   - page     : page number for pagination
//
// Example requests:
//   GET /tutorials
//   GET /tutorials?category=go
//   GET /tutorials?category=go&level=beginner
//
// The function currently returns mocked data and the filters applied.
// Replace the mock logic with database queries as needed.

func AllTutorials(c *gin.Context) {
	// Read optional query parameters
	category := c.Query("category")
	author := c.Query("author")
	level := c.Query("level")
	limit := c.Query("limit")
	sort := c.Query("sort")
	page := c.Query("page")
	search := c.Query("search")

	// Build a map of the applied filters
	filters := gin.H{}

	if category != "" {
		filters["category"] = category
	}
	if author != "" {
		filters["author"] = author
	}
	if level != "" {
		filters["level"] = level
	}
	if search != "" {
		filters["search"] = search
	}
	if limit != "" {
		filters["limit"] = limit
	}
	if sort != "" {
		filters["sort"] = sort
	}
	if page != "" {
		filters["page"] = page
	}

	// Respond with mock data for now
	c.JSON(http.StatusOK, gin.H{
		"tutorials": "Filtered tutorials",
		"filters":   filters,
	})
}
