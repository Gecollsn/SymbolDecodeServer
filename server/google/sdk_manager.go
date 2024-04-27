package google

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gecollsn.it/symparse/server/utils"
)

// download sdk via command line seems impossible, so it is enough to check if it exists.
type AndroidSDK struct {
	homePath string ""
}

// check whether sdk is good to go
func (sdk *AndroidSDK) CheckSdkExists() bool {
	ops := runtime.GOOS
	fmt.Println("operation system", ops)

	sdk.homePath = os.Getenv("ANDROID_HOME")
	if !utils.IsEmpty(sdk.homePath) {
		return true
	}

	defaultPath := filepath.Join(os.Getenv("HOME"), "Android", "Sdk")
	if _, err := os.Stat(defaultPath); err == nil {
		sdk.homePath = defaultPath
		return true
	}

	switch ops {
	case "darwin":
		fmt.Println("Running on Mac OS X")
	case "windows":
		fmt.Println("Running on Windows")
	case "linux":
		fmt.Println("Running on Linux")
	default:
		fmt.Println("Unknown operating system:", ops)
	}
	return sdk.homePath != ""
}

func (sdk *AndroidSDK) GetSdkHomePath() string {
	return ""
}

func (sdk *AndroidSDK) ListNdkVersions() []string {
	return nil
}

var sdk = new(AndroidSDK)

func GetAndroidSDK() *AndroidSDK {
	return sdk
}
