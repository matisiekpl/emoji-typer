package main

import (
	"github.com/aymerick/raymond"
	"github.com/forPelevin/gomoji"
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Message struct {
	Author  string
	Content string
}

var messages []Message
var box *packr.Box

func getEmojiString(list []gomoji.Emoji) string {
	output := ""
	for _, emoji := range list {
		output = output + emoji.Character
	}
	return output
}

func insert(c echo.Context) error {
	message := Message{
		Author:  getEmojiString(gomoji.FindAll(c.FormValue("author"))),
		Content: getEmojiString(gomoji.FindAll(c.FormValue("content"))),
	}
	if len(message.Author) == 0 || len(message.Content) == 0 {
		return c.HTML(http.StatusOK, "Z pewnego powodu, jedynie emotki są dostępne w tym forum 😎 <a href=\"/\">Cofnij</a>")
	}
	messages = append([]Message{message}, messages...)
	return c.Redirect(http.StatusSeeOther, "/")
}

func index(c echo.Context) error {
	tpl, _ := box.FindString("index.html")
	hostname, _ := os.Hostname()
	html, _ := raymond.Render(tpl, map[string]interface{}{
		"messages": messages,
		"hostname": hostname,
	})
	return c.HTML(http.StatusOK, html)
}

func main() {
	box = packr.New("Templates", "./templates")

	e := echo.New()
	e.GET("/", index)
	e.POST("/insert", insert)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
