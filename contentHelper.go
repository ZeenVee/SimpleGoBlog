package main

import (
	"net/http"
	"strings"
)

//this function is to check if there is any empty field of article inserted
//returns the corresponding Response Status, Message and data
//returns false (invalid) if there there is any empty field(s)
func articleValidityResponse(article Article) (ResponseBody, bool) {
	code := articleValidityCode(article.Title, 0) + articleValidityCode(article.Content, 1) +
		articleValidityCode(article.Author, 2)

	rbody := ResponseBody{
		Status:  http.StatusLengthRequired,
		Message: "Empty fields: ",
	}

	switch code {
	case 0:
		rbody.Message += "title, content, author"
		return rbody, false
	case 1:
		rbody.Message += "content, author"
		return rbody, false
	case 2:
		rbody.Message += "title, author"
		return rbody, false
	case 3:
		rbody.Message += "author"
		return rbody, false
	case 4:
		rbody.Message += "title, content"
		return rbody, false
	case 5:
		rbody.Message += "content"
		return rbody, false
	case 6:
		rbody.Message += "title"
		return rbody, false
	}

	return ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
	}, true
}

//this function is to check if the input is not empty
//returns 0 if input is 0 length string or space only string
//returns 2^order if input is not empty
func articleValidityCode(text string, order uint) int {
	if len(strings.TrimSpace(text)) > 0 {
		return (1 << order)
	}

	return 0
}
