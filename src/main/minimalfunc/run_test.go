package minimalfunc_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/minimalfunc"
	"os"
	"testing"
)

const MINIFUNC_RUN_TEST_FILE = "main.k"
const MINIFUNC_RUN_TEST_MOD = "k.mod"

func initRunEnv() {
	minimalfunc.Mod("new", []string{"console"})
	minimalfunc.Mod("restore", nil)
}

func destroyRunEnv() {
	os.Remove(MINIFUNC_RUN_TEST_FILE)
	os.Remove(MINIFUNC_RUN_TEST_MOD)
	os.RemoveAll("vendor")
}

func TestRun(t *testing.T) {
	initRunEnv()
	defer destroyRunEnv()
	minimalfunc.Run(MINIFUNC_RUN_TEST_FILE)
}

func TestRunMod(t *testing.T) {
	initRunEnv()
	defer destroyRunEnv()
	minimalfunc.Run(".")
}

func TestMod(t *testing.T) {
	minimalfunc.Mod("new", []string{"lib"})
	destroyRunEnv()
}

func TestModPanic(t *testing.T) {
	tester.PanicFx(t, func() {
		minimalfunc.Mod("new", []string{})
	}, tester.PanicExFxNotEmptyString)
	tester.PanicFx(t, func() {
		minimalfunc.Mod("new", []string{"114514"})
	}, tester.PanicExFxNotEmptyString)
	tester.PanicFx(t, func() {
		minimalfunc.Mod("hi", []string{"114514"})
	}, tester.PanicExFxNotEmptyString)
}
