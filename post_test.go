package main

import (
	"testing"
)

var (
	raw = `<!-- 
title:	this is title
summary: this summary
tag: java,python,golang
slug: hello-boy-google 
Time: 2022-05-13
-->

# Hello World`
)

func TestParseToPost(t *testing.T) {
	post := ParseToPost([]byte(raw), "/index.md")

	//Title
	if post.Title != "this is title" {
		t.Errorf("Title must be %s, but got %v", "this is title", post.Title)
	}
	//Summary
	if post.Summary != "this summary" {
		t.Errorf("Title must be %s, but got %v", "this summary", post.Summary)
	}
	//slug
	if post.Slug != "hello-boy-google" {
		t.Errorf("Title must be %s, but got %v", "hello-boy-google", post.Slug)
	}
	//tag
	if len(post.Tags) != 3 {
		t.Errorf("Title must be %d, but got %d", 3, len(post.Tags))
	}
}
