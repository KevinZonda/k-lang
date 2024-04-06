package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/fmtr"
	"github.com/KevinZonda/GoX/pkg/iox"
	"os"
	"strings"
	"testing"
	"time"
)

func TestJupyterEval(t *testing.T) {
	eval := NewReplInterpreter()
	eval.Eval("x = 18")
	vals, err := eval.Eval("x")
	if len(vals) != 1 || err != nil || vals[0] != 18 {
		t.Fatal(vals, err)
	}
}

func TestJupyterServ(t *testing.T) {
	j := `{
  "control_port": 50160,
  "shell_port": 57503,
  "transport": "tcp",
  "signature_scheme": "hmac-sha256",
  "stdin_port": 52597,
  "hb_port": 42540,
  "ip": "127.0.0.1",
  "iopub_port": 40885,
  "key": "a0436f6c-1916-498b-8eb9-e81ab9368e84"
}`
	iox.WriteAllText("jupyter.json", j)
	defer os.Remove("jupyter.json")
	cancel := async.Run(func() {
		JupyterKernel("jupyter.json")
	})
	time.Sleep(1 * time.Second)
	cancel()
}

func TestLsp(t *testing.T) {
	cancel := async.Run(func() {
		Lsp("")
	})
	time.Sleep(1 * time.Second)
	cancel()
}

func TestHelp(t *testing.T) {
	tester.CaptureStdout()
	Help()
	output := tester.StopCaptureStdout()
	tester.Assert(t, strings.TrimSpace(output), strings.TrimSpace(HelpMsg))
}

func TestFmt(t *testing.T) {
	buildconst.TestMod = true
	iox.WriteAllText("main.k", "println('Hello World!')")
	defer os.Remove("main.k")

	fmtr.Fmt("main.k")
	txt, _ := iox.ReadAllText("main.k")
	tester.Assert(t, txt, "println('Hello World!')")

	tester.PanicFx(t, func() { fmtr.Fmt("test.k") }, tester.PanicExFxNotEmptyString)

	iox.WriteAllText("fmt.k", "println()'Hello World!')")
	defer os.Remove("fmt.k")

	tester.PanicFx(t, func() { fmtr.Fmt("fmt.k") }, tester.PanicExFxNotEmptyString)
}
