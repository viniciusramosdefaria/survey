package main

import (
	"bytes"
	"fmt"
	"github.com/Netflix/go-expect"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is your name?",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: survey.Required,
	},
}

func main() {

	//answers := struct {
	//	Name  string
	//	Color string
	//}{}
	//
	//// ask the question
	//err := survey.Ask(simpleQs, &answers)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//// print the answers
	//fmt.Printf("%s chose %s.\n", answers.Name, answers.Color)

	createTerminal()
}

func createTerminal(){
	bufOut := new(bytes.Buffer)
	bufIn := new(bytes.Buffer)
	c, err := expect.NewConsole(expect.WithStdout(bufOut), expect.WithStdin(bufIn))
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	cmd := exec.Command("rit","set", "context")

	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()

	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.ExpectEOF()
	}()

	cmd.Start()

	c.SendLine("ty")

	c.ExpectString("New context:")

	c.SendLine("test\n")

	cmd.Wait()

	fmt.Println(bufOut.String())

}
