package main

import (
  "fmt"
  "log"
  "time"
  "bytes"
  "net/http"
  "text/template"
)

type Item struct {
  Title       string
  Description string
  Link        string
  Guid        string
  PubDate     time.Time
}

type RSSFeed struct {
  Title         string
  Description   string
  Link          string
  Copyright     string
  LastBuildDate time.Time
  PubDate       time.Time
  Items         []Item
}

var feed RSSFeed

func (feed *RSSFeed) String() string {
  temp, err := template.ParseFiles("rss_template.rss") 

  if err != nil {
    log.Println("Error while parsing RSS template")
    log.Fatal(err)
  }

  var buffer bytes.Buffer
  err = temp.Execute(&buffer, feed)

  if err != nil {
    log.Println("Error while parsing RSS template")
    log.Println(err)

    return ""
  }

  return buffer.String()
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("RSS feed requested")
  fmt.Fprintf(w, feed.String())
}

func main() {
  feed = RSSFeed{
    Title: "Marek's amazing RSS feed",
    Description: "This is a test RSS feed",
    Link: "http://127.0.0.51",
    Copyright: "2020 (example@email.com) All rights reserved",
    LastBuildDate: time.Now(),
    PubDate: time.Now(),
    Items: []Item{
      Item{
        Title: "First thing", 
        Description: "First things always go first...",
        Link: "http://127.0.0.51/blog/1",
        Guid: "1",
        PubDate: time.Now(),
      },
      Item{
        Title: "Second thing", 
        Description: "After the first thing goes.. you've guessed it.. the second thing!",
        Link: "http://127.0.0.51/blog/2",
        Guid: "2",
        PubDate: time.Now(),
      },
    },
  }
  
  http.HandleFunc("/rss", rssHandler)

  log.Println("Starting a server")
  http.ListenAndServe(":5050", nil)
}
