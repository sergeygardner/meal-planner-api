package main

import (
	"bufio"
	"flag"
	"fmt"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/sergeygardner/meal-planner-api/ui/cmd/cli/handler"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type commands []string

func (i *commands) String() string {
	return strings.Join(*i, "")
}

func (i *commands) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var flagCommands commands

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
	flag.Var(&flagCommands, "command", "Generate router documentation")
}

func main() {
	flag.Parse()
	_ = make(chan string)

	fmt.Println("Welcome to cli for Meal Planner.")

	_, _ = handler.Help("")

	scanner := bufio.NewScanner(os.Stdin)

	for _, command := range flagCommands {
		statusRunCommand, errorRunCommand := handler.Run(command)

		if statusRunCommand == handler.StatusExit {
			return
		} else if errorRunCommand != nil {
			log.Error(errorRunCommand)
		}
	}

	for {
		scanner.Scan()

		statusRunCommand, errorRunCommand := handler.Run(scanner.Text())

		if statusRunCommand == handler.StatusExit {
			return
		} else if errorRunCommand != nil {
			log.Error(errorRunCommand)
		}
	}
}
