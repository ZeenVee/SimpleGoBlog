package main

import (
	"net/http"
	"testing"
)

func TestArticleValidityCode(t *testing.T) {
	tables := []struct {
		inText  string
		inOrder uint
		outCode int
	}{
		{"", 0, 0},
		{"", 1, 0},
		{"", 2, 0},
		{"1", 0, 1},
		{"asdsd sds", 2, 4},
		{"1", 1, 2},
		{"asdsd", 2, 4},
		{"asdsd", 2, 4},
		{"asdsd", 3, 8},
		{"bdf sdf ", 4, 16},
		{"", 5, 0},
		{" ", 5, 0},
		{" ttrtr", 5, 32},
		{" ttr tr", 5, 32},
		{"\t", 5, 0},
		{"bdf sdf ", 5, 32},
	}

	for _, table := range tables {
		result := articleValidityCode(table.inText, table.inOrder)

		if result != table.outCode {
			t.Errorf("articleValidityCode(%v, %v) was incorrect, got %v, want %v.",
				table.inText, table.inOrder, result, table.outCode)
		}
	}
}

func TestArticleValidityResponse(t *testing.T) {
	tables := []struct {
		inArticle   Article
		outStatus   int
		outMessage  string
		outValidity bool
	}{
		{Article{
			ID:      0,
			Title:   " ",
			Content: "",
			Author:  ""},
			http.StatusLengthRequired,
			"Empty fields: title, content, author",
			false,
		},
		{Article{
			ID:      1,
			Title:   "1",
			Content: "",
			Author:  ""},
			http.StatusLengthRequired,
			"Empty fields: content, author",
			false,
		},
		{Article{
			ID:      1,
			Title:   "Hello",
			Content: "Hello Go",
			Author:  "Go"},
			http.StatusOK,
			"Success",
			true,
		},
		{Article{
			ID:      7,
			Title:   "",
			Content: "Hello Go",
			Author:  ""},
			http.StatusLengthRequired,
			"Empty fields: title, author",
			false,
		},
	}

	for _, table := range tables {
		resultBody, resultValidity := articleValidityResponse(table.inArticle)

		if resultBody.Status != table.outStatus {
			t.Errorf("articleValidityResponse(%v) Status was incorrect, got %v, want %v.",
				table.inArticle, resultBody.Status, table.outStatus)
		}

		if resultBody.Message != table.outMessage {
			t.Errorf("articleValidityResponse(%v) Message was incorrect, got %v, want %v.",
				table.inArticle, resultBody.Message, table.outMessage)
		}

		if resultValidity != table.outValidity {
			t.Errorf("articleValidityResponse(%v) validity was incorrect, got %v, want %v.",
				table.inArticle, resultValidity, table.outValidity)
		}
	}
}
