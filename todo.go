package main

import (
  "fmt"
  "regexp"
  "github.com/fatih/color"
)

type Todo struct {
  Id int64 `json:"id"`
  Desc string `json:"desc"`
  Status string `json:"status"`
  Modified string `json:"modified"`
}

func (t *Todo) MakeOutput() {
  var symbole string
  var colorFunction func(...interface {}) string

  if(t.Status == "done"){
    colorFunction = color.New(color.FgGreen).SprintFunc()
    symbole = "✓"
  } else {
    colorFunction = color.New(color.FgRed).SprintFunc()
    symbole = "✕"
  }

  hashtag_reg := regexp.MustCompile("#[^\\s]*")

  if hashtag_reg.MatchString(t.Desc) {
    hashtag_output := color.New(color.FgYellow).SprintFunc()
    t.Desc = hashtag_reg.ReplaceAllString(t.Desc, hashtag_output(hashtag_reg.FindString(t.Desc)))
  }

  fmt.Println("\t", t.Id, "|", colorFunction(symbole), t.Desc)
}