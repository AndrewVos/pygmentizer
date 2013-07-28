package pygmentizer

import (
  "runtime"
  "path"
  "io"
  "os/exec"
  "fmt"
)

func Highlight(language string, code string) string {
  cmd:= exec.Command("/usr/bin/env", "python", pygmentizerPath(), "-l", language, "-f", "html")
  writer, _ := cmd.StdinPipe()

  io.WriteString(writer, code)
  writer.Close()
  output,err := cmd.CombinedOutput()

  if err != nil {
    fmt.Printf(string(output))
    return code
  }
  return string(output)
}

func pygmentizerPath() string {
  _, filename, _, _ := runtime.Caller(1)
  return path.Join(path.Dir(filename), "pygmentizer.py")
}
