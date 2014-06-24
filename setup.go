package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"os"
	"path/filepath"
	"net/url"
	"code.google.com/p/go-netrc/netrc"
	"github.com/cyberdelia/heroku-go/v3"
)

var (
	password = flag.String("apikey", "", "api key")
	repo     = flag.String("archive", "", "archive url")

	apiURL    = "https://api.heroku.com"
	netrcPath = filepath.Join(os.Getenv("HOME"), ".netrc")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

  if *password == "" {
  	u, _ := url.Parse(apiURL)
  	_,netrcpass := getCreds(u)
  	heroku.DefaultTransport.Password = netrcpass
  } else {
  	heroku.DefaultTransport.Password = *password
  }


	h := heroku.NewService(heroku.DefaultClient)

	setup, err := h.AppSetupCreate(heroku.AppSetupCreateOpts{
		SourceBlob: &struct {
			URL     *string `json:"url,omitempty"`
		}{
			URL: heroku.String(*repo),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	if setup.Status == "pending"{
		setup, err = h.AppSetupInfo(setup.ID)
		fmt.Print("--> Created app "+setup.App.Name)
		fmt.Print("\n----> App ID:"+setup.App.ID)
		fmt.Print("\n----> Setting up config vars and add-ons")
		for setup.Status == "pending" {
			setup, err = h.AppSetupInfo(setup.ID)

			if setup.Build.ID != "null" && setup.Build.ID != "" {
				if(setup.Build.Status == "pending"){
					fmt.Print(".Done.\n")
					fmt.Print("\n--> Build "+setup.Build.ID+" pending.")
					for setup.Build.Status == "pending" {
						fmt.Print(".")
						time.Sleep(time.Second)
						setup, err = h.AppSetupInfo(setup.ID)
					}
					fmt.Print("\n----> Build "+setup.Build.Status+"\n")
				}

			}

			fmt.Print(".")
			time.Sleep(time.Second)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if setup.Status == "failed" {
		fmt.Print("\n--> Deleting app...Setup failed: ")
		fmt.Print(*setup.FailureMessage)
	}

	if setup.Status == "succeeded" {
		fmt.Print("\n--> Postdeploy script completed with exit code ")
		fmt.Print(setup.Postdeploy.ExitCode)
		if(setup.Postdeploy.Output != "null" && setup.Postdeploy.Output != ""){
			fmt.Print(" and output: "+setup.Postdeploy.Output)
		}
		fmt.Print("\n--> App setup complete.")
	}
}

func getCreds(u *url.URL) (user, pass string) {

	m, err := netrc.FindMachine(netrcPath, u.Host)
	if err != nil {
		fmt.Printf("netrc error (%s): %v", u.Host, err)
	}

	return m.Login, m.Password
}
