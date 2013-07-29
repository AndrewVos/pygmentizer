package pygmentizer

import (
  "runtime"
  "path"
  "io"
  "io/ioutil"
  "os"
  "os/exec"
  "fmt"
  "errors"
  "archive/zip"
  "log"
)

func Highlight(language string, code string) (string, error) {
  if _, err := os.Stat(path.Join(tempPath(), "pygments", "pygmentizer.py")); err != nil {
    installPygments()
  }

  cmd:= exec.Command("/usr/bin/env", "python", path.Join(tempPath(), "pygments", "pygmentizer.py"), "-l", language, "-f", "html")
  writer, _ := cmd.StdinPipe()

  io.WriteString(writer, code)
  writer.Close()
  output,err := cmd.CombinedOutput()

  if err != nil {
    fmt.Printf(string(output))
    return code, errors.New(string(output))
  }
  return string(output), nil
}

func tempPath() string {
  _, filename, _, _ := runtime.Caller(1)
  return path.Join(path.Dir(filename), "tmp")
}

func installPygments() {
  fmt.Printf("[pygmentizer] Installing pygments\n")
  os.MkdirAll(tempPath(), 0700)

  zipPath := path.Join(tempPath(), "pygments.zip")
  ioutil.WriteFile(zipPath, PygmentsZip(), 0700)

  r, err := zip.OpenReader(zipPath)
  if err != nil {
    log.Fatal(err)
  }
  defer r.Close()

  for _, file := range r.File {
    if file.FileInfo().IsDir() == false {
      rc, err := file.Open()
      if err != nil {
        log.Fatal(err)
      }

      allBytes,err := ioutil.ReadAll(rc)
      if err != nil {
        log.Fatal(err)
      }

      filePath := path.Join(tempPath(), file.Name)
      os.MkdirAll(path.Dir(filePath), 0700)

      err = ioutil.WriteFile(filePath, allBytes, 0700)
      if err != nil {
        log.Fatal(err)
      }

      rc.Close()
    }
  }
}
