package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "add", "Fill ITR").CombinedOutput()

	if err != nil {
		fmt.Printf("something went wrong : %s", err)
	}
	//fmt.Printf("The task output is %s\n", out)
	str := fmt.Sprintf("%s", out)
	fmt.Println(str)

	flag := strings.Contains(str, "Fill ITR")
	assert.Equal(t, flag, true)
}

func TestAdd1(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "add").CombinedOutput()

	if err != nil {
		fmt.Printf("Length of arguments is not sufficient : %s", err)
	}
	str := fmt.Sprintf("%s", out)
	fmt.Println(str)
}

func TestList(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "list").CombinedOutput()

	if err != nil {
		fmt.Printf("something went wrong : %s", err)
	}
	//fmt.Printf("The task output is %s\n", out)
	str := fmt.Sprintf("%s", out)
	fmt.Println(str)
}

func TestList1(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "list1").CombinedOutput()

	if err != nil {
		fmt.Printf("Could not list the task : %s", err)
	}
	//fmt.Printf("The task output is %s\n", out)
	str := fmt.Sprintf("%s", out)
	fmt.Println(str)
}

func TestDone(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "done", "Fill ITR").CombinedOutput()

	if err != nil {
		fmt.Printf("Could not list the task : %s", err)
	}
	//fmt.Printf("The task output is %s\n", out)
	str := fmt.Sprintf("%s", out)
	fmt.Println(str)
}
