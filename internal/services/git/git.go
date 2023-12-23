package git

import (
	"fmt"
	"go-gituser/utils/logger"
	"os/exec"
	"strings"
)

func CurrentAccount() (string, string) {
	cmdName := exec.Command("git", "config", "--global", "user.name")
	cmdEmail := exec.Command("git", "config", "--global", "user.email")

	name, nameErr := cmdName.CombinedOutput()
	if nameErr != nil {
		logger.PrintErrorExecutingMode()
	}

	email, emailErr := cmdEmail.CombinedOutput()
	if emailErr != nil {
		logger.PrintErrorExecutingMode()
	}

	return strings.TrimSpace(string(email)), strings.TrimSpace(string(name))
}

func SetAccount(name, email string) {
	setConfigName(name)
	setConfigEmail(email)
}

func setConfigName(name string) {
	cmd := exec.Command("git", "config", "--global", "user.name", name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("👤 " + name + " was set as username")
}

func setConfigEmail(email string) {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	_, err := cmd.CombinedOutput()
	if err != nil {
		logger.PrintErrorExecutingMode()
	}
	fmt.Println("📧 " + email + " was set as email")
}
