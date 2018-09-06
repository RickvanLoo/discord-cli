# discord-cli
## Project is on hold
This project was started as an experiment for me to get familiar with Go. Back in the day the Discord API was also a bit simpler and Linux support was not really there unfortunately. Since then Linux support is there luckily, new features have been added to Discord that are not being supported by the cli currently.
I've personally moved on to other things, and since I do not really need a Discord CLI I will not be actively developing the client anymore.

If people feel like they want to actually manage the project and actively develop for it then contact me. Otherwise if you have an addition that improves the client, please do a pull request. I will look at it.  If there are any big problems for further feature development such as API breaks, you can also contact me and we will look if anything can still be done :)
______

Minimalistic Command-Line Interface for Discord

Master (Semi-Stable): [![Build Status](https://travis-ci.org/Rivalo/discord-cli.svg?branch=master)](https://travis-ci.org/Rivalo/discord-cli), Develop (Default Git Branch): [![Build Status](https://travis-ci.org/Rivalo/discord-cli.svg?branch=develop)](https://travis-ci.org/Rivalo/discord-cli)

Join our Discord Chat! https://discord.gg/0pXWCo5RQbVuFHDM

![I suck at English, while 256 colors is enough for everyone](screenshot.png)

<sub>Disclaimer: Currently only tested on Linux.</sub>

### How to Install the Master branch?
Currently the easiest working way to install is to use the Go tools. I'm looking at using GCCGO and makefiles to reduce installation steps, and make setting PATHS unnecessary.
* Install the Go Tools and setup the `$GOPATH` (There are loads of tutorial for this part)
* `$ go get -u github.com/Rivalo/discord-cli`
* Go to the `bin` folder inside your `$GOPATH`
* `./discord-cli`

For trying the develop branch, do a git checkout and reinstall the application.

### (Master) Configuration Settings
Configuration files are being stored in JSON format and are automatically created when you first run discord-cli. Do not change the 'key' value inside `{"key":"value"}`, this is the part that discord-cli uses for parsing, missing keys will definitely return errors.

| Setting       | Function         |
| ------------- |-------------|
| username      | Discord Username (emailaddress) |
| password      | Discord Password |
| messagedefault| (true or false) Display messages automatically|
| messages   | Amount of Messages kept in memory |

NOTE: The Configuration settings are likely to change. Breaking updates are stated in the release section. To solve problems, delete `~/.config/discord-cli/config.json` and restart discord-cli.

### (Master) Chat Commands
When inside a text channel, the following commands are available:

| Command       | Function         |
| ------------- |-------------|
| :q      | Quits discord-cli |
| :g      | Change listening Guild|
| :c      | Change listening Channel inside Guild |
| :m [n]      | Display last [n] messages: ex. `:m 2` displays last two messages |
