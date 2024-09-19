# client_trivia

Fetches the master server list with all active servers that are not full, sorted by number of players, and connects to the first one.
Fetches new questions from an openly accessible trivia API and can be started with the chat command `!trivia` in order to print the first question.
Users can answer with `1`, `2`, `3`, or `4` to select the answer they think is correct or just write the answer.

The command `!score` will show your current score.
The command `!top` will show the top players.
The command `!leave` will tell the bot to leave the server.
The `!help` or `!h` command will show all available commands.

```shell
go build .
./client_trivia
```

