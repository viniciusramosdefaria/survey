package main

import (
	"bytes"
	"fmt"
	"github.com/Netflix/go-expect"
	"os/exec"
	"regexp"
	"time"
)

type testCase struct {
	input     string
	expectOut string
}

func main() {

	terminalRunner()

}

func terminalRunner() {
	bufOut := new(bytes.Buffer)
	c, err := expect.NewConsole(expect.WithStdout(bufOut))
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	cmd := exec.Command("rit", "scaffold", "generate", "coffee-go")

	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()

	go func() {
		c.ExpectEOF()
	}()

	testCases := []testCase{
		{input: "test", expectOut: ".*Type your name.*"},
		{input: "es", expectOut: ".*Pick your coffee.*"},
		{input: "", expectOut: ".*Delivery.*"},
	}

	go func() {
		for i := 0; i < len(testCases); i++ {
			for {
				matched, _ := regexp.MatchString(testCases[i].expectOut, bufOut.String())
				if matched {
					c.SendLine(testCases[i].input)
					break
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	cmd.Start()

	cmd.Wait()

	fmt.Println(bufOut.String())
}
