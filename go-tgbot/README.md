# Go Telegram Bot

A simple Telegram bot that echoes messages back to the user. This project serves as a basic template for building more complex bots using the `go-telegram/bot` library.

## Features

- Responds to any text message by echoing it back.
- Uses a structured configuration for the bot token.
- Built with the zero-dependency `go-telegram/bot` framework.
- Gracefully handles shutdown on interrupt signals.

## Project Structure

```
go-tgbot/
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── internal/
│   ├── config/
│   │   └── config.go   # Handles loading configuration (e.g., bot token)
│   └── types/
│       └── types.go    # Struct definitions for configuration
└── README.md           # This file
```

## Installation & Running

### Prerequisites
- Go 1.18 or higher
- A valid Telegram Bot Token from [BotFather](https://t.me/botfather)

### Setup
1.  Navigate to the `go-tgbot` directory:
    ```bash
    cd go-tgbot
    ```

2.  Install the necessary dependencies:
    ```bash
    go mod tidy
    ```

3.  Configure your bot token by editing `internal/config/config.go` and replacing the placeholder token with your own:
    ```go
    // in internal/config/config.go
    cfg.TelegramBotToken = "YOUR_REAL_BOT_TOKEN" 
    ```

## Usage

### Running the Application
To start the bot, run the following command from the `go-tgbot` directory:
```bash
go run .\main.go
```
Once the bot is running, send it any message on Telegram, and it will reply.

### Building an Executable
You can compile the bot into a single executable file.

1.  **Build the executable:**
    ```bash
    go build -o tgbot.exe .
    ```
    This will create `tgbot.exe` in the current directory.

2.  **Run the executable:**
    ```bash
    ./tgbot.exe
    ``` 