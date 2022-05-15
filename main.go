package main

import (
	"context"
	_ "embed"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/mattn/entgo-bbs/ent"
	"github.com/mattn/entgo-bbs/ent/entry"

	"github.com/mattn/go-slim"
	_ "github.com/mattn/go-sqlite3"
)

type Payload struct {
	Content string `json:"content"`
}

//go:embed template/index.slim
var tmpl string

func main() {
	var dsn string

	flag.StringVar(&dsn, "dsn", "file:bbs.sqlite?cache=shared&_fk=1", "connection string")
	flag.Parse()

	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	t, err := slim.Parse(strings.NewReader(tmpl))
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		eq := client.Entry.Query().Order(ent.Desc(entry.FieldCreatedAt)).Limit(10)
		entries := eq.AllX(context.Background())
		c.Request().Header.Set("content-type", "text/html")
		return t.Execute(c.Response(), map[string]interface{}{
			"entries": entries,
		})
	})
	e.POST("/add", func(c echo.Context) error {
		e := client.Entry.Create()
		e.SetContent(c.FormValue("content"))
		if _, err := e.Save(context.Background()); err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
