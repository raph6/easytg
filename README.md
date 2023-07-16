# Easy Telegram for Golang

## How to use
```go
import "github.com/raph6/easytg"

func main() {
    bot := easytg.NewBot("xxxx_YOUR_API_KEY_xxxx")
	bot.Msg("test", 99999999) // conversation id
	err := bot.Send()
	if err != nil {
		...
	}
}
```