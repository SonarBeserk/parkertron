
# parkertron

![Parkertron Logo](images/parkertron_logo.png)

A simple discord chat bot with a simple configuration. Written using [discordgo](https://github.com/bwmarrin/discordgo).


# Large changes to the keyword/command configs

This most recent update will break your current keyword and command configs by moving the response one level deeper. This is to allow for reactions to be added on keywords. The bot can also mention a user for matched using the &user& string in the keyword response.

Reactions need to be copied in as a unicode emoji character.

### Old:
```
keyword:
  help:
    - "Please check the github page at <https://github.com/parkervcp/parkertron>"
    - "The default config is a good example of how to set commands up. Try `.help command`"
    - "My base chat parsing function is also available. Try `.help keyword` for more info"
```

### New:
```
keyword:
  help:
    reaction:
      - "💪"
    response:
      - "&user& Please check the github page at <https://github.com/parkervcp/parkertron>"
      - "The default config is a good example of how to set commands up. Try `.help command`"
      - "My base chat parsing function is also available. Try `.help keyword` for more info"
```


### Requirements:
tesseract-ocr w/ english trainign files (May support other languages but has not been tested.)
libleptonica (for tesseract)

Working on adding other services and additions.

The checklist so far

- [x] Full OSS release

- [x] Read config files
  - [x] Separate service config (currently only discord)

- [ ] support multiple services
  - [x] Discord
  - [ ] Slack
  - [x] IRC

#### Discord specific support
  - [x] Support @mentions for the bot
  - [x] Use @mentions for other users
  - [x] Watch for @mentions on groups


#### General support
- [x] Get inbound messages
  - [x] Listen to specific channels
  - [x] Listen for mentions

- [x] Respond to inbound messages
  - [x] respond to commands with prefix
  - [x] respond to key words/phrases
  - [x] Comma separated lists
  - [ ] Separate server commands

- [x] Image parsing
  - [x] image from url support
    - [x] png support
    - [x] jpg support
  - [x] direct embedded images

- [x] Respond with correct output from configs

- [x] Respond with multi-line output in a single message

- [x] Impliment blacklist/whitelist mode (Blacklist by User ID only)

- [ ] Mitigate spam with cooldown per user/channel/global
  - [ ] global cooldown
  - [ ] channel cooldown
  - [ ] user cooldown

- [x] Permissions
  - [ ] Permissions management

- [ ] Logging
  - [ ] Log user join/leave 
  - [x] Log chats (only logs channels it is watching to cut on logging)
  - [ ] Log edits (show original and edited)
  - [ ] Log chats/edits to separate files/folders
  
- [ ] Join voice channels
  - [ ] Play audio from links


So far I have the chat bot part down with no limiting or administration.

Configuration is done in yaml/json.  
If you have a Bot account already you can add the token and client ID's on your own.  
If you don't you will need to set your own account up.

The "owner" option in the configs is basically a super admin that will not be able to be blacklisted.

The prefix is the command prefix and is customizable.  
Set to "." by default it can be changed to whatever you want.


The Commands set up is simple and is also in json.  
See the commands.json for examples.  
