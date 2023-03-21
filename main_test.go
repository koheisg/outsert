package main

import (
    "bytes"
    "os/exec"
    "testing"
)

func TestMain(t *testing.T) {
    cmd := exec.Command("./outsert")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        t.Fatalf("cmd.Run() failed with %s", err)
    }
    expected := "hello world\n"
    if out.String() != expected {
        t.Fatalf("expected %q but got %q", expected, out.String())
    }
}
