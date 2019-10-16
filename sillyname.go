package sillyname

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var adjectives = []string{"black", "white", "gray", "brown", "red", "pink", "crimson", "carnelian", "orange", "yellow", "ivory", "cream", "green", "viridian", "aquamarine", "cyan", "blue", "cerulean", "azure", "indigo", "navy", "violet", "purple", "lavender", "magenta", "rainbow", "iridescent", "spectrum", "prism", "bold", "vivid", "pale", "clear", "glass", "translucent", "misty", "dark", "light", "gold", "silver", "copper", "bronze", "steel", "iron", "brass", "mercury", "zinc", "chrome", "platinum", "titanium", "nickel", "lead", "pewter", "rust", "metal", "stone", "quartz", "granite", "marble", "alabaster", "agate", "jasper", "pebble", "pyrite", "crystal", "geode", "obsidian", "mica", "flint", "sand", "gravel", "boulder", "basalt", "ruby", "beryl", "scarlet", "citrine", "sulpher", "topaz", "amber", "emerald", "malachite", "jade", "abalone", "lapis", "sapphire", "diamond", "peridot", "gem", "jewel", "bevel", "coral", "jet", "ebony", "wood", "tree", "cherry", "maple", "cedar", "branch", "bramble", "rowan", "ash", "fir", "pine", "cactus", "alder", "grove", "forest", "jungle", "palm", "bush", "mulberry", "juniper", "vine", "ivy", "rose", "lily", "tulip", "daffodil", "honeysuckle", "fuschia", "hazel", "walnut", "almond", "lime", "lemon", "apple", "blossom", "bloom", "crocus", "rose", "buttercup", "dandelion", "iris", "carnation", "fern", "root", "branch", "leaf", "seed", "flower", "petal", "pollen", "orchid", "mangrove", "cypress", "sequoia", "sage", "heather", "snapdragon", "daisy", "mountain", "hill", "alpine", "chestnut", "valley", "glacier", "forest", "grove", "glen", "tree", "thorn", "stump", "desert", "canyon", "dune", "oasis", "mirage", "well", "spring", "meadow", "field", "prairie", "grass", "tundra", "island", "shore", "sand", "shell", "surf", "wave", "foam", "tide", "lake", "river", "brook", "stream", "pool", "pond", "sun", "sprinkle", "shade", "shadow", "rain", "cloud", "storm", "hail", "snow", "sleet", "thunder", "lightning", "wind", "hurricane", "typhoon", "dawn", "sunrise", "morning", "noon", "twilight", "evening", "sunset", "midnight", "night", "sky", "star", "stellar", "comet", "nebula", "quasar", "solar", "lunar", "planet", "meteor", "sprout", "pear", "plum", "kiwi", "berry", "apricot", "peach", "mango", "pineapple", "coconut", "olive", "ginger", "root", "plain", "fancy", "stripe", "spot", "speckle", "spangle", "ring", "band", "blaze", "paint", "pinto", "shade", "tabby", "brindle", "patch", "calico", "checker", "dot", "pattern", "glitter", "glimmer", "shimmer", "dull", "dust", "dirt", "glaze", "scratch", "quick", "swift", "fast", "slow", "clever", "fire", "flicker", "flash", "spark", "ember", "coal", "flame", "chocolate", "vanilla", "sugar", "spice", "cake", "pie", "cookie", "candy", "caramel", "spiral", "round", "jelly", "square", "narrow", "long", "short", "small", "tiny", "big", "giant", "great", "atom", "peppermint", "mint", "butter", "fringe", "rag", "quilt", "truth", "lie", "holy", "curse", "noble", "sly", "brave", "shy", "lava", "foul", "leather", "fantasy", "keen", "luminous", "feather", "sticky", "gossamer", "cotton", "rattle", "silk", "satin", "cord", "denim", "flannel", "plaid", "wool", "linen", "silent", "flax", "weak", "valiant", "fierce", "gentle", "rhinestone", "splash", "north", "south", "east", "west", "summer", "winter", "autumn", "spring", "season", "equinox", "solstice", "paper", "motley", "torch", "ballistic", "rampant", "shag", "freckle", "wild", "free", "chain", "sheer", "crazy", "mad", "candle", "ribbon", "lace", "notch", "wax", "shine", "shallow", "deep", "bubble", "harvest", "fluff", "venom", "boom", "slash", "rune", "cold", "quill", "love", "hate", "garnet", "zircon", "power", "bone", "void", "horn", "glory", "cyber", "nova", "hot", "helix", "cosmic", "quark", "quiver", "holly", "clover", "polar", "regal", "ripple", "ebony", "wheat", "phantom", "dew", "chisel", "crack", "chatter", "laser", "foil", "tin", "clever", "treasure", "maze", "twisty", "curly", "fortune", "fate", "destiny", "cute", "slime", "ink", "disco", "plume", "time", "psychadelic", "relic", "fossil", "water", "savage", "ancient", "rapid", "road", "trail", "stitch", "button", "bow", "nimble", "zest", "sour", "bitter", "phase", "fan", "frill", "plump", "pickle", "mud", "puddle", "pond", "river", "spring", "stream", "battle", "arrow", "plume", "roan", "pitch", "tar", "cat", "dog", "horse", "lizard", "bird", "fish", "saber", "scythe", "sharp", "soft", "razor", "neon", "dandy", "weed", "swamp", "marsh", "bog", "peat", "moor", "muck", "mire", "grave", "fair", "just", "brick", "puzzle", "skitter", "prong", "fork", "dent", "dour", "warp", "luck", "coffee", "split", "chip", "hollow", "heavy", "legend", "hickory", "mesquite", "nettle", "rogue", "charm", "prickle", "bead", "sponge", "whip", "bald", "frost", "fog", "oil", "veil", "cliff", "volcano", "rift", "maze", "proud", "dew", "mirror", "shard", "salt", "pepper", "honey", "thread", "bristle", "ripple", "glow", "zenith"}

