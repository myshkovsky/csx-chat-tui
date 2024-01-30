# csx-chat-tui
Basic Twitch chat TUI, written in Go.

Inspired by [LCOLONQ](https://twitch.tv/LCOLONQ)'s fig-chat, featured during his streams.

I used this project as training grounds while learning Go. It is still missing some features, but the basic functionality is in place.

It is not a priority of mine, so I thought I'd publish it in case someone finds it interesting and/or wants to continue developing it.

## Feature list
- [x] Connect to any Twitch chat (set in config.yaml)
- [x] Change badge emojis for Broadcaster, Subscriber, Moderator, etc. (set in config.yaml)
- [x] Receive messages from the chat in real time
- [x] Display unicode emojis (if your terminal supports it)
- [ ] Display Twitch global & channel emotes
- [ ] Display 7TV emotes
- [ ] Send messages directly from the TUI
- [ ] Create a template config if no config file is found

## Usage

### Release
Download a release zip, unpack it, set up your config and run the executable.

### Build from source
Download source, set up your config and run directly with `go run main.go` or build your exe with `go build`.

As of right now, the config file is not automatically created by the executable, so you'll have to make your own or copy the one from source. Latest release always includes a template config.

## Notes
I can't wrap my head around displaying the emotes correctly in the terminal. So far, all of my attempts failed miserably. The furthest I've gotten is downloading the emotes and caching them, so the Twitch API doesn't get spammed by the client. It also saves bandwidth.

If you know how to implement any of the missing features, feel free to contribute.

## License
No warranty, no expectations. You may do whatever you want with the source code.
Credit is greatly appreciated but not required.
