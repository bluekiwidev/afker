# afker

Holds down your left mouse button and waits a specific amount of time until switching to the next hotbar slot when your pickaxe breaks. Perfect for AFK mining in Minecraft!

## Features

- Cross-platform support (Windows, macOS, Linux)
- Automatic hotbar slot switching (slots 1-9)
- Configurable timing for different tools
- Simple and easy to use

## Installation

### Prerequisites

You need [Go](https://go.dev/dl/) installed on your system.

**For Windows users:**
1. Download and install Go from https://go.dev/dl/
2. During installation, make sure "Add to PATH" is selected
3. After installation, open a new Command Prompt or PowerShell window to verify: `go version`

**For Windows users (optional):**
- If you want to build from source, you'll need GCC (via MinGW-w64 or TDM-GCC)
- However, you can also download pre-built binaries from the releases page

### Quick Start (Windows)

1. **Clone the repository:**
   ```powershell
   git clone https://github.com/bluekiwidev/afker.git
   cd afker
   ```

2. **Get dependencies:**
   ```powershell
   go mod tidy
   ```

3. **Build the application:**
   ```powershell
   go build -o afker.exe .
   ```

4. **Run the application:**
   ```powershell
   .\afker.exe
   ```
   Or simply double-click `afker.exe` in File Explorer.

### Quick Start (macOS/Linux)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/bluekiwidev/afker.git
   cd afker
   ```

2. **Get dependencies:**
   ```bash
   go mod tidy
   ```

3. **Build the application:**
   ```bash
   go build -o afker .
   ```

4. **Run the application:**
   ```bash
   ./afker
   ```

## Usage

1. Place pickaxes in your hotbar slots (slots 1-9)
2. **Select hotbar slot 1 in-game** (the script assumes you start with slot 1)
3. Position your mouse where you want to mine
4. Run the application
5. You have 10 seconds to switch back to your game window
6. The script will:
   - Hold down the left mouse button
   - Wait for the configured time (default 144 seconds)
   - Automatically switch to the next hotbar slot (2, then 3, then 4, etc.)
   - Continue until all pickaxes in slots 1-9 are used

### Customization

Edit the constants in `main.go` to customize behavior:

```go
const min = 1         // First hotbar slot with a pickaxe
const max = 9         // Last hotbar slot with a pickaxe
const sleepTime = 144 // Seconds until switching to next slot
```

## Troubleshooting (Windows)

### "robotgo" build errors
If you encounter build errors related to robotgo, you may need to install GCC:
- Download and install TDM-GCC from https://jmeubank.github.io/tdm-gcc/
- Or install MinGW-w64 from https://www.mingw-w64.org/

### Permission Issues
If the script doesn't control your mouse/keyboard:
- Run the application as Administrator (right-click → "Run as administrator")
- Some games may block input automation - ensure your game allows it

### Application doesn't start
- Make sure you're running the application from Command Prompt or PowerShell, not just double-clicking (so you can see any error messages)
- Verify Go is properly installed: `go version`

## How It Works

The application uses [robotgo](https://github.com/go-vgo/robotgo) for cross-platform mouse and keyboard automation. It:
1. Holds down the left mouse button
2. Waits for the configured time (default: 144 seconds)
3. Presses the next hotbar number key (1-9)
4. Repeats until all slots are used

## License

This project is provided as-is for educational and personal use.
