package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

// StrGenerateRandomNumber generates a random number from
// a given max value as a string
func StrGenerateRandomNumber(max string) int {
	num, err := strconv.Atoi(max)
	if num < 1 {
		return 0
	}

	if err != nil {
		fmt.Printf("Supplied value %v is not a number", num)
		return 0
	} else {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(num)
	}
}

// GenerateRandomNumber generates a random number from
// a given max value as a int
func GenerateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// ElevatedPrivsMessage is checking a given message twitch.PrivateMessage
// if it came from a moderator/vip/or broadcaster and returns a bool
func ElevatedPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["vip"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}

// ModPrivsMessage is checking a given message twitch.PrivateMessage
// if it came from a moderator or broadcaster and returns a bool
func ModPrivsMessage(message twitch.PrivateMessage) bool {
	if message.User.Badges["moderator"] == 1 ||
		message.User.Badges["broadcaster"] == 1 {
		return true
	} else {
		return false
	}
}
