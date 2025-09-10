package bot

import "github.com/bwmarrin/discordgo"

var (
	// treat this as a constant, for some reason go won't let me make it const
	minValue   = 0
	minRitual  = 1
	minCon     = 1
	minSchool  = 1
	minClasses = 1
)

func spellModal(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		CustomID: "spell_modal_" + i.Interaction.Member.User.ID,
		Title:    "Add Spell",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "name",
						Label:       "Give the spell a name",
						Style:       discordgo.TextInputShort,
						Placeholder: "spell name goes here",
						Required:    true,
						MaxLength:   100,
						MinLength:   1,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "description",
						Label:       "Describe the spell ",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Description",
						Required:    true,
						MaxLength:   3000,
						MinLength:   100,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "higher level",
						Label:       "describe it",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Description",
						Required:    true,
						MaxLength:   1000,
						MinLength:   50,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "range",
						Label:       "Range in feet",
						Style:       discordgo.TextInputShort,
						Placeholder: "range",
						Required:    true,
						MaxLength:   50,
						MinLength:   1,
					},
					// discordgo.SelectMenu{
					// 	CustomID:    "components",
					// 	Placeholder: "no selection",
					// 	MaxValues:   3,
					// 	MinValues:   &minValue,
					// 	Options: []discordgo.SelectMenuOption{
					// 		{
					// 			Label:       "V",
					// 			Description: "Verbal spell",
					// 			Value:       "V",
					// 			Default:     false,
					// 			Emoji: &discordgo.ComponentEmoji{
					// 				Name: "🗣️",
					// 			},
					// 		},
					// 		{
					// 			Label:       "S",
					// 			Description: "Somatic spell",
					// 			Value:       "S",
					// 			Default:     false,
					// 			Emoji: &discordgo.ComponentEmoji{
					// 				Name: "👏",
					// 			},
					// 		},
					// 		{
					// 			Label:       "M",
					// 			Description: "Material spell",
					// 			Value:       "M",
					// 			Default:     false,
					// 			Emoji: &discordgo.ComponentEmoji{
					// 				Name: "⚗️",
					// 			},
					// 		},
					// 	},
					// },
					// discordgo.SelectMenu{
					// 	CustomID:    "ritual",
					// 	Placeholder: "false",
					// 	MaxValues:   1,
					// 	MinValues:   &minRitual,
					// 	Options: []discordgo.SelectMenuOption{
					// 		{
					// 			Label:       "true",
					// 			Description: "This is a ritual spell",
					// 			Value:       "true",
					// 			Default:     false,
					// 			Emoji: &discordgo.ComponentEmoji{
					// 				Name: "🪄",
					// 			},
					// 		},
					// 		{
					// 			Label:       "false",
					// 			Description: "This is not a ritual spell",
					// 			Value:       "false",
					// 			Default:     true,
					// 			Emoji: &discordgo.ComponentEmoji{
					// 				Name: "🚫",
					// 			},
					// 		},
					// 	},
					// },
				},
			},
			// discordgo.ActionsRow{
			// 	Components: []discordgo.MessageComponent{
			// 		// discordgo.SelectMenu{
			// 		// 	CustomID:    "concentration",
			// 		// 	Placeholder: "does this spell require concentration",
			// 		// 	MaxValues:   1,
			// 		// 	MinValues:   &minCon,
			// 		// 	Options: []discordgo.SelectMenuOption{
			// 		// 		{
			// 		// 			Label:       "true",
			// 		// 			Description: "This is a concentration spell",
			// 		// 			Value:       "true",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "👁️",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "false",
			// 		// 			Description: "This is not a concentration spell",
			// 		// 			Value:       "false",
			// 		// 			Default:     true,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🚫",
			// 		// 			},
			// 		// 		},
			// 		// 	},
			// 		// },
			// 		discordgo.TextInput{
			// 			CustomID:    "casting time",
			// 			Label:       "provide casting",
			// 			Style:       discordgo.TextInputShort,
			// 			Placeholder: "casting time",
			// 			Required:    true,
			// 			MaxLength:   15,
			// 			MinLength:   3,
			// 		},
			// 		// discordgo.TextInput{
			// 		// 	CustomID:    "Level",
			// 		// 	Label:       "provide level",
			// 		// 	Style:       discordgo.TextInputShort,
			// 		// 	Placeholder: "0",
			// 		// 	Required:    true,
			// 		// 	MaxLength:   1,
			// 		// 	MinLength:   1,
			// 		// },
			// 	},
			// },
			// discordgo.ActionsRow{
			// 	Components: []discordgo.MessageComponent{
			// 		// discordgo.SelectMenu{
			// 		// 	CustomID:    "school",
			// 		// 	Placeholder: "Select a school of magic",
			// 		// 	MaxValues:   1,
			// 		// 	MinValues:   &minSchool,
			// 		// 	Options: []discordgo.SelectMenuOption{
			// 		// 		{
			// 		// 			Label:       "abjuration",
			// 		// 			Description: "Emphasizing protection",
			// 		// 			Value:       "abjuration",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🛡️",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "chronurgy",
			// 		// 			Description: "Time magic",
			// 		// 			Value:       "chronurgy",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🕰️",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "conjuration",
			// 		// 			Description: "Magic for creating objects",
			// 		// 			Value:       "conjuration",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🪑",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "divination",
			// 		// 			Description: "Crystal ball shit idk",
			// 		// 			Value:       "divination",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🔮",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "enchantment",
			// 		// 			Description: "enhancing shit",
			// 		// 			Value:       "enchantment",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "✨",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "evocation",
			// 		// 			Description: "shoot stuff (can you tell I'm getting lazy)",
			// 		// 			Value:       "evocation",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "⚡",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "graviturgy",
			// 		// 			Description: "gravity magic",
			// 		// 			Value:       "graviturgy",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🍎",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "illusion",
			// 		// 			Description: "spoopy scary",
			// 		// 			Value:       "illusion",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🔎",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "necromancy",
			// 		// 			Description: "spoopy scary (with skull emoji)",
			// 		// 			Value:       "necromancy",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "💀",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "transmutation",
			// 		// 			Description: "oops it's all primordial soup now",
			// 		// 			Value:       "transmutation",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "💎",
			// 		// 			},
			// 		// 		},
			// 		// 	},
			// 		// },
			// 		// discordgo.SelectMenu{
			// 		// 	CustomID:    "classes",
			// 		// 	Placeholder: "choose what classes can learn this spell",
			// 		// 	MaxValues:   14,
			// 		// 	MinValues:   &minClasses,
			// 		// 	Options: []discordgo.SelectMenuOption{
			// 		// 		{
			// 		// 			Label:       "artificer",
			// 		// 			Description: "they make stuff with their hands",
			// 		// 			Value:       "artificer",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "⚙️",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "barbarian",
			// 		// 			Description: "they break stuff with their hands",
			// 		// 			Value:       "barbarian",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🪓",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "bard",
			// 		// 			Description: "music or something",
			// 		// 			Value:       "bard",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🪈",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "blood hunter",
			// 		// 			Description: "jjk??????",
			// 		// 			Value:       "blood hunter",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🩸",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "cleric",
			// 		// 			Description: "healer (who would've guessed)",
			// 		// 			Value:       "cleric",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🏥",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "druid",
			// 		// 			Description: "naturopath (becky who learned reiki)",
			// 		// 			Value:       "druid",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🌿",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "fighter",
			// 		// 			Description: "human male sword fighter",
			// 		// 			Value:       "fighter",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🗡️",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "monk",
			// 		// 			Description: "bald ahh",
			// 		// 			Value:       "monk",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🦲",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "paladin",
			// 		// 			Description: "they love killing (in the name of god)",
			// 		// 			Value:       "paladin",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🔱",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "ranger",
			// 		// 			Description: "yeah fuck these guys lmao",
			// 		// 			Value:       "ranger",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🧝",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "rogue",
			// 		// 			Description: "tinkle dinkle was a rogue",
			// 		// 			Value:       "rogue",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🔓",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "sorcerer",
			// 		// 			Description: "NOT a wizard",
			// 		// 			Value:       "sorcerer",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "⚡",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "warlock",
			// 		// 			Description: "guy with black and red colour scheme",
			// 		// 			Value:       "warlock",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "😈",
			// 		// 			},
			// 		// 		},
			// 		// 		{
			// 		// 			Label:       "wizard",
			// 		// 			Description: "NOT sorcerer",
			// 		// 			Value:       "wizard",
			// 		// 			Default:     false,
			// 		// 			Emoji: &discordgo.ComponentEmoji{
			// 		// 				Name: "🧙‍♂️",
			// 		// 			},
			// 		// 		},
			// 		// 	},
			// 		// },
			// 		discordgo.TextInput{
			// 			CustomID:    "material",
			// 			Label:       "Describe M",
			// 			Style:       discordgo.TextInputShort,
			// 			Placeholder: "material",
			// 			Required:    true,
			// 			MaxLength:   100,
			// 			MinLength:   0,
			// 		},
			// 	},
			// },
			// discordgo.ActionsRow{
			// 	Components: []discordgo.MessageComponent{
			// 		discordgo.TextInput{
			// 			CustomID:    "duration",
			// 			Label:       "describe duration",
			// 			Style:       discordgo.TextInputShort,
			// 			Placeholder: "duration",
			// 			Required:    true,
			// 			MaxLength:   100,
			// 			MinLength:   1,
			// 		},
			// 	},
			// },
		},
	}
}