var nouns = []string{"head", "crest", "crown", "tooth", "fang", "horn", "frill", "skull", "bone", "tongue", "throat", "voice", "nose", "snout", "chin", "eye", "sight", "seer", "speaker", "singer", "song", "chanter", "howler", "chatter", "shrieker", "shriek", "jaw", "bite", "biter", "neck", "shoulder", "fin", "wing", "arm", "lifter", "grasp", "grabber", "hand", "paw", "foot", "finger", "toe", "thumb", "talon", "palm", "touch", "racer", "runner", "hoof", "fly", "flier", "swoop", "roar", "hiss", "hisser", "snarl", "dive", "diver", "rib", "chest", "back", "ridge", "leg", "legs", "tail", "beak", "walker", "lasher", "swisher", "carver", "kicker", "roarer", "crusher", "spike", "shaker", "charger", "hunter", "weaver", "crafter", "binder", "scribe", "muse", "snap", "snapper", "slayer", "stalker", "track", "tracker", "scar", "scarer", "fright", "killer", "death", "doom", "healer", "saver", "friend", "foe", "guardian", "thunder", "lightning", "cloud", "storm", "forger", "scale", "hair", "braid", "nape", "belly", "thief", "stealer", "reaper", "giver", "taker", "dancer", "player", "gambler", "twister", "turner", "painter", "dart", "drifter", "sting", "stinger", "venom", "spur", "ripper", "swallow", "devourer", "knight", "lady", "lord", "queen", "king", "master", "mistress", "prince", "princess", "duke", "dutchess", "samurai", "ninja", "knave", "slave", "servant", "sage", "wizard", "witch", "warlock", "warrior", "jester", "paladin", "bard", "trader", "sword", "shield", "knife", "dagger", "arrow", "bow", "fighter", "bane", "follower", "leader", "scourge", "watcher", "cat", "panther", "tiger", "cougar", "puma", "jaguar", "ocelot", "lynx", "lion", "leopard", "ferret", "weasel", "wolverine", "bear", "raccoon", "dog", "wolf", "kitten", "puppy", "cub", "fox", "hound", "terrier", "coyote", "hyena", "jackal", "pig", "horse", "donkey", "stallion", "mare", "zebra", "antelope", "gazelle", "deer", "buffalo", "bison", "boar", "elk", "whale", "dolphin", "shark", "fish", "minnow", "salmon", "ray", "fisher", "otter", "gull", "duck", "goose", "crow", "raven", "bird", "eagle", "raptor", "hawk", "falcon", "moose", "heron", "owl", "stork", "crane", "sparrow", "robin", "parrot", "cockatoo", "carp", "lizard", "gecko", "iguana", "snake", "python", "viper", "boa", "condor", "vulture", "spider", "fly", "scorpion", "heron", "oriole", "toucan", "bee", "wasp", "hornet", "rabbit", "bunny", "hare", "brow", "mustang", "ox", "piper", "soarer", "flasher", "moth", "mask", "hide", "hero", "antler", "chill", "chiller", "gem", "ogre", "myth", "elf", "fairy", "pixie", "dragon", "griffin", "unicorn", "pegasus", "sprite", "fancier", "chopper", "slicer", "skinner", "butterfly", "legend", "wanderer", "rover", "raver", "loon", "lancer", "glass", "glazer", "flame", "crystal", "lantern", "lighter", "cloak", "bell", "ringer", "keeper", "centaur", "bolt", "catcher", "whimsey", "quester", "rat", "mouse", "serpent", "wyrm", "gargoyle", "thorn", "whip", "rider", "spirit", "sentry", "bat", "beetle", "burn", "cowl", "stone", "gem", "collar", "mark", "grin", "scowl", "spear", "razor", "edge", "seeker", "jay", "ape", "monkey", "gorilla", "koala", "kangaroo", "yak", "sloth", "ant", "roach", "weed", "seed", "eater", "razor", "shirt", "face", "goat", "mind", "shift", "rider", "face", "mole", "vole", "pirate", "llama", "stag", "bug", "cap", "boot", "drop", "hugger", "sargent", "snagglefoot", "carpet", "curtain"}

func randomNoun() string {
	return nouns[rand.Intn(len(nouns)-1)]
}

func randomAdjective() string {
	return adjectives[rand.Intn(len(adjectives)-1)]
}

func GenerateStupidName() string {
	return GenerateRandomNameN(0, true, true)
}

func GenerateUsername() string {
	return GenerateUsernameN(0)
}

func GenerateUsernameN(maxLength int) string {
	return GenerateRandomNameN(maxLength, false, false)
}

func GenerateRandomNameN(maxLength int, surname bool, capitalize bool) string {
	name := ""
	for ok := true; ok; ok = maxLength > 0 && len(name) > maxLength {
		name = randomNoun() + randomAdjective()
		if surname {
			name += " " + randomNoun()
		}
		if capitalize {
			name = strings.Title(name)
		}
	}
	return name
}
