package internal

import (
	"fmt"
	"os"
	"strings"
)

func detect_color_tag_type(tag string) ColorTagType {
	if strings.Count(tag, ";") > 1 {
		return ColorTagCombined
	}

	for color_string := range StandardColorList {
		if strings.Contains(tag, color_string) {
			return ColorTagStandard
		}
	}

	return ColorTagHex
}

func ParseString(s string) string {
	/*
		Initial variables
	*/
	var color_tags []string

	/*
		Functions
	*/

	/* Detect tags */
	var in_tag bool
	var color_tag_buffer string
	for _, char := range s {
		if char == '<' {
			in_tag = true
		}

		if in_tag {
			color_tag_buffer += string(char)
			if char == '>' {
				in_tag = false

				color_tags = append(color_tags, color_tag_buffer)
				color_tag_buffer = ""
				continue
			}
		}
	}

	/* Parse tags */
	for _, tag := range color_tags {
		color := tag[1 : len(tag)-1]
		if color == "/" {
			color = "reset"
		}
		if color == "" {
			s = strings.ReplaceAll(s, tag, "")
			continue
		}

		color_tag_type := detect_color_tag_type(color)

		switch color_tag_type {
		case ColorTagHex:
			if color[0] == '#' {
				color = color[1:]
			}

			rgb_color, err := NewColorFromHex(color)
			if err != nil {
				fmt.Println(err, `Codigo de linea: 8f2cc96b-1391-4a86-88b5-03ed2f58a789`)
				os.Exit(2)
			}

			ansi_code := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", rgb_color.R, rgb_color.B, rgb_color.G)

			// Replace colors
			s = strings.ReplaceAll(s, tag, ansi_code)

		case ColorTagStandard:
			ansi_code := fmt.Sprintf("\x1b[1;%dm", StandardColorList[color])

			s = strings.ReplaceAll(s, tag, ansi_code)

		case ColorTagCombined:
			sub_color_tag_splitted := strings.Split(color, ";")

			// Fix background
			var color_tag_splitted_fixed []string
			var background_tag_splitted bool
			for tag_position, sub_tag := range sub_color_tag_splitted {
				if sub_tag == "b" {
					background_tag_splitted = true
					continue
				}

				if background_tag_splitted {
					fixed_tag := strings.Join([]string{sub_color_tag_splitted[tag_position-1], sub_tag}, ";")

					color_tag_splitted_fixed = append(color_tag_splitted_fixed, fixed_tag)
					background_tag_splitted = false
					continue
				}

				color_tag_splitted_fixed = append(color_tag_splitted_fixed, sub_tag)
			}

			// Parse fixed tags
			var combined_tag string
			for _, sub_tag := range color_tag_splitted_fixed {
				color_tag_type := detect_color_tag_type(sub_tag)

				switch color_tag_type {
				case ColorTagHex:
					if color[0] == '#' {
						color = color[1:]
					}

					rgb_color, err := NewColorFromHex(sub_tag)
					if err != nil {
						fmt.Println(err, `Codigo de linea: 8f2cc96b-1391-4a86-88b5-03ed2f58a789`)
						os.Exit(2)
					}

					ansi_code := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", rgb_color.R, rgb_color.B, rgb_color.G)

					// Replace colors
					combined_tag += ansi_code

				case ColorTagStandard:
					ansi_code := fmt.Sprintf("\x1b[1;%dm", StandardColorList[sub_tag])
					combined_tag += ansi_code
				}
			}

			// Replace combined tag
			s = strings.ReplaceAll(s, tag, combined_tag)

		case ColorTagStyle:
			var ansi_id uint8 = StandardColorList[color]
			if strings.Contains(color, "/") {
				if ansi_id == 1 {
					ansi_id++
				}
				ansi_id += 10
			}

			ansi_code := fmt.Sprintf("\x1b[%dm", ansi_id)

			s = strings.ReplaceAll(s, tag, ansi_code)
		}

	}

	return s
}
