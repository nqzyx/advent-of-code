package pipes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
)

type ASCII rune

const (
	ASCII_EQUAL   ASCII = '='
	ASCII_MINUS   ASCII = '-'
	ASCII_NEWLINE ASCII = '\n'
	ASCII_PIPE    ASCII = '|'
	ASCII_PLUS    ASCII = '+'
	ASCII_SPACE   ASCII = 0x20
	ASCII_STAR    ASCII = '*'
)

func (a ASCII) String() string {
	return string(a)
}

func formatHeadersAndFooters(columns int) ([]string, []string, string) {
	l1Hdr := "****"
	l2Hdr := "****"
	l3Hdr := "===="

	l1Fmt := "%v" + strings.Repeat(string(ASCII_SPACE), 9)
	l2Fmt := "%v" + string(ASCII_SPACE)
	l3Fmt := strings.Repeat(string(ASCII_EQUAL), 10)

	l1Sep := string(ASCII_PIPE) + string(ASCII_SPACE)
	l2Sep := string(ASCII_PIPE) + string(ASCII_SPACE)
	l3Sep := string(ASCII_PLUS) + string(ASCII_EQUAL)

	l1sb := new(strings.Builder)
	l1sb.WriteString(l1Hdr + l1Sep)
	l2sb := new(strings.Builder)
	l2sb.WriteString(l2Hdr + l2Sep)
	l3sb := new(strings.Builder)
	l3sb.WriteString(l3Hdr + l3Sep)

	// maxLen := utils.Min[int](c%5, columns-c)*2
	for c := 0; c < columns; c++ {
		// every column
		l2sb.WriteString(fmt.Sprintf(l2Fmt, c%10))
		switch c % 5 {
		case 0: // before every 5th column
			maxLength := utils.Min(5, columns-c) * 2
			l1sb.WriteString(fmt.Sprintf(l1Fmt, c/10)[:maxLength])
			l3sb.WriteString(l3Fmt[:maxLength])
		case 4: // after every 5th column
			l1sb.WriteString(l1Sep)
			l2sb.WriteString(l2Sep)
			l3sb.WriteString(l3Sep)
		}
		if c == columns-1 { // final {
			l1sb.WriteString(l1Sep[:1])
			l2sb.WriteString(l2Sep[:1])
			l3sb.WriteString(l3Sep[:1])
		}
	}

	l1sb.WriteString(l1Hdr + string(ASCII_NEWLINE))
	l2sb.WriteString(l2Hdr + string(ASCII_NEWLINE))
	l3sb.WriteString(l3Hdr + string(ASCII_NEWLINE))

	headers := make([]string, 0, 3)
	headers = append(headers, l1sb.String(), l2sb.String(), l3sb.String())
	footers := make([]string, 0, len(headers[1:]))
	copy(footers, headers[1:])
	slices.Reverse(footers)
	separator := strings.ReplaceAll(l3sb.String(), string(ASCII_EQUAL), string(ASCII_MINUS))

	return headers, footers, separator
}

func formatTileRow(r int, row TileList) string {
	leftFmt := "%3v" + string(ASCII_SPACE)
	rightFmt := "%3v"
	grpSep := string(ASCII_PIPE) + string(ASCII_SPACE)
	tileFmt := "%v" + string(ASCII_SPACE)

	sb := new(strings.Builder)

	var rOut any

	if r%5 == 0 {
		rOut = r
	} else {
		rOut = r % 10
	}

	sb.WriteString(fmt.Sprintf(leftFmt, rOut) + grpSep)

	columns := len(row)
	for c, tile := range row {
		var tileChar string
		if tile.OnPath {
			tileChar = tile.PipeType.String()
		} else {
			tileChar = ASCII_SPACE.String()
		}
		sb.WriteString(fmt.Sprintf(tileFmt, tileChar))
		if c%5 == 4 || c == columns-1 {
			sb.WriteString(grpSep)
		}
	}

	rOutRunesReversed := []rune(fmt.Sprintf(rightFmt, rOut))
	slices.Reverse(rOutRunesReversed)

	sb.WriteString(string(rOutRunesReversed[:1]) + string(ASCII_NEWLINE))

	return sb.String()
}

func (m Map) String() string {
	lineLength := m.ColCount()
	pageHeaders, pageFooters, groupSeparator := formatHeadersAndFooters(lineLength)

	result := new(strings.Builder)
	for _, pageHeader := range pageHeaders {
		result.WriteString(pageHeader)
	}

	// TODO: Refactor this
	// for r, tileRow := range m.Tiles {
	// 	result.WriteString(formatTileRow(r, *tileRow))
	// 	if r%5 == 4 && r < m.RowCount()-1 {
	result.WriteString(groupSeparator)
	// 	}
	// }

	for _, footer := range pageFooters {
		result.WriteString(footer)
	}

	return result.String()
}
