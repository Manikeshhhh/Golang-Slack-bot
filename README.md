# Golang-Slack-bot

### How this bot works
This bot uses a slack socket module built by shomali

At this time, it is very simple bot that uses slack app and bot key and connect to your slack workspace,
you can install this bot in your slack workspace. 

You can call this bot by mentioning this bot in your slack workspace and sending required commands.
At this time if you send ping it simply send pong back but you can definitely add more functions according 
to your need.


### How to install

1. Create slack workspace and go to slack api and create a app with all the required permissions
2. Get the app and bot api key
3. Paste it in the main the go file

```
	os.Setenv("SLACK_BOT_TOKEN","xoxb-********-*******-**********")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-***********-********-*******************************")
```
you can create and env file and do that as well
4. Get the slacker Framework by 
```
root@x71no: go get github.com/shomali11/slacker
``` 
5. Customize the bot according to your neeed
6. Run the bot by 
```
root@x71no: go build main.go
root@x71no: go run main.go
```

### Thank you