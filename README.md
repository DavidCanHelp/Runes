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
- ᚠ **Fehu** (FAY-hoo) - Wealth, Prosperity, Abundance, Cattle, New Beginnings
- ᚢ **Uruz** (OO-rooz) - Strength, Vitality, Physical Health, Primal Power, Wild Ox
- ᚦ **Thurisaz** (THOOR-ee-saz) - Protection, Defense, Giants, Thorns, Reactive Force
- ᚨ **Ansuz** (AHN-sooz) - Communication, Wisdom, Divine Inspiration, Signals, Odin
- ᚱ **Raidho** (RYE-though) - Journey, Travel, Movement, Rhythm, Right Action
- ᚲ **Kenaz** (KEN-az) - Knowledge, Torch, Light, Creativity, Transformation
- ᚷ **Gebo** (GEH-bow) - Gift, Generosity, Partnership, Love, Balance
- ᚹ **Wunjo** (WOO-nyo) - Joy, Happiness, Harmony, Peace, Fellowship

### Heimdall's Aett
- ᚺ **Hagalaz** (HAH-gah-laz) - Hail, Disruption, Natural Forces, Change, Uncontrolled
- ᚾ **Nauthiz** (NOW-theez) - Need, Necessity, Resistance, Constraint, Endurance
- ᛁ **Isa** (EE-sah) - Ice, Stillness, Stagnation, Clarity, Focus
- ᛃ **Jera** (YEH-rah) - Year, Harvest, Cycles, Rewards, Natural Timing
- ᛇ **Eihwaz** (AY-waz) - Yew Tree, Defense, Endurance, Reliability, Strength
- ᛈ **Perthro** (PER-throw) - Mystery, Fate, Secrets, Hidden Knowledge, Chance
- ᛉ **Algiz** (AL-geez) - Protection, Shield, Higher Self, Connection to Divine
- ᛋ **Sowilo** (SO-wee-lo) - Sun, Success, Wholeness, Life Force, Victory

### Tyr's Aett
- ᛏ **Tiwaz** (TEE-waz) - Warrior, Justice, Honor, Victory, Tyr the God
- ᛒ **Berkano** (BER-kah-no) - Birch, Growth, Renewal, Fertility, New Beginnings
- ᛖ **Ehwaz** (EH-waz) - Horse, Movement, Progress, Trust, Partnership
- ᛗ **Mannaz** (MAH-naz) - Humanity, Self, Community, Intelligence, Rationality
- ᛚ **Laguz** (LAH-gooz) - Water, Flow, Emotions, Intuition, Dreams
- ᛜ **Ingwaz** (ING-waz) - Fertility, Potential, Internal Growth, Ing the God
- ᛞ **Dagaz** (DAH-gaz) - Day, Dawn, Breakthrough, Awakening, Transformation
- ᛟ **Othala** (OH-thah-lah) - Heritage, Home, Ancestral Property, Legacy, Inheritance

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
