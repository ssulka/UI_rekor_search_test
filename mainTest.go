package main

import (
	//"fmt"

	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	chromePath := "/usr/bin/google-chrome"
	log.Println("Launching browser...")
	browser := rod.New().ControlURL(launcher.New().Bin(chromePath).Headless(true).MustLaunch()).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("http://localhost:3000")
	// waits for the page to load
	page.MustWaitLoad()
	log.Println("waiting..")
	// if needed, this will scroll to the button (added in case some changes will be added to the UI)
	page.MustElement("#rekor-search-attribute").MustScrollIntoView()
	log.Println("scrolling...")
	// will click on the rolldown button
	page.MustElement("#rekor-search-attribute").MustClick()
	log.Println("click1..")
	// waits for the rolldown to load
	page.MustElement("#rekor-search-attribute").MustWaitVisible()
	// will choose email option from the rolldown
	page.MustElement("#email").MustClick()

	log.Println("Option has been selected")

	//page.MustElement("#")
	//fill textfield with chosen input
	page.MustElement("#rekor-search-email").MustInput("jdoe@redhat.com").MustType(input.Enter)

}
