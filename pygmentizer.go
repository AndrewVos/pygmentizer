package pygmentizer

import (
  "io"
  "os/exec"
  "fmt"
  "errors"
)

func Highlight(language string, code string) (string, error) {
  cmd:= exec.Command("/usr/bin/env", "pygmentize", "-l", language, "-f", "html")
  writer, _ := cmd.StdinPipe()

  io.WriteString(writer, code)
  writer.Close()
  output,err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println()
    fmt.Printf("pygmentizer:")
    fmt.Printf(string(output))
    fmt.Println()
    return code, errors.New(string(output))
  }
  return string(output), nil
}
