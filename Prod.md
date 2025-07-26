# Production Deployment on IONOS Ubuntu 24.04

This guide provides step-by-step instructions for deploying the Instagram Downloader bot and its API to a production environment on an IONOS VPS running Ubuntu 24.04.

## 1. Prerequisites

- An IONOS VPS with Ubuntu 24.04 installed.
- SSH access to your VPS.
- A domain or subdomain pointed to your VPS's IP address (optional, but recommended for HTTPS).
- `git` installed on your VPS (`sudo apt update && sudo apt install git`).
- `go` installed on your VPS (`sudo apt install golang-go`).

## 2. Clone the Repository

Connect to your VPS via SSH and clone the project repository:

```bash
git clone <your-repository-url>
cd <your-repository-directory>
```

## 3. Build the Applications

Build the `go-api` and `go-tgbot` applications for a Linux environment.

### Build `go-api`

```bash
cd go-api
go build -o instagram-api-linux
cd ..
```

### Build `go-tgbot`

```bash
cd go-tgbot
go build -o tgbot-linux
cd ..
```

You should now have two executable files: `go-api/instagram-api-linux` and `go-tgbot/tgbot-linux`.

## 4. Configure the Telegram Bot

The Telegram bot is configured using environment variables. Create a `.env` file in the `go-tgbot` directory.

```bash
cd go-tgbot
nano .env
```

Add the following content to the file, replacing the placeholder values with your actual data:

```
# Telegram Bot Token from BotFather
TELEGRAM_BOT_TOKEN=YOUR_TELEGRAM_BOT_TOKEN

# Base URL for the Instagram Downloader API
# This should be the public URL of your API server
API_BASE_URL=http://85.215.239.20:8080
```

Save the file and exit (`Ctrl+X`, `Y`, `Enter`).

## 5. Set up systemd Services

To ensure the applications run continuously and restart automatically, we will create `systemd` services for them.

### `go-api` Service

Create a new service file for the API:

```bash
sudo nano /etc/systemd/system/instagram-api.service
```

Add the following configuration. **Make sure to replace `/root/projects/mediav1/Insta1` with the actual path to your project directory if it differs.**

```ini
[Unit]
Description=Instagram Downloader API Service
After=network.target

[Service]
User=root
Group=root
WorkingDirectory=/root/projects/mediav1/Insta1/go-api
ExecStart=/root/projects/mediav1/Insta1/go-api/instagram-api-linux
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

### `go-tgbot` Service

Create a new service file for the Telegram bot:

```bash
sudo nano /etc/systemd/system/tgbot.service
```

Add the following configuration, again replacing `/root/projects/mediav1/Insta1` with the correct path if it differs:

```ini
[Unit]
Description=Telegram Instagram Downloader Bot Service
After=network.target instagram-api.service
Requires=instagram-api.service

[Service]
User=root
Group=root
WorkingDirectory=/root/projects/mediav1/Insta1/go-tgbot
ExecStart=/root/projects/mediav1/Insta1/go-tgbot/tgbot-linux
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

**Note:** The `After` and `Requires` directives in the `tgbot.service` ensure that the API service is started before the bot service.

## 6. Start and Enable the Services

Reload the `systemd` daemon to recognize the new services, then start and enable them.

```bash
sudo systemctl daemon-reload
sudo systemctl start instagram-api
sudo systemctl start tgbot
sudo systemctl enable instagram-api
sudo systemctl enable tgbot
```

## 7. Verify the Services

Check the status of the services to ensure they are running correctly:

```bash
sudo systemctl status instagram-api
sudo systemctl status tgbot
```

You can also view the logs for each service:

```bash
sudo journalctl -u instagram-api -f
sudo journalctl -u tgbot -f
```

## 8. Firewall Configuration (Optional but Recommended)

If you have `ufw` (Uncomplicated Firewall) enabled, make sure to allow traffic on the port your API is using (default is `8080`).

```bash
sudo ufw allow 8080/tcp
sudo ufw enable
```

## 9. Next Steps: Reverse Proxy and HTTPS

For a production environment, it is highly recommended to run the API behind a reverse proxy like Nginx and secure it with an SSL certificate from Let's Encrypt. This is a more advanced topic but provides significant security and performance benefits.

Congratulations! Your Instagram Downloader bot and API should now be running in production on your IONOS VPS. 