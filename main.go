package main

import (
	"fmt"
	"os"
	"strings"
)

// Rune represents an Elder Futhark rune
type Rune struct {
	Symbol        string
	Name          string
	Pronunciation string
	Meaning       string
	Keywords      []string
}

// Aett represents a group of 8 runes
type Aett struct {
	Name  string
	Runes []Rune
}

var (
	// First Aett - Freya's Aett
	freyaAett = Aett{
		Name: "Freya's Aett (Fertility and Creation)",
		Runes: []Rune{
			{
				Symbol:        "ᚠ",
				Name:          "Fehu",
				Pronunciation: "FAY-hoo",
				Meaning:       "Wealth, Prosperity, Abundance, Cattle, New Beginnings",
				Keywords:      []string{"wealth", "prosperity", "abundance", "cattle", "new beginnings", "money", "success", "fortune"},
			},
			{
				Symbol:        "ᚢ",
				Name:          "Uruz",
				Pronunciation: "OO-rooz",
				Meaning:       "Strength, Vitality, Physical Health, Primal Power, Wild Ox",
				Keywords:      []string{"strength", "vitality", "physical health", "primal power", "wild ox", "energy", "endurance"},
			},
			{
				Symbol:        "ᚦ",
				Name:          "Thurisaz",
				Pronunciation: "THOOR-ee-saz",
				Meaning:       "Protection, Defense, Giants, Thorns, Reactive Force",
				Keywords:      []string{"protection", "defense", "giants", "thorns", "reactive force", "thorn", "challenge"},
			},
			{
				Symbol:        "ᚨ",
				Name:          "Ansuz",
				Pronunciation: "AHN-sooz",
				Meaning:       "Communication, Wisdom, Divine Inspiration, Signals, Odin",
				Keywords:      []string{"communication", "wisdom", "divine inspiration", "signals", "odin", "speech", "god"},
			},
			{
				Symbol:        "ᚱ",
				Name:          "Raidho",
				Pronunciation: "RYE-though",
				Meaning:       "Journey, Travel, Movement, Rhythm, Right Action",
				Keywords:      []string{"journey", "travel", "movement", "rhythm", "right action", "progress", "path"},
			},
			{
				Symbol:        "ᚲ",
				Name:          "Kenaz",
				Pronunciation: "KEN-az",
				Meaning:       "Knowledge, Torch, Light, Creativity, Transformation",
				Keywords:      []string{"knowledge", "torch", "light", "creativity", "transformation", "enlightenment", "fire"},
			},
			{
				Symbol:        "ᚷ",
				Name:          "Gebo",
				Pronunciation: "GEH-bow",
				Meaning:       "Gift, Generosity, Partnership, Love, Balance",
				Keywords:      []string{"gift", "generosity", "partnership", "love", "balance", "exchange", "relationship"},
			},
			{
				Symbol:        "ᚹ",
				Name:          "Wunjo",
				Pronunciation: "WOO-nyo",
				Meaning:       "Joy, Happiness, Harmony, Peace, Fellowship",
				Keywords:      []string{"joy", "happiness", "harmony", "peace", "fellowship", "contentment", "bliss"},
			},
		},
	}

	// Second Aett - Heimdall's Aett
	heimdallAett = Aett{
		Name: "Heimdall's Aett (Protection and Challenges)",
		Runes: []Rune{
			{
				Symbol:        "ᚺ",
				Name:          "Hagalaz",
				Pronunciation: "HAH-gah-laz",
				Meaning:       "Hail, Disruption, Natural Forces, Change, Uncontrolled",
				Keywords:      []string{"hail", "disruption", "natural forces", "change", "uncontrolled", "crisis", "storm"},
			},
			{
				Symbol:        "ᚾ",
				Name:          "Nauthiz",
				Pronunciation: "NOW-theez",
				Meaning:       "Need, Necessity, Resistance, Constraint, Endurance",
				Keywords:      []string{"need", "necessity", "resistance", "constraint", "endurance", "hardship", "survival"},
			},
			{
				Symbol:        "ᛁ",
				Name:          "Isa",
				Pronunciation: "EE-sah",
				Meaning:       "Ice, Stillness, Stagnation, Clarity, Focus",
				Keywords:      []string{"ice", "stillness", "stagnation", "clarity", "focus", "patience", "frozen"},
			},
			{
				Symbol:        "ᛃ",
				Name:          "Jera",
				Pronunciation: "YEH-rah",
				Meaning:       "Year, Harvest, Cycles, Rewards, Natural Timing",
				Keywords:      []string{"year", "harvest", "cycles", "rewards", "natural timing", "season", "patience"},
			},
			{
				Symbol:        "ᛇ",
				Name:          "Eihwaz",
				Pronunciation: "AY-waz",
				Meaning:       "Yew Tree, Defense, Endurance, Reliability, Strength",
				Keywords:      []string{"yew tree", "defense", "endurance", "reliability", "strength", "protection", "rebirth"},
			},
			{
				Symbol:        "ᛈ",
				Name:          "Perthro",
				Pronunciation: "PER-throw",
				Meaning:       "Mystery, Fate, Secrets, Hidden Knowledge, Chance",
				Keywords:      []string{"mystery", "fate", "secrets", "hidden knowledge", "chance", "divination", "luck"},
			},
			{
				Symbol:        "ᛉ",
				Name:          "Algiz",
				Pronunciation: "AL-geez",
				Meaning:       "Protection, Shield, Higher Self, Connection to Divine",
				Keywords:      []string{"protection", "shield", "higher self", "connection to divine", "guardian", "safety"},
			},
			{
				Symbol:        "ᛋ",
				Name:          "Sowilo",
				Pronunciation: "SO-wee-lo",
				Meaning:       "Sun, Success, Wholeness, Life Force, Victory",
				Keywords:      []string{"sun", "success", "wholeness", "life force", "victory", "achievement", "light"},
			},
		},
	}

	// Third Aett - Tyr's Aett
	tyrAett = Aett{
		Name: "Tyr's Aett (Spirituality and Community)",
		Runes: []Rune{
			{
				Symbol:        "ᛏ",
				Name:          "Tiwaz",
				Pronunciation: "TEE-waz",
				Meaning:       "Warrior, Justice, Honor, Victory, Tyr the God",
				Keywords:      []string{"warrior", "justice", "honor", "victory", "tyr", "courage", "sacrifice"},
			},
			{
				Symbol:        "ᛒ",
				Name:          "Berkano",
				Pronunciation: "BER-kah-no",
				Meaning:       "Birch, Growth, Renewal, Fertility, New Beginnings",
				Keywords:      []string{"birch", "growth", "renewal", "fertility", "new beginnings", "birth", "nurture"},
			},
			{
				Symbol:        "ᛖ",
				Name:          "Ehwaz",
				Pronunciation: "EH-waz",
				Meaning:       "Horse, Movement, Progress, Trust, Partnership",
				Keywords:      []string{"horse", "movement", "progress", "trust", "partnership", "cooperation", "teamwork"},
			},
			{
				Symbol:        "ᛗ",
				Name:          "Mannaz",
				Pronunciation: "MAH-naz",
				Meaning:       "Humanity, Self, Community, Intelligence, Rationality",
				Keywords:      []string{"humanity", "self", "community", "intelligence", "rationality", "mind", "social"},
			},
			{
				Symbol:        "ᛚ",
				Name:          "Laguz",
				Pronunciation: "LAH-gooz",
				Meaning:       "Water, Flow, Emotions, Intuition, Dreams",
				Keywords:      []string{"water", "flow", "emotions", "intuition", "dreams", "psychic", "unconscious"},
			},
			{
				Symbol:        "ᛜ",
				Name:          "Ingwaz",
				Pronunciation: "ING-waz",
				Meaning:       "Fertility, Potential, Internal Growth, Ing the God",
				Keywords:      []string{"fertility", "potential", "internal growth", "ing", "peace", "gestation", "family"},
			},
			{
				Symbol:        "ᛞ",
				Name:          "Dagaz",
				Pronunciation: "DAH-gaz",
				Meaning:       "Day, Dawn, Breakthrough, Awakening, Transformation",
				Keywords:      []string{"day", "dawn", "breakthrough", "awakening", "transformation", "clarity", "hope"},
			},
			{
				Symbol:        "ᛟ",
				Name:          "Othala",
				Pronunciation: "OH-thah-lah",
				Meaning:       "Heritage, Home, Ancestral Property, Legacy, Inheritance",
				Keywords:      []string{"heritage", "home", "ancestral property", "legacy", "inheritance", "tradition", "family"},
			},
		},
	}

	allAetts = []Aett{freyaAett, heimdallAett, tyrAett}
)

