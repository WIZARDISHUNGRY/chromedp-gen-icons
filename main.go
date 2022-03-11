// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
		chromedp.WithBrowserOption(),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx,
		elementScreenshot(`http://127.0.0.1:4000/icon`, `#icon`, &buf),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
		// chromedp.CaptureScreenshot(res),
	}
}
