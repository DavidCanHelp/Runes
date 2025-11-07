# Elder Futhark Runes CLI

A command-line tool for working with the 24 Elder Futhark runes, organized by their traditional three aetts.

## Installation

```bash
go build -o runes
```

## Usage

### Display All Runes

Simply run the command without arguments to see all 24 runes organized by their three aetts:

```bash
./runes
```

This displays:
- **Freya's Aett** (Fertility and Creation) - 8 runes
- **Heimdall's Aett** (Protection and Challenges) - 8 runes
- **Tyr's Aett** (Spirituality and Community) - 8 runes

### Search by Keyword

Search for runes by meaning or associated concepts:

```bash
./runes love
./runes wealth
./runes protection
./runes journey
```

The search is flexible and matches partial keywords. For example:
- `runes protec` will find runes related to "protection"
- `runes fertil` will find runes related to "fertility"

### Search by Rune Name

You can also search for a specific rune by its name:

```bash
./runes fehu
./runes algiz
./runes dagaz
./runes wunjo
```

Name searches are case-insensitive and support partial matching.

### Search by Symbol

You can paste a UTF-8 runic symbol directly to look up its name and meaning:

```bash
./runes ᚠ
./runes ᛉ
./runes ᛞ
./runes ᚹ
```

This is useful when you encounter a rune symbol and want to identify it.

## The 24 Elder Futhark Runes

### Freya's Aett
- ᚠ **Fehu** - Wealth, Prosperity
- ᚢ **Uruz** - Strength, Vitality
- ᚦ **Thurisaz** - Chaos, Defense
- ᚨ **Ansuz** - Communication, Wisdom
- ᚱ **Raidho** - Journey, Travel
- ᚲ **Kenaz** - Knowledge, Enlightenment
- ᚷ **Gebo** - Gift, Partnership, Love
- ᚹ **Wunjo** - Joy, Harmony

### Heimdall's Aett
- ᚺ **Hagalaz** - Disruption, Change
- ᚾ **Nauthiz** - Need, Necessity
- ᛁ **Isa** - Ice, Stillness
- ᛃ **Jera** - Harvest, Cycles
- ᛇ **Eihwaz** - Endurance, Defense
- ᛈ **Perthro** - Fate, Mystery
- ᛉ **Algiz** - Protection, Sanctuary
- ᛊ **Sowilo** - Sun, Success

### Tyr's Aett
- ᛏ **Tiwaz** - Justice, Honor
- ᛒ **Berkano** - Growth, Fertility
- ᛖ **Ehwaz** - Movement, Partnership
- ᛗ **Mannaz** - Humanity, Self
- ᛚ **Laguz** - Water, Intuition
- ᛜ **Ingwaz** - Fertility, Completion
- ᛞ **Dagaz** - Awakening, Breakthrough
- ᛟ **Othala** - Heritage, Home

## Example Searches

Common keywords you can search for:
- Wealth: `runes wealth`, `runes money`, `runes prosperity`
- Love: `runes love`, `runes relationship`
- Protection: `runes protection`, `runes defense`, `runes shield`
- Wisdom: `runes wisdom`, `runes knowledge`, `runes learning`
- Journey: `runes journey`, `runes travel`, `runes path`
- Success: `runes success`, `runes victory`, `runes achievement`

## About Elder Futhark

The Elder Futhark is the oldest form of the runic alphabets, used from around 150-800 CE. The 24 runes are divided into three groups of eight, called "aetts" (families), each associated with different Norse deities and aspects of life.
