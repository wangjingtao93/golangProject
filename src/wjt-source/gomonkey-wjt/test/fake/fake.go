package fake

import (
    "fmt"
    "strings"
    "errors"
    "os/exec"
)

var (
    ErrActual = errors.New("actual")
    ERR_ELEM_EXIST = errors.New("elem already exist")
    ERR_ELEM_NT_EXIST = errors.New("elem not exist")
)

func Exec(cmd string, args ...string) (string, error) {
    cmdPath, err := exec.LookPath(cmd)
    if err != nil {
        fmt.Errorf("exec.LookPath err: %v, cmd: %s", err, cmd)
        return "", errors.New("any")
    }

    var output []byte
    output, err = exec.Command(cmdPath, args...).CombinedOutput()
    if err != nil {
        fmt.Errorf("exec.Command.CombinedOutput err: %v, cmd: %s", err, cmd)
        return "", errors.New("any")
    }
    fmt.Println("CMD[", cmdPath, "]ARGS[", args, "]OUT[", string(output), "]")
    return string(output), nil
}

//wjt add
func ExecWjt(cmd string, args ...string) (string, error) {
    output, err := Exec(cmd, args...)
    return output, err
}

func Belong(points string, lines []string) bool {
    flag := false
    for _, line := range lines {
        flag = true
        for _, r := range points {
            if !strings.ContainsRune(line, r) {
                flag = false
                break
            }
        }
        if flag {
            return true
        }
    }
    return false
}


type Slice []int

func NewSlice() Slice {
    return make(Slice, 0)
}

func (this* Slice) Add(elem int) error {
    for _, v := range *this {
        if v == elem {
            fmt.Printf("Slice: Add elem: %v already exist\n", elem)
            return ERR_ELEM_EXIST
        }
    }
    *this = append(*this, elem)
    fmt.Printf("Slice: Add elem: %v succ\n", elem)
    return nil
}

func (this* Slice) Remove(elem int) error {
    found := false
    for i, v := range *this {
        if v == elem {
            if i == len(*this) - 1 {
                *this = (*this)[:i]

            } else {
                *this = append((*this)[:i], (*this)[i+1:]...)
            }
            found = true
            break
        }
    }
    if !found {
        fmt.Printf("Slice: Remove elem: %v not exist\n", elem)
        return ERR_ELEM_NT_EXIST
    }
    fmt.Printf("Slice: Remove elem: %v succ\n", elem)
    return nil
}

func ReadLeaf(url string) (string, error) {
    output := fmt.Sprintf("%s, %s!", "Hello", "World")
    return output, nil
}

type Etcd struct {

}

func (this *Etcd) Retrieve(url string) (string, error) {
    output := fmt.Sprintf("%s, %s!", "Hello", "World")
    return output, nil
}

var Marshal = func(v interface{}) ([]byte, error) {
    return nil, nil
}


//wjt add
func AddOne(t int32) int32 {
	return t+1
}

func MinusOne(t int32) int32 {
	return t-1
}

func MultiAddOne(t int32) int32 {
	t = MinusOne(t)
	t = AddOne(t)
	t = AddOne(t)
	return t
}


