package share_package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

func TestTelegram() {
	token := getEnv("TELEGRAM_TOKEN", "-1")
	chatId := getEnvAsInt("TELEGRAM_CHAT_ID", -1)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(chatId, "test")
	_, err = bot.Send(msg)
	if err != nil {
		return
	}

}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int64) int64 {
	valueStr := getEnv(name, "0")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}

	return defaultVal
}
