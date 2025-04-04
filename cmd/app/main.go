package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"email-push-service/pkg/constants"
	"email-push-service/pkg/listners"
	"email-push-service/pkg/logger"

	"email-push-service/internal/server"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Web Content Downloader Pipeline"
	app.Usage = "Triggers pipeline to url's and download it's content"
	app.Version = "latest"

	app.Commands = []*cli.Command{
		{
			Name:      "listen",
			Usage:     "Listens to a topic and listener of your choice",
			UsageText: "trigger -l listner -t topic",
			Action:    RunServerAndListen,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "Listner",
					Aliases:  []string{"l", "LISTNER"},
					Usage:    "Name of listener",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "Topic",
					Aliases:  []string{"t", "TOPIC"},
					Usage:    "Topic to listen to",
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf("%s", err)
		os.Exit(1)
	}
}

// RunPipeline - Function that will run the server
func RunServerAndListen(c *cli.Context) error {
	go RunServer()
	listner := c.String("LISTNER")
	topic := c.String("TOPIC")
	if isValid, ok := constants.MapOfListeners[strings.ToLower(listner)]; !ok || !isValid {
		logger.Errorf("Listnerer not available: %s", listner)
		return errors.New("Listnerer not available")
	}

	listnerObject := listners.GetListnerObject(strings.ToLower(listner))
	if listnerObject == nil {
		logger.Errorf("Unable to create %s listner object", listner)
		return errors.New("Unable to create listner object")
	}
	if err := listnerObject.InitializeListner(); err != nil {
		logger.Errorf("Unable to connect to %s listner: %s", listner, err)
		return err
	}
	logger.Info("Listing to topic: ", topic)
	listnerObject.Listen(topic)
	return nil
}

func RunServer() error {
	http.HandleFunc("/", server.HandleHome)
	http.HandleFunc("/login", server.HandleLogin)
	http.HandleFunc("/oauth2callback", server.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	return http.ListenAndServe(":8080", nil)

}
