package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func main() {
	url := "https://ln" + "cn.org"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// ele := ".copy" + "-all"

	time.Sleep(time.Second * 10)

	js := `
	new Promise((resolve) => {
		resolve(document.querySelector("html").outerHTML);
	  });
	  
	`

	var str string

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// chromedp.WaitVisible(ele, chromedp.ByQuery),
		// chromedp.Sleep(time.Second*40),
		chromedp.Evaluate(js, &str, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
			return p.WithAwaitPromise(true)
		}),
	); err != nil {
		fmt.Println("err.Error(fdfaaff )")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(str)

}
