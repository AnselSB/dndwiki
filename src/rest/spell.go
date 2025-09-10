package rest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"wikiBot/src/cache"

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
			{
				Name:   "Concentration",
				Value:  strconv.FormatBool(spellEntity.Concentration),
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

func AddSpell(spellEmbed *discordgo.MessageEmbed) error {
	// we know that this thing is a spell so what we gotta do is set up all the variables we need to define the spell
	var spellFileName string
	higherLevel := make([]string, 0)
	var spellName string
	desc := make([]string, 0)
	var spellRange string
	components := make([]string, 0)
	var isRitual bool = false
	var spellDuration string
	var isConcentration bool = false
	var castingTime string
	var level int
	// attacktype and damage we'll set as constant values for now
	spellAttack := ""
	spellDamage := Damage{DamageType: DamageType{Name: "damage"}, DamageAtCharacterLevel: DamageAtCharacterLevel{First: "1", Fifth: "5", Eleventh: "11", Seventeenth: "17"}}
	spellSchool := MagicSchool{Name: "placeholder"}
	spellClasses := make([]Class, 0)
	var spellMaterial string

	// now we loop through each field and set the according variable
	for _, field := range spellEmbed.Fields {
		// TODO: Fix this shit, make it into a look up table or something

		if field.Name == "Spell Name" {
			spellFileName = field.Value
			spellName = strings.ReplaceAll(field.Value, "-", " ")
		} else if field.Name == "Description" {
			desc = append(desc, field.Value)
		} else if field.Name == "Higher Level" {
			higherLevel = append(higherLevel, field.Value)
		} else if field.Name == "Range" {
			spellRange = field.Value
		} else if strings.Contains(field.Name, "Components") {
			components = append(components, field.Value)
		} else if field.Name == "Ritual" && field.Value == "yes" {
			isRitual = true
		} else if field.Name == "Concentration" && field.Value == "yes" {
			isConcentration = true
		} else if field.Name == "Casting time" {
			castingTime = field.Value
		} else if field.Name == "Level" {
			parsedLevel, err := strconv.ParseInt(field.Value, 10, 64)
			level = int(parsedLevel)
			if err != nil {
				fmt.Println("Error adding spell. Unable to parse spell level")
				return err
			}
		} else if field.Name == "School" {
			spellSchool.Name = field.Value
		} else if strings.Contains(field.Name, "Classes") {
			spellClasses = append(spellClasses, Class{Name: field.Value})
		} else if field.Name == "Material" {
			spellMaterial = field.Value
		} else if field.Name == "Duration" {
			spellDuration = field.Value
		}
	}
	// once all properties are set we generate the spell instance and encode into json
	newSpellEntity := Spell{
		HigherLevel:   higherLevel,
		Name:          spellName,
		Desc:          desc,
		Range:         spellRange,
		Components:    components,
		Ritual:        isRitual,
		Duration:      spellDuration,
		Concentration: isConcentration,
		CastingTime:   castingTime,
		Level:         level,
		AttackType:    spellAttack,
		SpellDamage:   spellDamage,
		School:        spellSchool,
		Classes:       spellClasses,
		Material:      spellMaterial,
	}
	// now that the entity has been made we need to encode it into json
	encodedSpell, err := json.Marshal(newSpellEntity)
	if err != nil {
		return err
	}
	err = cache.SetObject(fmt.Sprintf("spells/%v", spellFileName), encodedSpell)
	if err != nil {
		return err
	}
	return nil
}
