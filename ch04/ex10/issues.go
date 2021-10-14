package main

import (
	"basic-go/ch04/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	month_ago := now.AddDate(0, -1, 0)
	year_ago := now.AddDate(-1, 0, 0)

	var items_less_than_month []*github.Issue
	var items_less_than_year []*github.Issue
	var items_more_than_year []*github.Issue

	for _, item := range result.Items {
		if item.CreatedAt.Sub(month_ago) > 0 {
			items_less_than_month = append(items_less_than_month, item)
		} else if item.CreatedAt.Sub(year_ago) > 0 {
			items_less_than_year = append(items_less_than_year, item)
		} else {
			items_more_than_year = append(items_more_than_year, item)
		}
	}

	fmt.Println("一ヶ月未満")
	for _, item := range items_less_than_month {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("一年未満")
	for _, item := range items_less_than_year {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("一年以上")
	for _, item := range items_more_than_year {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

}
