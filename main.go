package main

import (
	"fmt"
	"os"
	"strings"
)

// Rune represents an Elder Futhark rune
type Rune struct {
	Symbol   string
	Name     string
	Meaning  string
	Keywords []string
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
				Symbol:   "ᚠ",
				Name:     "Fehu",
				Meaning:  "Cattle, Wealth, Prosperity",
				Keywords: []string{"wealth", "prosperity", "abundance", "cattle", "money", "success", "fortune"},
			},
			{
				Symbol:   "ᚢ",
				Name:     "Uruz",
				Meaning:  "Aurochs, Strength, Wild Ox",
				Keywords: []string{"strength", "power", "vitality", "health", "energy", "endurance", "ox"},
			},
			{
				Symbol:   "ᚦ",
				Name:     "Thurisaz",
				Meaning:  "Giant, Thorn, Chaos",
				Keywords: []string{"chaos", "destruction", "giant", "thorn", "conflict", "challenge", "defense"},
			},
			{
				Symbol:   "ᚨ",
				Name:     "Ansuz",
				Meaning:  "God, Communication, Divine Inspiration",
				Keywords: []string{"communication", "wisdom", "inspiration", "god", "odin", "speech", "divination"},
			},
			{
				Symbol:   "ᚱ",
				Name:     "Raidho",
				Meaning:  "Journey, Wagon, Travel",
				Keywords: []string{"journey", "travel", "movement", "progress", "wagon", "quest", "path"},
			},
			{
				Symbol:   "ᚲ",
				Name:     "Kenaz",
				Meaning:  "Torch, Knowledge, Enlightenment",
				Keywords: []string{"knowledge", "enlightenment", "torch", "creativity", "learning", "fire", "illumination"},
			},
			{
				Symbol:   "ᚷ",
				Name:     "Gebo",
				Meaning:  "Gift, Partnership, Exchange",
				Keywords: []string{"gift", "partnership", "love", "generosity", "exchange", "relationship", "balance"},
			},
			{
				Symbol:   "ᚹ",
				Name:     "Wunjo",
				Meaning:  "Joy, Pleasure, Harmony",
				Keywords: []string{"joy", "happiness", "harmony", "pleasure", "peace", "contentment", "bliss"},
			},
		},
	}

	// Second Aett - Heimdall's Aett
	heimdallAett = Aett{
		Name: "Heimdall's Aett (Protection and Challenges)",
		Runes: []Rune{
			{
				Symbol:   "ᚺ",
				Name:     "Hagalaz",
				Meaning:  "Hail, Disruption, Uncontrolled Forces",
				Keywords: []string{"hail", "disruption", "destruction", "crisis", "change", "uncontrolled", "storm"},
			},
			{
				Symbol:   "ᚾ",
				Name:     "Nauthiz",
				Meaning:  "Need, Necessity, Constraint",
				Keywords: []string{"need", "necessity", "constraint", "hardship", "resistance", "requirement", "survival"},
			},
			{
				Symbol:   "ᛁ",
				Name:     "Isa",
				Meaning:  "Ice, Stillness, Stagnation",
				Keywords: []string{"ice", "stillness", "stagnation", "standstill", "patience", "frozen", "waiting"},
			},
			{
				Symbol:   "ᛃ",
				Name:     "Jera",
				Meaning:  "Year, Harvest, Cycles",
				Keywords: []string{"harvest", "year", "cycle", "reward", "season", "patience", "growth"},
			},
			{
				Symbol:   "ᛇ",
				Name:     "Eihwaz",
				Meaning:  "Yew Tree, Endurance, Defense",
				Keywords: []string{"endurance", "defense", "yew", "protection", "reliability", "death", "rebirth"},
			},
			{
				Symbol:   "ᛈ",
				Name:     "Perthro",
				Meaning:  "Fate, Mystery, Secrets",
				Keywords: []string{"fate", "mystery", "secrets", "chance", "divination", "hidden", "luck"},
			},
			{
				Symbol:   "ᛉ",
				Name:     "Algiz",
				Meaning:  "Elk, Protection, Sanctuary",
				Keywords: []string{"protection", "defense", "sanctuary", "elk", "shield", "guardian", "safety"},
			},
			{
				Symbol:   "ᛊ",
				Name:     "Sowilo",
				Meaning:  "Sun, Success, Wholeness",
				Keywords: []string{"sun", "success", "wholeness", "victory", "achievement", "light", "power"},
			},
		},
	}

	// Third Aett - Tyr's Aett
	tyrAett = Aett{
		Name: "Tyr's Aett (Spirituality and Community)",
		Runes: []Rune{
			{
				Symbol:   "ᛏ",
				Name:     "Tiwaz",
				Meaning:  "Tyr, Justice, Honor",
				Keywords: []string{"justice", "honor", "courage", "sacrifice", "tyr", "law", "warrior"},
			},
			{
				Symbol:   "ᛒ",
				Name:     "Berkano",
				Meaning:  "Birch, Growth, Fertility",
				Keywords: []string{"growth", "fertility", "birch", "birth", "nurture", "new beginnings", "renewal"},
			},
			{
				Symbol:   "ᛖ",
				Name:     "Ehwaz",
				Meaning:  "Horse, Movement, Partnership",
				Keywords: []string{"horse", "movement", "partnership", "trust", "cooperation", "progress", "teamwork"},
			},
			{
				Symbol:   "ᛗ",
				Name:     "Mannaz",
				Meaning:  "Man, Humanity, Self",
				Keywords: []string{"humanity", "self", "man", "community", "mind", "intelligence", "social"},
			},
			{
				Symbol:   "ᛚ",
				Name:     "Laguz",
				Meaning:  "Water, Flow, Intuition",
				Keywords: []string{"water", "flow", "intuition", "emotion", "psychic", "dreams", "unconscious"},
			},
			{
				Symbol:   "ᛜ",
				Name:     "Ingwaz",
				Meaning:  "Ing, Fertility, Completion",
				Keywords: []string{"fertility", "completion", "ing", "peace", "gestation", "potential", "family"},
			},
			{
				Symbol:   "ᛞ",
				Name:     "Dagaz",
				Meaning:  "Day, Awakening, Breakthrough",
				Keywords: []string{"day", "awakening", "breakthrough", "transformation", "clarity", "hope", "dawn"},
			},
			{
				Symbol:   "ᛟ",
				Name:     "Othala",
				Meaning:  "Ancestral Property, Heritage, Home",
				Keywords: []string{"heritage", "home", "ancestry", "property", "inheritance", "tradition", "family"},
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
			fmt.Printf("│  %s  %-12s  %s\n", rune.Symbol, rune.Name, rune.Meaning)
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
		fmt.Printf("  %s  %s\n", rune.Symbol, rune.Name)
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
