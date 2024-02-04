package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anaskhan96/soup"
)

const (
	DAY_GAINERS_YAHOO string = "https://finance.yahoo.com/screener/predefined/day_gainers"
	DAY_LOSERS_YAHOO  string = "https://finance.yahoo.com/screener/predefined/day_losers"
)

func ParseYahooListings(url string) ([]YahooListing, []string) {
	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println("Error fetching gainers...")
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	table := doc.Find("tbody").FindAll("tr")
	var listings []YahooListing
	var symbols []string

	for _, lst := range table {
		symbol := lst.Children()[0].Find("a").Text()
		listing := YahooListing{
			Symbol:    symbol,
			Name:      lst.Children()[1].Text(),
			Price:     ParseFloat64(lst.Children()[2].Children()[0].Attrs()["value"]),
			Change:    ParseFloat64(lst.Children()[3].Children()[0].Attrs()["value"]),
			PChange:   ParseFloat64(lst.Children()[4].Children()[0].Attrs()["value"]),
			Volume:    ParseInt64(lst.Children()[5].Children()[0].Attrs()["value"]),
			MarketCap: ParseInt64(lst.Children()[7].Children()[0].Attrs()["value"]),
		}

		symbols = append(symbols, symbol)
		listings = append(listings, listing)
	}

	return listings, symbols
}

func ParseFloat64(source string) float64 {
	target, _ := strconv.ParseFloat(source, 64)
	return target
}

func ParseInt64(source string) int64 {
	target, _ := strconv.ParseInt(source, 10, 64)
	return target
}

type YahooListing struct {
	Symbol, Name    string
	Price           float64
	Change, PChange float64
	Volume          int64
	MarketCap       int64
}
