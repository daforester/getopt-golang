package main

import (
	"github.com/daforester/getopt-golang/examples/commands"
	"github.com/daforester/getopt-golang/getopt"
	"github.com/daforester/getopt-golang/getopt/errors"
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
)

const NAME = "run"
const VERSION = "1.0-beta"

func main() {
	// Create new argument processor
	opt, err := getopt.NewGetOpt("", nil)

	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	// Setup global command options
	addOptions(opt)

	// Setup commands
	addCommands(opt)

	// Process command line arguments
	err = opt.Process()

	if err != nil {
		logrus.Errorln(err)
		if err.(errors.GetOptError).Type() == errors.ERROR_UNEXPECTED_ARGUMENT {
			fmt.Println("\n" + opt.GetHelpText(nil))
		}
		os.Exit(1)
	}

	// Version requested so output & quit
	optValue := opt.GetOptionValue("version")
	if optValue != nil && len(optValue) > 0 {
		fmt.Println(fmt.Sprintf("%s: %s", NAME, VERSION))
		os.Exit(0)
	}

	// Acquire selected command
	command := opt.GetCommand("")

	// Generate help
	optHelp := opt.GetOptionValue("help")

	// No command specified, output help
	if command == nil || (optHelp != nil && len(optHelp) > 0) {
		fmt.Println(opt.GetHelpText(nil))
		os.Exit(0)
	}

	// Run Handle function for specified command
	handlerFunc := command.GetHandler()
	if handlerFunc != nil {
		err := handlerFunc(opt)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("No handlerFunc found")
	}
}

func addCommands(opt *getopt.GetOpt) {
	c1, _ := getopt.NewCommand("test-setup", func(o *getopt.GetOpt) error {
		fmt.Println("When you see this message the setup works." + "\n")
		return nil
	})
	c1.SetDescription("Check if setup works")

	c2, err := commands.NewDemo()
	c2.SetDescription("Example of command with parameters")

	if err != nil {
		logrus.Error(err)
	}

	_, _ = opt.AddCommands(c1, c2)
}

func addOptions(opt *getopt.GetOpt) {
	o1, _ := getopt.NewOption('\x00', "version", getopt.NO_ARGUMENT)
	o1.SetDescription("Show version information and quit")

	o2, _ := getopt.NewOption('?', "help", getopt.NO_ARGUMENT)
	o2.SetDescription("Show this help and quit")

	o3, _ := getopt.NewOption('f', "config", getopt.OPTIONAL_ARGUMENT)
	o3.SetValidation(getopt.FileIsReadable, nil)
	o3.SetDescription("Specify configuration file to use")

	_, _ = opt.AddOptions(o1, o2, o3)
}
