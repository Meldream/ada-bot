package botReactions

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Meldream/ada-bot/settings"
)

type dice struct {
	Trigger string
}

func (d *dice) Help() string {
	return fmt.Sprintf("Roll a dice! DnD style, %sdice xdy+z", settings.Settings.Discord.BotPrefix)
}

func (d *dice) HelpDetail() string {
	return d.Help()
}

var diceRegexp = regexp.MustCompile(`(?i)([0-9]+)d([0-9]+)(?:\+([0-9]+))?`)
var validSides = [6]int{4, 6, 8, 10, 12, 20}

func (d *dice) Reaction(m *discordgo.Message, a *discordgo.Member, mType string) Reaction {
	var numDice, numSides, addNum, response string
	var roll int

	request := strings.TrimSpace(m.Content[len(settings.Settings.Discord.BotPrefix)+len(d.Trigger):])
	if !(len(request) > 0) {
		request = "1d6"
	}
	diceRoll := ""
	total := 0

	dMatch := diceRegexp.FindStringSubmatch(request)
	if len(dMatch) > 0 {
		numDice, numSides, addNum = dMatch[1], dMatch[2], dMatch[3]
		if _numDice, err := strconv.Atoi(numDice); err == nil {
			if _numDice > 7 {
				response = "But I have small hands, I can't hold that many dice :frowning:"
				return Reaction{Text: response}
			}
			if _numSides, err := strconv.Atoi(numSides); err == nil {
				_validNumSides := 0
				for _, nSide := range validSides {
					if nSide == _numSides {
						_validNumSides = nSide
					}
				}
				if _validNumSides == 0 {
					response = "Wow those are strange dice, I don't even know how to roll 'em :confused:"
					return Reaction{Text: response}
				}
				for dice := 0; dice < _numDice; dice++ {
					if _numSides > 0 {
						roll = rand.Intn(_numSides) + 1
					} else {
						roll = 0
					}
					diceRoll = fmt.Sprintf("%s %d", diceRoll, roll)
					total += roll
				}
			} else {
				log.Printf("error: %v", err) // Non fatal error at strconv.Atoi() call
			}
		} else {
			log.Printf("error: %v", err) // Non fatal error at strconv.Atoi() call
		}

		if len(addNum) > 0 {
			if _addNum, err := strconv.Atoi(addNum); err == nil {
				if _addNum > 99 {
					response = "Are you trying to cheat here? :open_mouth:"
					return Reaction{Text: response}
				}
				total += _addNum
				diceRoll = fmt.Sprintf("%s %d", diceRoll, _addNum)
			} else {
				log.Printf("error: %v", err) // Non fatal error at strconv.Atoi() call
			}
		} else {
			addNum = "0"
		}
	}

	if len(diceRoll) > 0 {
		response = fmt.Sprintf("```Dice roll [%sd%s+%s]: %s\tTotal: %d```", numDice, numSides, addNum, diceRoll, total)
	}
	return Reaction{Text: response}
}

func init() {
	rand.Seed(time.Now().Unix())

	_dice := &dice{
		Trigger: "dice",
	}
	addReaction(_dice.Trigger, "CREATE", _dice)
}
