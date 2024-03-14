# Go CLI Application

## What is CLI Application and why we use them?
* CLI application is text-based application that can run on your command. Such as GitHub and npm.
* Unlike GUI, CLI is light-weight and have more performance than GUI because It's contains only text not graphics.
* CLI can be complied to binary files which can also run cross-platform. 
* Easy to code.

## Why we create CLI with Go?

- Spf13/cobra
- Spf13/viper
- Urfave and cli
- Delve
- Chzyer and readline

GitHub CLI also use Go to do CLI Application

First we create go mod init in each folder like this \
`go mod init github.com/Celesca/Lec02_simple_app `

## Why we use Cobra framework
Cobra is fastest Static Site Generator (SSG) and a lot more features such as

1. Nested commands (Can implement complex features) - because the standard libraries is very hard to implement
2. Powerful flags (POSIX) - Particular command (Options and Description)
3. Customize Help
4. Shell Auto Completion (You can tab to generate auto completion)
