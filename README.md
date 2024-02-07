# OS library

## Candidate functions :

|   Func Name    |                   Prototype                   |                             Description                              |         Comments          |
|:--------------:|:---------------------------------------------:|:--------------------------------------------------------------------:|:-------------------------:|
|     Chown      |    Chown(path string, uid int, gid int) {}    |            Chown changes the owner and group of the file             | Does not work on Windows  |
|    ClearEnv    |                 ClearEnv() {}                 |                   Clears all environment variables                   |            N/A            |
|     Create     |            Create(name string) {}             |              Creates a file and returns the file object              |            N/A            |
|    Getegid     |               Getegid() int {}                |          Gets the effective group ID of the calling process          |            N/A            |
|     GetEnv     |         GetEnv(key string) string {}          |                   Returns an environment variable                    |            N/A            |
|    Geteuid     |               Geteuid() int {}                |          Gets the effective user ID of the calling process           |            N/A            |
|     Getgid     |                Getgid() int {}                |               Gets the group ID of the calling process               |            N/A            |
|  GetHostname   |            GetHostname() string {}            |                   Gets the hostname of the machine                   |            N/A            |
|     Getpid     |                Getpid() int {}                |              Gets the process ID of the calling process              |            N/A            |
|    Getppid     |               Getppid() int {}                |       Gets the process ID of the parent of the calling process       |            N/A            |
|     Getuid     |                Getuid() int {}                |               Gets the user ID of the calling process                |            N/A            |
| GetUserHomeDir |          GetUserHomeDir() string {}           |             Gets the home directory of the current user              |            N/A            |
|     Getwd      |               Getwd() string {}               |                  Gets the current working directory                  |            N/A            |
|     Mkdir      |             Mkdir(name string) {}             |                       Creates a new directory                        |            N/A            |
|    ReadDir     |       ReadDir(name string) []string {}        | Reads a directory and returns the names of the files and directories |            N/A            |
|    ReadFile    |        ReadFile(str string) string {}         |                    Returns the content of a file                     |            N/A            |
|     Remove     |          Remove(filename string) {}           |                            Removes a file                            |            N/A            |
|   RemoveAll    |           RemoveAll(path string) {}           |               Removes a directory and all its contents               |            N/A            |
|     SetEnv     |      SetEnv(key string, value string) {}      |                     Sets an environment variable                     |            N/A            |
|  SetEnvByFile  |       SetEnvByFile(filename string) {}        |            Sets an environment variable by reading a file            | Uses the format key=value |
|    UnsetEnv    |            UnsetEnv(key string) {}            |                    Unsets an environment variable                    |            N/A            |
|   WriteFile    | WriteFile(filename string, content string) {} |                 Writes a file with the given content                 |            N/A            |
