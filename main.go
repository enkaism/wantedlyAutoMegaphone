package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"time"
	"os"
	"strconv"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	// login facebook
	if err := page.Navigate("https://www.facebook.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	email := page.FindByID("email")
	if err := email.SendKeys(os.Getenv("FACEBOOK_EMAIL")); err != nil {
		fmt.Println(err)
	}

	password := page.FindByID("pass")
	if err := password.SendKeys(os.Getenv("FACEBOOK_PASSWORD")); err != nil {
		fmt.Println(err)
	}

	loginButton := page.Find("#loginbutton input[type=\"submit\"]")
	if err := loginButton.Click(); err != nil {
		fmt.Println(err)
	}
	sleep()

	// login wantedly
	if err := page.Navigate("https://www.wantedly.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	sleep()

	// cheer
	for i,arg := range os.Args {
		if i == 0 {
			continue
		}
		if err := page.Navigate(arg); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}

		if err := page.First("div.project-supporters-list-wrapper>div.project-support-link").Click(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)

		if err := page.Find(".sns-checkbox input:checked[type=\"checkbox\"]").Uncheck(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)

		if err := page.Find("#support-with-comment-form button[type=\"submit\"]").Click(); err != nil {
			fmt.Println(err)
		}
		sleep()
	}
}

func sleep() {
	var sleepTime time.Duration
	timeout := os.Getenv("TIMEOUT_SECONDS")
	if timeout == "" {
		sleepTime = 3
	} else {
		t, err := strconv.Atoi(timeout)
		if err != nil {
			log.Fatalf("Failed Atoi:%v", err)
		}
		sleepTime = time.Duration(t)
	}
	time.Sleep(sleepTime * time.Second)
}