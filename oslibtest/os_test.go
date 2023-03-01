package oslibtest

import (
	"fmt"
	PackageOs "github.com/Eclalang/os"
	"os"
	"reflect"
	"testing"
)

func TestChown(t *testing.T) {
	PackageOs.Chown("test.txt", 0, 0)
	uid := PackageOs.Getuid()
	gid := PackageOs.Getgid()
	if uid != 0 || gid != 0 {
		t.Error("Expected uid and gid to change, but it did not")
	}
}

func TestClearEnv(t *testing.T) {
	PackageOs.ClearEnv()
	expected := os.Getenv("ecla")
	actual := PackageOs.GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestCreate(t *testing.T) {
	PackageOs.Create("newfile.txt")
	files := PackageOs.ReadDir(".")
	for _, file := range files {
		if file == "newfile.txt" {
			return
		}
	}
	t.Error("Expected to find newfile.txt, but it was not found")
}

func TestGetegid(t *testing.T) {
	expected := os.Getegid()
	actual := PackageOs.Getegid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetEnv(t *testing.T) {
	expected := os.Getenv("ecla")
	actual := PackageOs.GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGeteuid(t *testing.T) {
	expected := os.Geteuid()
	actual := PackageOs.Geteuid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetgid(t *testing.T) {
	expected := os.Getgid()
	actual := PackageOs.Getgid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetHostname(t *testing.T) {
	expected, _ := os.Hostname()
	actual := PackageOs.GetHostname()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetpid(t *testing.T) {
	expected := os.Getpid()
	actual := PackageOs.Getpid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetppid(t *testing.T) {
	expected := os.Getppid()
	actual := PackageOs.Getppid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetuid(t *testing.T) {
	expected := os.Getuid()
	actual := PackageOs.Getuid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetUserHomeDir(t *testing.T) {
	expected, _ := os.UserHomeDir()
	actual := PackageOs.GetUserHomeDir()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetwd(t *testing.T) {
	expected, _ := os.Getwd()
	actual := PackageOs.Getwd()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestMkdir(t *testing.T) {
	PackageOs.Mkdir("tempodirectory")
	files := PackageOs.ReadDir(".")
	for _, file := range files {
		if file == "tempodirectory" {
			return
		}
	}
	t.Error("Expected to find tempodirectory, but it was not found")
}

func TestReadDir(t *testing.T) {
	expect, _ := os.ReadDir(".")
	var expected []string
	for _, file := range expect {
		expected = append(expected, file.Name())
	}
	actual := PackageOs.ReadDir(".")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	expect, _ := os.ReadFile("test.txt")
	expected := string(expect)
	actual := PackageOs.ReadFile("test.txt")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRemove(t *testing.T) {
	PackageOs.Create("newfile.txt")
	PackageOs.Remove("newfile.txt")
	files := PackageOs.ReadDir(".")
	for _, file := range files {
		if file == "newfile.txt" {
			t.Error("Expected newfile.txt to be removed, found it")
		}
	}
}

func TestRemoveAll(t *testing.T) {
	PackageOs.Mkdir("tempodirectory")
	PackageOs.Create("tempodirectory/newfile.txt")
	PackageOs.RemoveAll("tempodirectory")
	files := PackageOs.ReadDir(".")
	for _, file := range files {
		if file == "tempodirectory" {
			t.Error("Expected tempodirectory to be removed, found it")
		}
	}
}

func TestSetEnv(t *testing.T) {
	PackageOs.SetEnv("ecla", "Hello we are Ecla Team!")
	expected := os.Getenv("ecla")
	actual := PackageOs.GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// DOES NOT WORK
func TestSetEnvByFile(t *testing.T) {
	PackageOs.SetEnvByFile(".env")
	expected := os.Getenv("ecla")
	actual := PackageOs.GetEnv("ecla")
	fmt.Println(expected)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestUnsetEnv(t *testing.T) {
	PackageOs.SetEnv("ecla", "Hello we are Ecla Team!")
	fmt.Println("Environment variable ecla is set to: ", os.Getenv("ecla"))
	PackageOs.UnsetEnv("ecla")
	expected := os.Getenv("ecla")
	actual := PackageOs.GetEnv("ecla")
	fmt.Println("Environment variable ecla is unset: ", os.Getenv("ecla"))
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestWriteFile(t *testing.T) {
	PackageOs.WriteFile("test.txt", "Hello we are Ecla Team!")
	expect, _ := os.ReadFile("test.txt")
	expected := string(expect)
	actual := PackageOs.ReadFile("test.txt")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
