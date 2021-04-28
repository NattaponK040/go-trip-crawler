package pkg

import (
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"go-crawler/repository"
	"os"
	"strings"
	"time"
)

type TripadvisorProc struct {
	mg  *repository.MgRepository
	c   *colly.Collector
	set map[string]bool
}

const THAILAND = "Thailand"
const HOTELS = "Hotels"
const Restaurants = "Restaurants"
const DOMAIN = "tripadvisor"

func NewTripasdvisor(mg *repository.MgRepository) *TripadvisorProc {
	return &TripadvisorProc{
		mg: mg,
		set: map[string]bool{},
	}
}

func (t *TripadvisorProc) Runprocess() {
	t.c = colly.NewCollector()
	t.c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "tripadvisor",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	t.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if t.set[link] {
			return
		}
		if strings.Contains(link, "tripadvisor") {
			t.set[link] = true
			e.Request.Visit(link)
		} else if strings.Index(link, "/"+HOTELS) == 0 || strings.Index(link, "/"+Restaurants) == 0 {
			link = "https://www.tripadvisor.com" + link
			t.set[link] = true
			e.Request.Visit(link)
		} else {
			return
		}
	})

	t.c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting: ", r.URL)
		if t.set[r.URL.String()] {
			return
		}
		t.set[r.URL.String()] = true
		r.Visit(r.URL.String())
	})
	t.c.OnResponse(func(r *colly.Response) {
		fmt.Println("resp: ")
		if !(strings.Contains(r.Request.URL.String(), "Hotel_Review") || strings.Contains(r.Request.URL.String(), "Restaurant_Review")) {
			fmt.Println("No type: ")
			return
		}
		fmt.Println("parse: ")
		doc, err := xmlquery.Parse(bytes.NewBuffer(r.Body))
		if err != nil {
			return
		}
		fmt.Println("title", doc.SelectElements("h1._1mTlpMC3"))

		//r.ForEach("p", func(_ int, elem *colly.HTMLElement) {
		//	if strings.Contains(elem.Text, "golang") {
		//		fmt.Println(elem.Text)
		//	}
		//})
		//err := ioutil.WriteFile("index.html", r.Body, 0755)
		//log.Error(err)
		os.Exit(0)
	})

	t.c.Visit("https://www.tripadvisor.com/Hotels-g293915-Thailand-Hotels.html")
}
