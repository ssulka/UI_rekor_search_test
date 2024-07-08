package main

import (
	"fmt"
	"testing"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/ysmood/got"
)

type G struct {
	got.G

	browser *rod.Browser
}

var setup = func() func(t *testing.T) G {
	return func(t *testing.T) G {
		launch := launcher.New().Headless(false)
		url := launch.MustLaunch()
		browser := rod.New().ControlURL(url).MustConnect() // create a new browser instance for each test
		t.Parallel()                                       // run each test concurrently
		return G{got.New(t), browser}
	}
}

// a helper function to create an incognito page.
func (g G) page(url string) *rod.Page {
	page := g.browser.MustPage(url)
	g.Cleanup(func() {
		if err := page.Close(); err != nil {
			g.Logf("Failed to close page: %v", err)
		}
	})
	return page
}

const appURL = "http://localhost:3000" //will be replaced by URL from ocp

// test for email
func TestEmail(t *testing.T) {
	fmt.Println("came here0")
	g := setup()(t) // invoke setup to get a new instance of G for this test
	fmt.Println("came here1")
	p := g.page(appURL)
	fmt.Println("came here2")
	//ensure the element is ready before interacting with it
	attrElement := p.MustElement("#rekor-search-attribute")
	attrElement.MustWaitVisible().MustClick()

	fmt.Println("came here3")
	// Use MustSelect to interact with <select> dropdowns
	attrElement.MustElementR("option", "Email")
	attrElement.MustClick()

	//select the "email" option from the dropdown

	// fill the text field with the email "jdoe@redhat.com"
	emailInput := p.MustElement("#rekor-search-email")
	emailInput.MustWaitVisible().MustInput("jdoe@redhat.com")
	fmt.Println("came here5")

	// verify the input value
	inputValue := emailInput.MustProperty("value").String()
	g.Eq(inputValue, "jdoe@redhat.com")
	fmt.Println("came here6")

	searchButton := p.MustElement("#search-form-button")
	searchButton.MustClick()

}
