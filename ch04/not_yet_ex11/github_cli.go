package main

import (
	"basic-go/ch04/ex11/github"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"

)

func main(){
	createIssue("a", "bb")





}


// curl -u platypus:$git_token \
//   -X POST \
//   -H "Accept: application/vnd.github.v3+json" \
//   https://api.github.com/repos/platypus-12/basic-go/issues \
//   -d '{"title":"title"}' 
func createIssue(title, body string){
	issue:= github.Issue{Title: title, Body: body}
	fmt.Println(issue)
	issue_json, _ := json.Marshal(issue)
	fmt.Println(string(issue_json))

	req, _ := http.NewRequest("POST", "https://api.github.com/repos/platypus-12/basic-go/issues", 
		bytes.NewBuffer(issue_json))

	client := new(http.Client)
	resp, err := client.Do(req)
	fmt.Println(resp, "||||",err)

	var result github.Issue
	jsonNewDecoder(resp)
}