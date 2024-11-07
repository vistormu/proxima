package evaluator

import (
    "io"
    "os/exec"
    "sync"
    "bufio"
    "strings"
)

type Interpreter struct {
    cmd    *exec.Cmd
    stdin  io.WriteCloser
    stdout io.ReadCloser
    stderr io.ReadCloser
    mutex  sync.Mutex
}

func NewInterpreter(command string, args ...string) (*Interpreter, error) {
    cmd := exec.Command(command, args...)
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return nil, err
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return nil, err
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        return nil, err
    }
    if err := cmd.Start(); err != nil {
        return nil, err
    }
    return &Interpreter{
        cmd:    cmd,
        stdin:  stdin,
        stdout: stdout,
        stderr: stderr,
    }, nil
}

func (interp *Interpreter) Execute(script string) (string, string, error) {
    interp.mutex.Lock()
    defer interp.mutex.Unlock()

    // Write the script to stdin
    _, err := io.WriteString(interp.stdin, script+"\n")
    if err != nil {
        return "", "", err
    }

    // Read the output
    output, err := bufio.NewReader(interp.stdout).ReadString('\n')
    if err != nil {
        return "", "", err
    }

    // Read any errors
    errOutput, _ := io.ReadAll(interp.stderr)

    return strings.TrimSpace(output), strings.TrimSpace(string(errOutput)), nil
}
