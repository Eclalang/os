package os

import (
	"os"
	"strings"
)

// Chown changes the owner and group of the file
func Chown(path string, uid int, gid int) {
	os.Chown(path, uid, gid)
}

// ClearEnv clears all environment variables
func ClearEnv() {
	os.Clearenv()
}

// Create creates a file and returns the file object
func Create(name string) {
	file, err := os.Create(name)
	if err != nil {
		return
	}
	file.Close()
	return
}

// Getegid gets the effective group ID of the calling process
func Getegid() int {
	return os.Getegid()
}

// GetEnv gets an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// Geteuid gets the effective user ID of the calling process
func Geteuid() int {
	return os.Geteuid()
}

// Getgid gets the group ID of the calling process
func Getgid() int {
	return os.Getgid()
}

// GetHostname gets the hostname of the machine
func GetHostname() string {
	name, _ := os.Hostname()
	return name
}

// Getpid gets the process ID of the calling process
func Getpid() int {
	return os.Getpid()
}

// Getppid gets the process ID of the parent of the calling process
func Getppid() int {
	return os.Getppid()
}

// Getuid gets the user ID of the calling process
func Getuid() int {
	return os.Getuid()
}

// GetUserHomeDir gets the home directory of the current user
func GetUserHomeDir() string {
	home, _ := os.UserHomeDir()
	return home
}

// Getwd gets the current working directory
func Getwd() string {
	wd, _ := os.Getwd()
	return wd
}

// Mkdir creates a new directory
func Mkdir(name string) {
	os.Mkdir(name, 0750)
}

// ReadDir reads a directory and returns the names of the files and directories
func ReadDir(name string) []string {
	files, err := os.ReadDir(name)
	if err != nil {
		return nil
	}
	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}
	return names
}

// ReadFile reads a file and returns the contents
func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}

// Remove removes a file
func Remove(filename string) {
	os.Remove(filename)
}

// RemoveAll removes a directory and all its contents
func RemoveAll(path string) {
	os.RemoveAll(path)
}

// SetEnv sets an environment variable
func SetEnv(key string, value string) {
	os.Setenv(key, value)
}

// SetEnvByFile sets an environment variable by reading a file
func SetEnvByFile(filename string) error {
	str := ReadFile(filename)
	if str == "" {
		return nil
	}
	for _, line := range strings.Split(str, "\n") {
		if strings.Contains(line, "=") {
			part := strings.Split(line, "=")
			os.Setenv(part[0], part[1])
		}
	}
	return nil
}

// UnsetEnv unsets an environment variable
func UnsetEnv(key string) {
	os.Unsetenv(key)
}

// WriteFile writes a file with the given content
func WriteFile(filename string, content string) {
	file, _ := os.Create(filename)
	file.WriteString(content)
	file.Close()
}
