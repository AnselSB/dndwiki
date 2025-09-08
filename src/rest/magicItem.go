package rest

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MagicItem struct {
	Name     string            `json:"name"`
	EquipCat EquipmentCategory `json:"equipment_category"`
	Rarity   Rarity            `json:"rarity"`
	Variants []Variant         `json:"variants"`
	Desc     []string          `json:"desc"`
}

func GetMagicItem(itemName string) (*discordgo.MessageEmbed, error) {
	// make the endpoint url
	endpointURL := fmt.Sprintf("magic-items/%v", itemName)
	// get the response
	itemResponse, err := makeRequest(endpointURL)
	if err != nil {
		return nil, err
	}

	// unmarshal the information into the struct we defined
	var itemEntity MagicItem
	err = json.Unmarshal(itemResponse, &itemEntity)
	if err != nil {
		return nil, err
	}
	// we now have to start formatting shit so that we can send things through as an embed
	fullDesc := formatMultiValues(itemEntity.Desc)

	var fullVariants string
	if len(itemEntity.Variants) == 0 {
		fullVariants = "N/A"
	} else {
		fullVariants = formatVariants(itemEntity.Variants)
	}

	embed := &discordgo.MessageEmbed{
		Title:       itemEntity.Name,
		URL:         "",
		Description: fullDesc,
		Color:       BotColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Equipment Category",
				Value:  itemEntity.EquipCat.Name,
				Inline: false,
			},
			{
				Name:   "Rarity",
				Value:  itemEntity.Rarity.Name,
				Inline: false,
			}, {
				Name:   "Variants",
				Value:  fullVariants,
				Inline: false,
			},
		},
	}

	return embed, nil
}

func formatVariants(values []Variant) string {
	var builder strings.Builder
	for index, value := range values {
		builder.WriteString(value.Name)
		if index < len(values)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
