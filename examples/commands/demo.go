package commands

import (
	"errors"
	"fmt"
	"github.com/daforester/getopt-golang/getopt"
	"lwebco.de/cosmic-calendar-go-library/examples/libs/commands/traits"
	"time"
)

type Demo struct {
	getopt.Command
	traits.Configurable
}

func NewDemo() (*Demo, error) {
	a := new(Demo)

	o1, _ := getopt.NewOption('u', "user", getopt.REQUIRED_ARGUMENT, "Specify a username")
	o2, _ := getopt.NewOption('\x00', "end", getopt.OPTIONAL_ARGUMENT, "Specify end")

	_, err := a.Command.Build("demo", a.Handle, o1, o2)

	return a, err
}

func (a *Demo) Handle(opt *getopt.GetOpt) error {
	user := opt.GetOptionString("user")

	if user == "" {
		return errors.New("User not configured for request")
	}

	var end time.Time
	var err error

	if opt.GetOptionString("end") != "" {
		end, err = time.Parse("2006-01-02T15:04:05Z", opt.GetOptionString("end"))

		if err != nil {
			return err
		}
	}

	fmt.Println("Demo Command Run")
	fmt.Print("user: ")
	fmt.Println(user)
	fmt.Print("end:")
	fmt.Println(end)

	return nil
}
