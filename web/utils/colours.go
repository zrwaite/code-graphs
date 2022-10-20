package utils

func GetColour(language string) string {
	switch language {
	case "TypeScript":
		return "#0099ff"
	case "JavaScript":
		return "#ecec13"
	case "C":
		return "#666666"
	case "C#":
		return "#9332bf"
	case "JSON":
		return "#339933"
	case "PHP":
		return "#9999ff"
	case "Python":
		return "#0066cc"
	case "HTML":
		return "#ff471a"
	case "Docker":
		return "#1aa3ff"
	case "SQL":
		return "#e6b800"
	case "Java":
		return "#e60000"
	case "Dart":
		return "rgb(23, 174, 255)"
	case "SCSS":
		return "rgb(201, 85, 146)"
	case "CSS":
		return "#1337ed"
	case "Rust":
		return "#ff8b6e"
	case "Racket":
		return "rgb(100, 13, 20)"
	case "Markdown":
		return "#333333"
	case "C++":
		return "rgb(83, 136, 200)"
	case "VHDL":
		return "grey"
	case "Go":
		return "rgb(20, 156, 206)"
	case "Swift":
		return "rgb(234, 80, 41)"
	case "GraphQL":
		return "rgb(215, 0, 135)"
	case "Svelte":
		return "rgb(235, 62, 39)"
	case "Ruby":
		return "rgb(217, 10, 0)"
	case "Scala":
		return "rgb(153, 0, 0)"
	case "Bash":
		return "rgb(3, 252, 53)"
	case "Arduino":
		return "rgb(12, 145, 166)"
	case "Assembly":
		return "#888888"
	case "Groovy":
		return "rgb(43, 158, 204)"
	default:
		return "lightgrey"
	}
}
