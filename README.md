# Chitty Chat

## Table of contents

- [Chitty Chat](#chitty-chat)
- [Installing and running the code](#installing-and-running-the-code)
  - [Downloading Go](#downloading-go)
  - [Downloading the project](#downloading-the-project)
- [Usage](#usage)
  - [Set-up server](#set-up-server)
  - [Set-up client](#set-up-client)
  - [Exit](#exit)

## Installing and running the code

### Downloading go

- Install the latest version of go, from the official go website
- Insure that go is installed to path on your device
    - Test it by running the command `go version`
    - Should output something like the following `go version go1.23.0 darwin/arm64``

### Downloading the project

- Download this repository as a `.zip` file
- Unzip the the zip file
- Naviagate to the unzipped file and open it in the terminal
- Once in the directory, run `go mod download` to download all dipendencies
- Then run `go build`to build the program
- The run `go run .`to run the program

### Usage

#### Set-up server
When running the program you will be presented with a prompt:
```
Select an option: 
  ▸ Start Server
    Start new Client
    Exit
````

Selecting `Start Server` will present you with the following:

```
2024/10/28 22:11:09 Setting up server on port: :4002 with the name Chitty-Chat-Server
2024/10/28 22:11:09 Starting gRPC server with name: Chitty-Chat-Server
2024/10/28 22:11:09 Server is now listening on port: :4002
````

This indicates a server has been setup on the default port that is 4002.

#### Set-up client

Selecting `Start new Client` from the prompt will result in the following:
````
✔ Start new Client
✔ Enter desired name: █
✔ Enter server address: █
 ````
You can enter the desired username as anything, and the server address as `localhost:4002`

When this is done the following will be displayed:
````
2024/10/28 22:14:50 Joining chat as user: TestClient on time 0...
2024/10/28 22:14:50 Welcome! You just joined the chat with status: TestClient at time 1
````

#### Exit

If you choose exit, then the program will output the following:

````
✔ Exit
2024/10/28 22:15:58 Exiting...
exit status 1
````
