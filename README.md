# Laraquest-Go
Sending requests and receiving Telegram updates.
- Bot API Version `8.2`

## Other Versions
- [Laraquest PHP](https://github.com/laraXgram/Laraquest)
- Laraquest Python
- Laraquest JavaScript
- Laraquest Rust
- Laraquest C#

---

# Installation :
```shell
go get github.com/laraXgram/Laraquest-Go
```

# Usage :

### Imports :
```go
import (
    "github.com/laraXgram/Laraquest-Go/request"
    "github.com/laraXgram/Laraquest-Go/updates"
    "github.com/laraXgram/Laraquest-Go/params"
)
```

### Config :
```go
request.AppConfig = request.Config{
    Token: "XXXXXX:XXXXXX",
    API_Server: "https://api.telegram.org", // optional
    Server_IP: "x.x.x.x", // optional
    Server_Port: 9000, // optional
}
```

### Receiving updates :

All updates sent by Telegram, similar to the Telegram structure, are available through the Update struct and will be highlighted by the IDE or editor.

```go
go request.Start(func(update updates.Updates) {
    // Code ...
})
```

### Sending request :

All Telegram methods come with a struct of received parameters, and you can easily use them.

- With response

If the second parameter is true, the bot will wait for a response from Telegram.
```go
params := params.SendMessageParams{
    Chat_id: update.Message.Chat.Id,
    Text: update.Message.Text,
}

request.SendMessage(params, true)
```

- Without response

If the second parameter is false, the bot will not wait for a response and will immediately continue executing the code. This method is very fast, but it will not return the response from Telegram.
```go
params := params.SendMessageParams{
    Chat_id: update.Message.Chat.Id,
    Text: update.Message.Text,
}

request.SendMessage(params, false)
```

### Start Listen :
```go
request.Serve()
```

This method will run a web server, and you can set the webhook on it.
You can configure the `Server_IP` and `Server_Port` as needed.

### Overview :
```go
package main

import (
    "github.com/laraXgram/Laraquest-Go/request"
    "github.com/laraXgram/Laraquest-Go/updates"
    "github.com/laraXgram/Laraquest-Go/params"
)

func main() {
    request.AppConfig = request.Config{
        Token:       "XXX:XXX",
        Server_IP:   "https://mydomain.com/mybot"		
    }

    go request.Start(func(update updates.Updates) {
        if update.Message.Chat.Type == "private" {
            params := params.SendMessageParams{
                Chat_id: update.Message.Chat.Id,
                Text:    "Hello, World!",
            }

            request.SendMessage(params, false)
        }
    })    

    request.Serve()
}
```

### ToDo :
- [x] Get Updates
- [x] Request Methods
- [x] Webhook Handler
- [x] Non-Response Method
- [ ] Long Polling
- [ ] Event Listeners
- [ ] Keyboard Builder
- [ ] Database Manager

### Contribution :

You can complete the project's To-Do list or share your ideas through a PR.
Please ensure your code is standard and follows Laraquest guidelines.

### Contact :
[Telegram](https://t.me/Amirh_krgr) - [Email](mailto:laraxgram@gmail.com)

### Version :
**0.1.1 Alfa**
