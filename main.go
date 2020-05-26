package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/irnes/go-mailer"
	"github.com/spf13/viper"
)

// config stores the configuration values
type config struct {
	Address   []string
	Port      []string
	Recipient map[string][]string
}

// Pinger pings the host and sends email if host is not avaliable
func pinger(configuration *config) {
	for {
		for i := 0; i < len(configuration.Address); i++ {

			//Ping syscall, -c ping count, -i interval, -w timeout
			out, _ := exec.Command("ping", configuration.Address[i], "-c 5", "-i 3", "-w 10").Output()
			if strings.Contains(string(out), "Destination Host Unreachable") {
				fmt.Println("Server down")
				var (
					host     = "xxx"
					user     = "xxx"
					pass     = "xxx"
					recipent = "xxx"
				)

				config := mailer.Config{
					Host: host,
					Port: 465,
					User: user,
					Pass: pass,
				}

				Mailer := mailer.NewMailer(config, true)

				mail := mailer.NewMail()
				mail.FromName = "Go Mailer"
				mail.From = user
				mail.SetTo(recipent)

				mail.Subject = "Server "
				mail.Body = "Your server is down"

				if err := Mailer.Send(mail); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Server is running")

			}
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Set the path to look for the configurations file
	viper.AddConfigPath("/home/alem/go/src/github.com/aTTiny73/microservice")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var b config
	err := viper.Unmarshal(&b)
	if err != nil {
		fmt.Println("unable to set struct host", err)
	}
	pinger(&b)
	fmt.Println(b.Address)
	fmt.Println(b.Port)
	fmt.Println(b.Recipient)
}