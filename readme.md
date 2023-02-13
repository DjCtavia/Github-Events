# Github Events

Listen to Github events with self hosted server through the application,  
and send the formatted event through a you registered discord hook.

## Installation

Make sure [GO is installed](https://go.dev) on your computer.

Clone the repository:

```bash
git clone https://github.com/DjCtavia/Github-Events.git
```

Move to the application directory:

```bash
cd Github-Events
```

Run the application:

```bash
go run main.go -hookUrl YourUniqueID@YourDiscordWebhookURL
```

You can add multiple `-hookUrl` at a time.

## Flags

- `-help`, `--help` : Show help for the application.
- `-hookUrl`, `--hookUrl` *\<uniqueID\>*@*\<discordHookURL\> : Link a discord hook to an unique ID.
- `-serverPort`, `--serverPort` *\<PortToListen\>* : Set a port for the server receiving Github Events.

## Contributing

Everyone is welcome to contribute to this project. Here are a few steps to get started:

1. Fork the project
2. Create a new branch (`git checkout -b feature/awesome-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/awesome-feature`)
5. Open a pull request

I appreciate your contributions to make this project better!