package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThrynSec/aes_implementation_poc/internal/app"
	"github.com/ThrynSec/aes_implementation_poc/internal/client"
)

func main() {
	//Check the signal for daemon usage
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		s := <-sigs
		if s == syscall.SIGURG {
			fmt.Println("received sigur")
		} else {
			log.Printf("RECEIVED SIGNAL: %s\n", s)
			os.Exit(1)
		}
	}()

	//flags to be used
	//Example : toolname -f to flag1
	//By default, toolanme -h will print list of flags with description
	//the "" part is the default value. You can put whatever you like.

	flagAPI := flag.String("api", "", "stop or start will stop or start the API backend \n"+
		"endpoints : https://localhost:8080/dns/unsecure/<ip> | https://localhost:8080/dns/secure/<ip> |"+
		" https://localhost:8080/decrypt/<crypted message> | https://localhost:8080/ping/")
	flagClient := flag.String("client", "", "Use client functionnalities : call, encrypt, ")
	flagStr := flag.String("c", "", "content to use for client, under quotes")

	flag.Parse()

	// OPTIONAL - No flag catcher
	if (*flagAPI == "") && (*flagClient == "") {
		fmt.Printf("Error, need at least an argument")
		os.Exit(0)
	} else if *flagAPI == "start" {
		bootAPI()
	} else {
		switch *flagClient {
		case "unsecure":
			client.SendUnsecure(*flagStr)
			os.Exit(1)

		case "secure":
			client.SendSecure(*flagStr)
			os.Exit(1)

		case "message":
			client.SendMessage(*flagStr)
			os.Exit(1)

		case "key":
			client.CreateKey(*flagStr)
			os.Exit(1)

		default:
			fmt.Println("Wrong use")
		}
	}

}

func bootAPI() {
	app.BootApp()
}
