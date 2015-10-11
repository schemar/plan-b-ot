# Plan-B-ot
Plan-B-ot is an integration for [slack](https://slack.com/) which allows for
Scrum Planning Poker inside the slack chat.

Plan-B-ot is written in [go](https://golang.org/).

## Usage
In a slack channel, use the slash command (see setup) to interact with the bot.
If your slash command is called `/planbot`, you can issue the following
commands:
```
/planbot task <your taskname>
```
Creates a new task with zero votes and the given name

```
/planbot vote <your vote>
````
Sets your vote on the current task.
If you had a vote before, it  gets overriden by the new vote.

````
/planbot results
```
Orders plan-b-ot to print all the voting results in the specified channel
(see setup)

## Setup
### Slack
You need to setup two integrations in slack: a slach command and a
incoming web hook

Setup a slash command so users can interact with plan-b-ot:
- Create a new slash command integration for your slack team.
- Pick the command you want to use, e.g. `planbot`
- Specify the URL where your plan-b-ot will be running. If your server is reachable at `example.com`, your port is `8786`, and your route is `/planbot`: then set this to `example.com:8786/planbot`. See also bot setup.
- Method is `POST`
- Token: you need the token for the bot setup.

Setup an incoming web hook so plan-b-ot can post to your channel:
- Create a new incoming web hook for your slack team.
- Set the `channel` that you want plan-b-ot to post to
- Webhook URL: you need the webhook URL for the bot setup

### Bot
Copy `config.json.example` to `config.json`.
Edit the contents of `config.json` to setup your plan-b-ot:
- `Port`: The port in which plan-b-ot listens
- `Route`: The route for plan-b-ot
- `Token`: the token from teh slack setup (slash command)
- `Webhook URL`: The webhook URL from the slack setup (incoming web hook)

Run plan-b-ot.
Now the server is running and waiting for input from slack.
