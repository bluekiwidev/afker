# afker
Holds down your left mouse button and waits a specific amount of time until switching to the next hotbar slot when your picaxe breaks

## Install
How to install:
I'm too lazy to put instructions for macos and linux but it does work with it.
### Pre reqs
You need [go](https://go.dev/dl/)

### Clone the repo & cd
```bash
https://github.com/bluekiwidev/afker.git
cd afker
```

### Get deps
```bash
go mod tidy
```

### Build the app
```bash
go build -o afker.exe .
```

### Run the script
```bash
run afker.exe
```
