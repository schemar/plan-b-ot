# Plan-B-ot
Plan-B-ot is an integration for
[Mattermost](http://www.mattermost.org/) and
[slack](https://slack.com/) which allows for Scrum Planning Poker
inside the Mattermost or slack chat.

Plan-B-ot is written in [go](https://golang.org/).
Mattermost is an open source project: [Mattermost on GitHub](https://github.com/mattermost).

## Usage
In a channel, use the slash command (see setup) to interact with
the bot. If your slash command is called `/planbot`, you can issue the
following commands:
```
/planbot task <your taskname>
```
Creates a new task with zero votes and the given name

```
/planbot vote <your vote>
```
Sets your vote on the current task.
If you had a vote before, it  gets overriden by the new vote.

```
/planbot results
```
Orders plan-b-ot to print all the voting results in the specified
channel (see setup).

## Setup
### Mattermost/Slack
You need to setup two integrations: a slach command and an
incoming web hook

Setup a slash command so users can interact with plan-b-ot:
- Create a new slash command integration for your team.
- Pick the command you want to use, e.g. `planbot`.
- Specify the URL where your plan-b-ot will be running.
If your server is reachable at `example.com`, your port is `8786`, and
your route is `/planbot`: then set this to `example.com:8786/planbot`.
See also bot setup.
- Method is `POST`.
- Token: you need the token for the bot setup.

Setup an incoming web hook so plan-b-ot can post to your channel:
- Create a new incoming web hook for your team.
- Set the `channel` that you want plan-b-ot to post to.
- Webhook URL: you need the webhook URL for the bot setup.

### Bot
Copy `config.json.example` to `config.json`.
Edit the contents of `config.json` to setup your plan-b-ot:
- `Port`: The port on which plan-b-ot listens.
- `Route`: The URL route for plan-b-ot.
- `Token`: the token from the Mattermost/slack setup (slash command).
- `Webhook URL`: The webhook URL from the Mattermost/slack setup
(incoming web hook on Mattermost's/'slack's side).

Run plan-b-ot.
Now the server is running and waiting for input from Mattermost/slack.

## Contributions
Contributions are always welcome and do not have to be in a specific
format.

Simply create a pull request on GitHub.
