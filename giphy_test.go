package gogiphy

import (
	"fmt"
	"testing"
)

var client *Client

func setup() *Client {
	if client != nil {
		return client
	}
	client = NewClient("dc6zaTOxFJmzC")
	return client
}

func TestSearch(t *testing.T) {
	client = setup()
	result, err := client.Search("lol", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) == 0 {
		t.Fatal("Search retured zero results")
	}
	t.Log(fmt.Sprintf("Number of gifs found = %d", len(result.Data)))
}

func TestSearchLimit(t *testing.T) {
	client := setup()
	params := &SearchParams{
		Limit: 10,
	}
	result, err := client.Search("lol", params)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) != 10 {
		t.Fatal("Search did not return 10 items")
	}
	t.Log(fmt.Sprintf("Number of gifs found = %d", len(result.Data)))
}

func TestSearchPagination(t *testing.T) {
	client = setup()
	result, err := client.Search("lol", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) == 0 {
		t.Fatal("Search retured zero results")
	}
	t.Log(fmt.Sprintf("Pagination.\nTotal = %d\nCount = %d\nOffset = %d\n", result.Pagination.TotalCount, result.Pagination.Count, result.Pagination.Offset))
}
