package Python

import (
	"fmt"
	"os/exec"
)

func CmdPython(filePath, newFilePath string)(err error)  {
	args := []string{"./Python/main.py", filePath, newFilePath}
	out, err:=exec.Command("python", args...).Output()
	if err != nil{
		fmt.Println("失败了", err)
		return
	}
	result := string(out)
	fmt.Println(result)
	return
}

