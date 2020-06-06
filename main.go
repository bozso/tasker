package main

import (
    "os"
    "os/exec"

    "github.com/bozso/gotoolbox/path"
    "github.com/bozso/gotoolbox/cli"
)

const taskCmd = "task"

var (
    parentDir = path.New("..")
)

func Main() (err error) {
    c := cli.New("tasker",
        "Search for Taskfile.yml in the current and parent directories and execute it.")
    
    err = c.Parse(os.Args[1:])
    if err != nil {
        return
    }
    
    taskFile := path.New("Taskfile.yml")
    
    for {
        b, err := taskFile.Exist()
        
        if err != nil {
            return err
        }
        
        if b {
            break
        } else {
            taskFile = parentDir.Join(taskFile.String())
        }
    }
    
    cmd := exec.Command(taskCmd,
        append(os.Args[1:], "--taskfile", taskFile.String())...)
    
    out, err := cmd.CombinedOutput()
    println(string(out))
    return
}

func main() {
    cli.Run(Main)
}
