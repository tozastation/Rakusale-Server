package util

// Enum
type Vegetable int

// Const
const (
	SQUASH      Vegetable = iota // 南瓜
	CABBAGE     Vegetable = iota // キャベツ
	CUCUMBER    Vegetable = iota // キュウリ
	SWEETPOTATO Vegetable = iota // さつまいも
	POTATO      Vegetable = iota // じゃがいも
	RADISH      Vegetable = iota // 大根
	ONION       Vegetable = iota // 玉ねぎ
	CARROT      Vegetable = iota // 人参
	BELLPEPPER  Vegetable = iota // ピーマン
	SPINACH     Vegetable = iota // ホウレンソウ
	LETTUCE     Vegetable = iota // レタス
)

// ConvertEnumVegetable is
func ConvertEnumVegetable(v string) string {
	switch v {
	case "SQUASH":
		return "カボチャ"
	case "CABBAGE":
		return "キャベツ"
	case "CUCUMBER":
		return "キュウリ"
	case "SWEETPOTATO":
		return "さつまいも"
	case "POTATO":
		return "じゃがいも"
	case "RADISH":
		return "大根"
	case "ONION":
		return "玉ねぎ"
	case "CARROT":
		return "人参"
	case "BELLPEPPER":
		return "ピーマン"
	case "SPINACH":
		return "ホウレンソウ"
	case "LETTUCE":
		return "レタス"
	default:
		return "Unknown"
	}
}
