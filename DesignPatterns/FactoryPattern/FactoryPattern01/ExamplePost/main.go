package main

import "fmt"

type Post struct{
	Content, Website, Author string
}

type PostFactory struct{
	Website, Author string
}

func NewPostFactory(website string, author string) func(content string) *Post{
	return func(content string) *Post{
		return &Post{content, website, author}
	}
}

func main(){
	devPostFactory := NewPostFactory("DEV.to", "Tomassirio, what does that mean?")
	mediumPostFactory := NewPostFactory("Medium.com", "Tomassirio, what does that mean?")
	hackerNewsPostFactory := NewPostFactory("Hackernews.com", "Tomassirio, what does that mean?")

	content := "This is what is posted today"

	devPost := devPostFactory(content)
	mediumPost := mediumPostFactory(content)
	hackerNewsPost := hackerNewsPostFactory(content)

	fmt.Println(devPost)
	fmt.Println(mediumPost)
	fmt.Println(hackerNewsPost)
}