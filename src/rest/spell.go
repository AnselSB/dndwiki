package rest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Spell struct {
	HigherLevel   []string    `json:"higher_level"`
	Name          string      `json:"name"`
	Desc          []string    `json:"desc"`
	Range         string      `json:"range"`
	Components    []string    `json:"components"`
	Ritual        bool        `json:"ritual"`
	Duration      string      `json:"duration"`
	Concentration bool        `json:"concentration"`
	CastingTime   string      `json:"casting_time"`
	Level         int         `json:"level"`
	AttackType    string      `json:"attack_type"`
	SpellDamage   Damage      `json:"damage"`
	School        MagicSchool `json:"school"`
	Classes       []Class     `json:"classes"`
	Material      string      `json:"material"`
}

type MagicSchool struct {
	Name string `json:"name"`
}

type Class struct {
	Name string `json:"name"`
}

func GetSpell(spell string) (*discordgo.MessageEmbed, error) {
	// make the url
	endpointURL := fmt.Sprintf("spells/%v", spell)

	// get the response
	spellResponse, err := makeRequest(endpointURL)
	if err != nil {
		return nil, err
	}

	// now try to unmarshal this information into the struct
	var spellEntity Spell
	err = json.Unmarshal(spellResponse, &spellEntity)
	if err != nil {
		return nil, err
	}
	// this now begins the process of formatting the spell into a string we'll send, hopefully I don't run into character limit issues
	formattedHigherLevel := formatMultiValues(spellEntity.HigherLevel)
	formattedDesc := formatMultiValues(spellEntity.Desc)
	formattedComponents := formatComponents(spellEntity.Components)
	formattedClasses := formatClassValues(spellEntity.Classes)

	if strings.Contains(formattedComponents, "M") {
		formattedComponents = fmt.Sprintf("%v (%v)", formattedComponents, spellEntity.Material)
	}

	fullDesc := formattedDesc + "\n" + formattedHigherLevel

	embed := &discordgo.MessageEmbed{
		Title:       spellEntity.Name,
		URL:         "",
		Description: fullDesc,
		Color:       BotColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Spell Level",
				Value:  strconv.Itoa(spellEntity.Level),
				Inline: false,
			},
			{
				Name:   "School",
				Value:  spellEntity.School.Name,
				Inline: false,
			},
			{
				Name:   "Casting Time",
				Value:  spellEntity.CastingTime,
				Inline: false,
			},
			{
				Name:   "Range",
				Value:  spellEntity.Range,
				Inline: false,
			},
			{
				Name:   "Components",
				Value:  formattedComponents,
				Inline: false,
			},
			{
				Name:   "Compatible Classes",
				Value:  formattedClasses,
				Inline: false,
			},
			{
				Name:   "Ritual",
				Value:  strconv.FormatBool(spellEntity.Ritual),
				Inline: false,
			},
		},
	}

	return embed, nil
}

func formatClassValues(values []Class) string {
	var builder strings.Builder
	for index, value := range values {
		builder.WriteString(value.Name)
		if index < len(values)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

func formatComponents(values []string) string {
	var builder strings.Builder
	for index, value := range values {
		builder.WriteString(value)
		if index < len(values)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
