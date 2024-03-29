package os

import (
	"os"
	"reflect"
	"testing"
)

func TestChown(t *testing.T) {
	oldUid := Getuid()
	oldGid := Getgid()
	Chown("unit_test_files/test.txt", 3, 3)
	newUid := Getuid()
	newGid := Getgid()
	if oldUid != newUid || oldGid != newGid {
		t.Error("Expected uid and gid to change, but it did not")
	}
}

/*
func TestClearEnv(t *testing.T) {
	ClearEnv()
	expected := os.Getenv("ecla")
	actual := GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

*/

func TestCreate(t *testing.T) {
	working := false
	Create("unit_test_files/newfile.txt")
	files := ReadDir("unit_test_files/")
	for _, file := range files {
		if file == "newfile.txt" {
			working = true
		}
	}
	if !working {
		t.Error("Expected to find newfile.txt, but it was not found")
	}

	working = false
	fileName := "no??/?extension"
	Create("unit_test_files/" + fileName)
	files = ReadDir("unit_test_files/")
	for _, file := range files {
		if file == fileName {
			working = true
		}
	}
	if working {
		t.Errorf("Did not expect to create a file named %s, but it was found", fileName)
	}
}

func TestGetegid(t *testing.T) {
	expected := os.Getegid()
	actual := Getegid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetEnv(t *testing.T) {
	expected := os.Getenv("ecla")
	actual := GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGeteuid(t *testing.T) {
	expected := os.Geteuid()
	actual := Geteuid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetgid(t *testing.T) {
	expected := os.Getgid()
	actual := Getgid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetHostname(t *testing.T) {
	expected, _ := os.Hostname()
	actual := GetHostname()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetpid(t *testing.T) {
	expected := os.Getpid()
	actual := Getpid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetppid(t *testing.T) {
	expected := os.Getppid()
	actual := Getppid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetuid(t *testing.T) {
	expected := os.Getuid()
	actual := Getuid()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetUserHomeDir(t *testing.T) {
	expected, _ := os.UserHomeDir()
	actual := GetUserHomeDir()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetwd(t *testing.T) {
	expected, _ := os.Getwd()
	actual := Getwd()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestMkdir(t *testing.T) {
	Mkdir("unit_test_files/tempodirectory")
	files := ReadDir("unit_test_files/")
	for _, file := range files {
		if file == "tempodirectory" {
			Remove("unit_test_files/tempodirectory/")
			return
		}
	}
	t.Error("Expected to find tempodirectory, but it was not found")
}

func TestReadDir(t *testing.T) {
	expect, _ := os.ReadDir("unit_test_files/")
	var expected []string
	for _, file := range expect {
		expected = append(expected, file.Name())
	}
	actual := ReadDir("unit_test_files/")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actualCrash := ReadDir("unit_test_not_exist/")
	if actualCrash != nil {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	expect, _ := os.ReadFile("unit_test_files/test.txt")
	expected := string(expect)
	actual := ReadFile("unit_test_files/test.txt")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actualCrash := ReadFile("unit_test_files/notexist.txt")
	if actualCrash != "" {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRemove(t *testing.T) {
	Create("unit_test_files/newfile.txt")
	Remove("unit_test_files/newfile.txt")
	files := ReadDir("unit_test_files/")
	for _, file := range files {
		if file == "newfile.txt" {
			t.Error("Expected newfile.txt to be removed, found it")
		}
	}
}

func TestRemoveAll(t *testing.T) {
	Mkdir("unit_test_files/tempodirectory")
	Create("unit_test_files/tempodirectory/newfile.txt")
	RemoveAll("unit_test_files/tempodirectory")
	files := ReadDir("unit_test_files/")
	for _, file := range files {
		if file == "tempodirectory" {
			t.Error("Expected tempodirectory to be removed, found it")
		}
	}
}

func TestSetEnv(t *testing.T) {
	SetEnv("ecla", "Hello we are Ecla Team!")
	expected := os.Getenv("ecla")
	actual := GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSetEnvByFile(t *testing.T) {
	err := SetEnvByFile("unit_test_files/.env")
	if err != nil {
		return
	}
	expected := os.Getenv("ecla")
	actual := GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	err = SetEnvByFile("unit_test_files/empty.env")
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	err = SetEnvByFile("unit_test_files/error.env")
	if err != nil {
		t.Errorf("Expected an error, got %v", err)
	}

}

func TestUnsetEnv(t *testing.T) {
	SetEnv("ecla", "Hello we are Ecla Team!")
	UnsetEnv("ecla")
	expected := os.Getenv("ecla")
	actual := GetEnv("ecla")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestWriteFile(t *testing.T) {
	WriteFile("unit_test_files/test.txt", "Hello we are Ecla Team!")
	expect, _ := os.ReadFile("unit_test_files/test.txt")
	expected := string(expect)
	actual := ReadFile("unit_test_files/test.txt")
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
