package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type car struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	AdUrl    string `json:"urlad"`
	PhotoUrl string `json:"urlphoto"`
}

func main() {
	c := colly.NewCollector()

	findNextPageOrEnd()
	var cars []car
	// Find and visit all links
	c.OnHTML("div.search-row.swipable", func(e *colly.HTMLElement) {

		car := car{
			Name:     e.ChildText("h2.title"),
			Price:    e.ChildText("span.price-no-decimals"),
			AdUrl:    "https://car.gr" + e.ChildAttr("a", "href"),
			PhotoUrl: e.ChildAttr("img.thumbnail__image.swipable.tw-max-w-full.tw-h-auto", "src"),
		}

		cars = append(cars, car)
		fmt.Println(cars)
		content, err := json.Marshal(cars)
		if err != nil {
			log.Println(err.Error())
		}
		os.WriteFile("cars.json", content, 0755)
		fmt.Println(len(cars))
	})

	c.OnHTML("[role=menuitem]", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		// fmt.Println("goes to: " + e.Attr("href"))
		if strings.Compare(e.Text, "»") > -1 {
			fmt.Println(e.Text)
			// fmt.Print(e.Attr("disabled"))
			c.Visit(nextPage)
		}

	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://www.car.gr/classifieds/cars/?condition=new&condition=used&distance=1000&engine_power-from=90&features=302&mileage-to=%3C125000&offer_type=sale&onlyprice=1&pg=1&postcode=56224&prefecture=7&price-to=%3C17500&price=%3E50&registration-from=%3E2012&rg=2&significant_damage=f&sort=pra/")
}

func findNextPageOrEnd() {
	c := colly.NewCollector()
	// var disabledNextPage string
	c.OnHTML("div.tw-flex.tw-w-full", func(e *colly.HTMLElement) {
		fmt.Println(e)
		// disabledNextPage = e.ChildText("h5")
		// // if disabledNextPage {
		// // }
		// // disabledNextPage = e.Text
		// fmt.Println(disabledNextPage)
		// stringPouPsaxnw := " Δεν βρέθηκαν αποτελέσματα για αυτά τα κριτήρια "
		// apotelesma := strings.Compare(disabledNextPage, stringPouPsaxnw)
		// if apotelesma > -1 {
		os.Exit(3)
		// }
	})

	// return disabledNextPage
}
