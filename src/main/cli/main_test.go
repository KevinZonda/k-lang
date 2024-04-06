package main_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	main "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli/funcs"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/rhysd/go-fakeio"
	"os"
	"testing"
	"time"
)

func TestMainM(t *testing.T) {
	buildconst.TestMod = true
	os.Args = []string{"cli", "wrong", "dew"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString)
}

func TestMainREPL(t *testing.T) {
	buildconst.TestMod = true
	os.Args = []string{"cli", "repl"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString) // cmd should panic due to readline has no input
	os.Args = []string{"cli"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString) // cmd should panic due to readline has no input

}

func TestMainFmt(t *testing.T) {
	buildconst.TestMod = true
	iox.WriteAllText("main.k", "println('Hello World!')")
	defer os.Remove("main.k")

	os.Args = []string{"cli", "fmt", "main.k"}
	main.M()
	txt, _ := iox.ReadAllText("main.k")
	tester.Assert(t, txt, "println('Hello World!')")

	os.Args = []string{"cli", "fmt"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString)

	os.Args = []string{"cli", "fmt", "test.k"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString)

	iox.WriteAllText("fmt.k", "println()'Hello World!')")
	defer os.Remove("fmt.k")
	os.Args = []string{"cli", "fmt", "fmt.k"}
	tester.PanicFx(t, main.M, tester.PanicExFxNotEmptyString)
}

func TestMainRun(t *testing.T) {
	iox.WriteAllText("main.k", "println('Hello World!')")
	defer os.Remove("main.k")

	os.Args = []string{"cli", "main.k"}
	tester.CaptureStdout()
	main.M()
	output := tester.StopCaptureStdout()
	fmt.Print(output)
	tester.Assert(t, output, "Hello World!\n")

	os.Args = []string{"cli", "run", "main.k"}
	tester.CaptureStdout()
	main.M()
	output = tester.StopCaptureStdout()
	fmt.Print(output)
	tester.Assert(t, output, "Hello World!\n")

}

func TestMainVersion(t *testing.T) {
	os.Args = []string{"cli", "version"}
	tester.CaptureStdout()
	main.M()
	output := tester.StopCaptureStdout()
	fmt.Print(output)
	tester.Assert(t, output, buildconst.Msg()+"\n")
}

func TestMainHelp(t *testing.T) {
	os.Args = []string{"cli", "help"}
	tester.CaptureStdout()
	main.M()
	output := tester.StopCaptureStdout()
	fmt.Print(output)
	tester.Assert(t, output, funcs.HelpMsg)
}

func TestREPL(t *testing.T) {
	buildconst.TestMod = true
	fake := fakeio.Stdin("println('Hello World!')\n:exit\n")
	defer fake.Restore()

	os.Args = []string{"cli", "repl"}
	tester.PanicFx(t, main.M, func(exp string) bool {
		return exp == "0"
	})
}

func TestLSP(t *testing.T) {
	buildconst.TestMod = true
	os.Args = []string{"cli", "lsp"}
	cancel := async.Run(main.M)
	time.Sleep(1 * time.Second)
	cancel()

}
