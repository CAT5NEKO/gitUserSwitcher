package utils

import (
	"fmt"
	"go-gituser/internal/models"
	"go-gituser/state"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

func ReadFileData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, errors.Wrap(err, "ReadFileData")
	}
	return data, nil
}

func ReadAccountsData(account models.Accounts) {
	fmt.Println("Hello, this is your accounts data")
	fmt.Println("")
	if account.PersonalUsername == "" {
		fmt.Println("🏠 | You have no personal account defined")
	} else {
		fmt.Println("🏠 | Personal Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.PersonalUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.PersonalEmail)
	}
	fmt.Println("")
	if account.SchoolUsername == "" {
		fmt.Println("📚 | You have no school account defined")
	} else {
		fmt.Println("📚 | School Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.SchoolUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.SchoolEmail)
	}
	fmt.Println("")
	if account.WorkUsername == "" {
		fmt.Println("💻 | You have no work account defined")
	} else {
		fmt.Println("💻 | Work Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.WorkUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.WorkEmail)
	}
	fmt.Println("")

}

func ReadCurrentAccountData(name string, email string, mode string) {
	fmt.Println("You are on the " + color.CyanString(mode) + " acccount")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", name)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", email)
}

func ReadUnsavedGitAccount(name string, email string) {
	fmt.Println("You are using the following account")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", name)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", email)

	fmt.Println("This account is " + color.YellowString("unsaved") + ". Run <gituser config> to save it to a " + color.CyanString("mode"))
}

func GetAccountsDataFile() (string, error) {
	dataFile, err := getLocalFile(accountsDataFileName)
	if err != nil {
		return "", errors.Wrap(err, "GetAccountsDataFile")
	}

	return dataFile, nil
}

func ensureLocalConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".config", "gituser")
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return configDir, nil
}

func getLocalFile(filename string) (string, error) {
	localConfigDir, err := ensureLocalConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "checkFile - UserHomeDir")
	}

	localDataFile := filepath.Join(localConfigDir, filename)
	_, err = os.Stat(localDataFile)
	if os.IsNotExist(err) {
		createdFile, err := os.Create(localDataFile)
		if err != nil {
			return "", errors.Wrap(err, "checkFile - create")
		}

		_, err = createdFile.Write([]byte("{}"))
		if err != nil {
			return "", errors.Wrap(err, "checkFile - write")
		}

	}
	return localDataFile, nil
}

func GitUsernameIsUnsaved(name string) bool {
	return state.SavedAccounts.PersonalUsername != name ||
		state.SavedAccounts.WorkUsername != name || state.SavedAccounts.SchoolUsername != name
}

func GitEmailIsUnsaved(email string) bool {
	return state.SavedAccounts.PersonalEmail != email ||
		state.SavedAccounts.WorkEmail != email || state.SavedAccounts.SchoolEmail != email
}
