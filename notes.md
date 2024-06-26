# Nofap bot
FriendlyBroccolli golang

Build android: CC=/home/titi/Android/Sdk/ndk/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build

Build and push: CC=/home/titi/Android/Sdk/ndk/android-ndk-r26b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build ; adb push "/home/titi/Dev/telegram/friendly-brocolli/main" /data/local/tmp ; adb push "/home/titi/Dev/telegram/friendly-brocolli/motivation" /data/local/tmp/motivation ; adb shell "chmod +x /data/local/tmp/main" ; adb shell "data/local/tmp/main"

Android build: Change motivationPath in config to absolute path

Get android device cpu architecture: adb shell getprop ro.product.cpu.abi

# To-do
[x] Middleware message count
[x] Bot.SetCommands()
[x] config.yml
[x] Add inline keyboard (/new)
[x] Finish getRank()
[x] Remove replace in go.mod
[x] Remove comment when issue fixed (i18n.SetDefaultLanguage("en-US")) // removed line
[x] Fix cauliflower and integrate
[x] /start
[x] /new
[x] /check
[x] /task
[x] /motivation
[x] /profile [@user=me]
[x] /account
[x] /ranks
[x] /ranks [rank]
[x] /help (channels, ~~donation~~)
[x] /add-task
[x] Fix profile and /check relapsed update db doesn't work
[x] Fix profile public entries markup
[x] Task max 3/day
[x] Change /motivations structs to db
[x] Finir /account activity
[x] Add resources to /help (easypeasymethod)
[x] Add images to motivation/
[x] Add /update to update motivation folder
[x] Update config notes.md
[-] Make reply keyboard markup (to uncomment -> work with .Edit()) (removed)
[-] Add tasks to config.yml
[-] /account download edit message instead of sending new one
[ ] README.md
[ ] Move tasks to db
[ ] Move motivations to db
[ ] Add map[motivation id]telebot.image
[ ] Add custom language
[x] Use layout.yml
[ ] Add motivation path to layout and update update()
[x] Add locales logic
[ ] Fix empty message (colon conflict in locales)
[ ] Remove comment (l.212 - 217)
[-] Split commands into different files

# Structure
Commands:
- /start -> tutorial
- /new -> new journey (days, save to db, rank system, update to db)
- /check -> new entry (max 3/day, relapse?, note, text, public?, save to db)
- /task -> random task to complete (max 3/day, completed?, save to db)
- /motivation -> random image
- /motivation list -> list categories
- /motivation [id] -> image id
- /motivation [category] -> random image from category
- /profile [@user=me] -> total score, current journey (start, days, rank, next rank, n. entries, n. tasks, score), all journeys (average length, total days, total entries), public entries (callback query button)
- /account -> score, rank, next rank, all entries, activity (new, check (id, note, relapse?), task), activity/journey, download
- /ranks -> ranks system overview
- /ranks [rank] -> full rank list
- /fix -> fix missing user
- /help -> command list, bot channel, personal channel, stats (users, uptime, messages count) contact, donation

Admin commands:
- /dummy -> make dummy user for test purpose
- /update -> update motivation table in database
- /add-task -> create new task and save into db

Reply markup:
- /new -> check, task, motivation, account
- /check relapsed -> new, motivation, account

Motivation filename:
- pack.packplace.category.languagecode.extension
- id.category.languagecode.extension
- id/pack must be unique
- category must not be equal to "list"

Score system:
- 1 point/check-in (3 checks max/day)
- 2 points/day
- 2-10 points/task (3 task max/day)

Config:
- Token: str
- Timeout: int
- SetCommands: bool (true only once)
- Database: string (db path -> :memory:)
- MotivationPath: string (motivation folder path without /)
- Owners: []int64 (users id allowed to run admin commands)
- NofapChannel: string (t.me link)
- PersonalChannel: string (t.me link)

Database:
Journey:
- ID (pk) - int
- User id - int
- Rank system - int
- Start - time.Time
- End - time.Time
- Text - str
Entry:
- ID (pk) - int
- Journey id (foreign key) - int
- Date (autodate) - time.Time
- Is public? - bool
- Note - int
- Text - str
Task:
- ID (pk) - int
- Journey id (foreign key) - int
- Date (autodate) - time.Time
- Task id

Locales:
- Use double indentation to escape colons (:)

# Installation

#### install go
curl https://go.dev/dl/go1.21.3.linux-arm64.tar.gz -o go.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz
export PATH=$PATH:/usr/local/go/bin
export CGO_ENABLED=1

#### install gcc
sudo apt install build-essential

#### install gh
type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
&& sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
&& echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
&& sudo apt update \
&& sudo apt install gh -y
export GITHUB_TOKEN=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