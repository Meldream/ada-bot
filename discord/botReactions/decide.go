package botReactions

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Meldream/ada-bot/settings"
)

type decide struct {
	Trigger string
}

func (d *decide) Help() string {
	return "Let the bot decide between two or more things for you!"
}

func (d *decide) HelpDetail() string {
	return d.Help()
}

func (d *decide) Reaction(m *discordgo.Message, a *discordgo.Member, mType string) Reaction {
	choices := strings.Split(m.Content[len(settings.Settings.Discord.BotPrefix)+len(d.Trigger):], " or ")
	theAnswer := choices[rand.Intn(len(choices))]
	response := fmt.Sprintf("The correct answer is **%s**", strings.TrimSpace(theAnswer))
	return Reaction{Text: response}
}

func init() {
	rand.Seed(time.Now().Unix())

	_decide := &decide{
		Trigger: "decide",
	}
	addReaction(_decide.Trigger, "CREATE", _decide)
}
