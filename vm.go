package main

import (
    "os/exec"
    "time"
    "os"
	"log"
    "fmt"
    "strings"
)

type VM struct {
    ID string
    Created int64
}


func runCommand(command string, args string) string {
    argArray := strings.Split(args, " ")
    cmd := exec.Command(command, argArray...)
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("ERROR RUNNING COMMAND", command)
        print(err)
    }
    return string(output)
}

func shellVM(id string) {
    cmd := exec.Command("docker", "exec", "-it", id, "/bin/bash")
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Println("Error:", err)
    }
}

func initVM() VM {
    db := initDB()
    defer db.Close()
    res := runCommand("docker", "run -d -it ubuntu")
    id := res[:12]
    vm := VM {ID: id, Created: time.Now().Unix()}
    update(db, "VM", []byte(id), SerializeVM(vm))
    return vm
}

func stopVM(id string) {
    runCommand("docker", "stop " + id)
}


func listVM() {
    db := initDB()
    defer db.Close()
    res := retrieveAll(db, "VM")
    for _, v := range res {
        vm := DeSerializeVM(v.Second)
        fmt.Println(string(v.First), vm.ID, vm.Created)
    }
}