// getAllRunes returns all runes in a flat list
func getAllRunes() []Rune {
	var runes []Rune
	for _, aett := range allAetts {
		runes = append(runes, aett.Runes...)
	}
	return runes
}

// displayAllRunes shows all three aetts with their runes
func displayAllRunes() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║           ELDER FUTHARK RUNES - The Old Norse Alphabet            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝\n")

	for _, aett := range allAetts {
		fmt.Printf("┌─ %s\n", aett.Name)
		fmt.Println("│")
		for _, rune := range aett.Runes {
			fmt.Printf("│  %s  %-12s (%s)\n", rune.Symbol, rune.Name, rune.Pronunciation)
			fmt.Printf("│      %s\n", rune.Meaning)
		}
		fmt.Println("│")
	}
	fmt.Println("└────────────────────────────────────────────────────────────────────\n")
	fmt.Println("Usage: runes [keyword|rune-name|symbol]")
	fmt.Println("Examples: runes love")
	fmt.Println("          runes fehu")
	fmt.Println("          runes ᚠ")
	fmt.Println()
}

// searchRunes searches for runes matching the given keyword, rune name, or symbol
func searchRunes(keyword string) {
	keyword = strings.TrimSpace(keyword)
	keywordLower := strings.ToLower(keyword)
	var matches []Rune

	for _, rune := range getAllRunes() {
		// Check if keyword matches the rune symbol (exact match)
		if rune.Symbol == keyword {
			matches = append(matches, rune)
			continue
		}

		// Check if keyword matches the rune name
		if strings.EqualFold(rune.Name, keyword) || strings.Contains(strings.ToLower(rune.Name), keywordLower) {
			matches = append(matches, rune)
			continue
		}

		// Check if keyword matches any of the rune's keywords
		for _, kw := range rune.Keywords {
			if strings.Contains(kw, keywordLower) || strings.Contains(keywordLower, kw) {
				matches = append(matches, rune)
				break
			}
		}
	}

	if len(matches) == 0 {
		fmt.Printf("\nNo runes found matching '%s'\n", keyword)
		fmt.Println("\nTry keywords like: love, wealth, strength, protection, journey, wisdom, etc.")
		fmt.Println("Or search by rune name: fehu, uruz, thurisaz, ansuz, etc.")
		fmt.Println("Or search by symbol: ᚠ, ᚢ, ᚦ, etc.")
		fmt.Println("Or run 'runes' without arguments to see all runes.\n")
		return
	}

	fmt.Printf("\n╔════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("\n║  Runes matching '%s':", keyword)
	for i := len(keyword); i < 49; i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝\n")

	for _, rune := range matches {
		fmt.Printf("  %s  %s (%s)\n", rune.Symbol, rune.Name, rune.Pronunciation)
		fmt.Printf("      Meaning: %s\n", rune.Meaning)
		fmt.Printf("      Related: %s\n\n", strings.Join(rune.Keywords, ", "))
	}
}

func main() {
	if len(os.Args) < 2 {
		// No arguments - display all runes
		displayAllRunes()
		return
	}

	// Search for runes by keyword
	keyword := strings.Join(os.Args[1:], " ")
	searchRunes(keyword)
}
