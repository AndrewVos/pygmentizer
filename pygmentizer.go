package pygmentizer

import (
  "runtime"
  "path"
  "io"
  "io/ioutil"
  "os/exec"
  "fmt"
)

func Highlight(language string, code string) string {
  cmd:= exec.Command("/usr/bin/env", "python", pygmentizerPath(), "-l", language, "-f", "html")
  writer, _ := cmd.StdinPipe()
  errors, _ := cmd.StderrPipe()
  allErrors,_ := ioutil.ReadAll(errors)
  fmt.Printf("pygmentizer errors:\n" + string(allErrors) + "\n")
  io.WriteString(writer, code)
  writer.Close()
  output,_ := cmd.Output()
  return string(output)
}

func pygmentizerPath() string {
  _, filename, _, _ := runtime.Caller(1)
  return path.Join(path.Dir(filename), "pygmentizer.py")
}
