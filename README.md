# Go Slack Age Bot

This is a simple Slack bot written in Go. It calculates the age of a user based on the year of birth they provide.

## Features

- Calculates age based on the year of birth
- Responds with a user-friendly message

## Installation

1. Clone this repository.
2. Create a `.env` file in the project directory.
3. Add the following environment variables to the `.env` file:
   - `SLACK_BOT_TOKEN` - The Slack bot token.
   - `SLACK_APP_TOKEN` - The Slack app token.
4. Navigate to the project directory.
5. Run `go build` to compile the project.

## Usage

1. Start the bot by running the compiled binary.
2. In Slack, use the command `My yob is <year>` to tell the bot your year of birth.
3. The bot will respond with your calculated age.

## Packages

1. "github.com/joho/godotenv"
2. "github.com/shomali11/slacker"

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
