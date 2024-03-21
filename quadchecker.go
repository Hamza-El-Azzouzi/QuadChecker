package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func counter(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
		} else {
			if flag {
				countX++
			}
		}
	}
	return countX, countY
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	var output []rune
	inputStr := ""

	for reader.Scan() {
		inputStr = inputStr + reader.Text() + "\n"
	}

	//if size == 0 --> not a raid func
	for _, a := range inputStr {
		output = append(output, a)
	}
	x, y := counter(output)
	//determineInput(output)
	strX := printNbr(x)
	strY := printNbr(y)
	// fmt.Print(inputStr)
	// result := ""

	// result, _ = runCommand("./quadC", strX, strY)
	// fmt.Print(result)
	checkQuad(strX, strY, inputStr)

}

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
func checkQuad(strX string, strY string, inputStr string) {

	Quad := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	result := ""
	str := []string{}
	for i := 0; i < len(Quad); i++ {
		result, _ = runCommand("./"+Quad[i], strX, strY)
		// fmt.Print(result)
		// fmt.Print(inputStr)
		// fmt.Print(inputStr)

		if result == inputStr && Quad[i] == "quadA" {
			str = append(str, "[quadA]"+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
			// fmt.Print("succed")
			// return str
		}

		if result == inputStr && Quad[i] == "quadB" {
			str = append(str, "[quadB]"+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
			// fmt.Print(len(str))
			// return str
		}

		if result == inputStr && Quad[i] == "quadC" {
			str = append(str, "[quadC]"+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
			// fmt.Print("succed")
			// fmt.Print(len(str))
			// return str
		}

		if result == inputStr && Quad[i] == "quadD" {
			str = append(str, "[quadD]"+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
			// fmt.Print("succed")
			// fmt.Print(len(str))
			// return str
		}

		if result == inputStr && Quad[i] == "quadE" {
			str = append(str, "[quadE]"+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
			// fmt.Print("succed")
			// fmt.Print(len(str))
			// return str
		}

	}
	// fmt.Print(len(str))
	if len(str) == 0 {
		fmt.Println("Not a quad function")
	}else{
		for i := 0; i <= len(str)-2; i++ {
		fmt.Print(str[i] + "||" + " ")
	}
	fmt.Println(str[len(str)-1])
	}
	

}

func printNbr(n int) string {
	res := ""

	if n/10 != 0 {
		res = res + printNbr(n/10)
	}
	mod := '0'
	for i := 0; i < (n % 10); i++ {
		mod++

	}
	if mod == '0' {
		res = res + "0"
	} else if mod == '1' {
		res = res + "1"
	} else if mod == '2' {
		res = res + "2"
	} else if mod == '3' {
		res = res + "3"
	} else if mod == '4' {
		res = res + "4"
	} else if mod == '5' {
		res = res + "5"
	} else if mod == '6' {
		res = res + "6"
	} else if mod == '7' {
		res = res + "7"
	} else if mod == '8' {
		res = res + "8"
	} else if mod == '9' {
		res = res + "9"
	}

	return res

}
