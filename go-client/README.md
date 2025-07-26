# Go Instagram Downloader Client

A simple command-line client to test the Go Instagram Downloader API. It allows you to download videos by providing an Instagram post URL.

## Features

- Fetches post metadata from the API.
- Downloads videos from Instagram URLs.
- Command-line interface for easy usage.
- Includes a test suite to verify the end-to-end download flow.

## Project Structure

```
go-client/
├── main.go             # Application entry point
├── main_test.go        # Test for the download flow
├── go.mod              # Go module definition
├── internal/
│   ├── api/
│   │   └── client.go   # Functions to interact with the downloader API
│   ├── types/
│   │   └── types.go    # Struct definitions for API responses
│   └── utils/
│       └── utils.go    # Utility functions (e.g., shortcode extraction)
└── README.md           # This file
```

## Installation & Running

### Prerequisites
- Go 1.21 or higher
- The Instagram Downloader API must be running.

### Setup
1. Navigate to the `go-client` directory:
```bash
cd go-client
```

2. Initialize Go modules:
```bash
go mod tidy
```

## Usage

### Running the Application
To download a video, run the application with the Instagram URL as an argument:
```bash
go run .\main.go https://www.instagram.com/reel/C0m_gs3y9Jk/
```

### Running Tests
To run the built-in test suite, which verifies the health of the API and the download process:
```bash
go test -v
```

### Building an Executable
You can compile the client into a single executable file for easy distribution and use.

1.  **Build the executable:**
    ```bash
    go build -o downloader.exe .
    ```
    This will create `downloader.exe` in the current directory.

2.  **Run the executable:**
    ```bash
    ./downloader.exe https://www.instagram.com/reel/DJeXBKNPFNM/?igsh=aTV1ajNmNXhtMDB3
    ``` 