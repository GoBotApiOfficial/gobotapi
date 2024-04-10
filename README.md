<p align="center">
    <br><br>
    <b>AutoGenerated Telegram BotAPI Framework for GoLang</b>
    <br>
    <a href="https://github.com/GoBotApiOfficial/gobotapi/tree/main/examples">
        Examples
    </a>
    •
    <a href="https://pkg.go.dev/github.com/GoBotApiOfficial/gobotapi#section-documentation">
        Documentation
    </a>
    •
    <a href="https://github.com/GoBotApiOfficial/gobotapi">
        Sources
    </a>
</p>

# GoBotAPI
[![pkg.go.dev](https://img.shields.io/github/v/tag/GoBotApiOfficial/gobotapi?color=00b7e0&label=go.pkg.dev)](https://pkg.go.dev/github.com/GoBotApiOfficial/gobotapi)
[![GO Version](https://img.shields.io/github/go-mod/go-version/GoBotApiOfficial/gobotapi)](https://go.dev/)
[![GitHub](https://img.shields.io/github/license/GoBotApiOfficial/gobotapi)](https://github.com/GoBotApiOfficial/gobotapi/blob/main/LICENSE.md)
![OS](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey)
![Architectures](https://img.shields.io/badge/architectures-x86__64%20%7C%20arm64v8%20%7C%20win__amd64%20%7C%20darwin__x64-blue)

> An elegant and modern BotAPI Framework for GoLang

``` go
package main

import "github.com/GoBotApiOfficial/gobotapi"
import "github.com/GoBotApiOfficial/gobotapi/types"
import "github.com/GoBotApiOfficial/gobotapi/methods"

func main() {
    client := gobotapi.NewClient("YOUR_TOKEN")
    client.OnMessage(func(ctx *gobotapi.Client, message types.Message) {
        ctx.Invoke(&methods.SendMessage{
            ChatID: message.Chat.ID,
            Text:   "Hello World!",
        })
    })
    client.Run()
}
```

**GoBotAPI** is a modern and elegant AutoGenerated [BotAPI](https://core.telegram.org/bots/api) Framework. This Framework provides a pure
Go implementation **without any external libs**

In addition to the official API, this Framework also provides some **high-level
functions** that make it easier to use the API.

> The Telegram API scheme depends on your build, but if you don't want to compile by your self you
can use the package compiled from [pkg.go.dev/github.com/GoBotApiOfficial/gobotapi](https://pkg.go.dev/github.com/GoBotApiOfficial/gobotapi).

## How to install?
Here's how to add the GoBotApi Framework to your project, the command are given below:
``` bash
go get -u github.com/GoBotApiOfficial/gobotapi
```

# GoBotAPI Generator
[![GO Version](https://img.shields.io/github/go-mod/go-version/GoBotApiOfficial/GoBotApiGenerator)](https://go.dev/)
[![GitHub](https://img.shields.io/github/license/GoBotApiOfficial/GoBotApiGenerator)](https://github.com/GoBotApiOfficial/GoBotApiGenerator/blob/main/LICENSE.md)
![OS](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey)
![Architectures](https://img.shields.io/badge/architectures-x86__64%20%7C%20arm64v8%20%7C%20win__amd64%20%7C%20darwin__x64-blue)

**GoBotAPI Generator** is the generator for the GoBotAPI Framework. It generates a new BotAPI Framework from the given
[BotAPI](https://core.telegram.org/bots/api) schema.

# How to use?
Just run the binary file from [releases](https://github.com/GoBotApiOfficial/GoBotApiCompiler/releases)
and choose the binary file you want to use on your operating system.

## Credits
Big thanks to [@Laky-64] for making this project possible, special thanks to [@geiccobs] for his own package as
starting point for this project, also thanks to [@empijei] for help about the project design and to
[@LucaTheHacker] for optimizations.

[@Laky-64]: https://github.com/Laky-64
[@geiccobs]: https://github.com/geiccobs
[@empijei]: https://github.com/empijei
[@LucaTheHacker]: https://github.com/LucaTheHacker
