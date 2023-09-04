package tools

import (
	"fmt"
	"os/exec"
)

func Encrypt(input string, secret string) (output string, err error) {
	// openssl 命令
	openssl := fmt.Sprintf("echo %s | openssl enc -e -aes-256-cbc -a -nosalt -pbkdf2 -iv 1 -k %s", input, secret)

	outputByte, err := exec.Command("sh", "-c", openssl).CombinedOutput()
	if err != nil {
		return
	}

	output = string(outputByte)
	return
}

func Decrypt(input string, secret string) (output string, err error) {
	// openssl 命令
	openssl := fmt.Sprintf("echo %s | openssl enc -d -aes-256-cbc -a -nosalt -pbkdf2 -iv 1 -k %s", input, secret)

	outputByte, err := exec.Command("sh", "-c", openssl).CombinedOutput()
	if err != nil {
		return
	}

	output = string(outputByte)
	return
}
