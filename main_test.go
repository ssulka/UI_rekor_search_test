// Package main ...
package main

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/ysmood/got"
)

type G struct {
	got.G

	browser *rod.Browser
}

var setup = func() func(t *testing.T) G {
	browser := rod.New().MustConnect()

	return func(t *testing.T) G {
		t.Parallel() // run each test concurrently

		return G{got.New(t), browser}
	}
}()

// a helper function to create an incognito page.
func (g G) page(url string) *rod.Page {
	page := g.browser.MustIncognito().MustPage(url)
	g.Cleanup(page.MustClose)
	return page
}

const appURL = "http://localhost:3000"

// test for email
func TestEmail(t *testing.T) {
	g := setup(t)

	p := g.page(appURL)

	p.MustElement("#rekor-search-attribute").MustClick()

	// Select the "email" option from the dropdown
	p.MustElementR("option", "email").MustClick()

	// Fill the text field with the email "jdoe@redhat.com"
	p.MustElement("#rekor-search-email").MustInput("jdoe@redhat.com")

	// Verify the input value
	g.Eq(p.MustElement("##rekor-search-email").MustProperty("value").String(), "jdoe@redhat.com")

}
