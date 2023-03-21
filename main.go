package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Table name: ")
	scanner.Scan()
	tableName := scanner.Text()

	fmt.Print("Columns (comma-separated): ")
	scanner.Scan()
	columnsStr := scanner.Text()

	columns := strings.Split(columnsStr, ",")

	var values []string
	for scanner.Scan() {
		row := scanner.Text()
		rowValues := strings.Split(row, "\t")
		valueStrs := make([]string, len(rowValues))

		for i, value := range rowValues {
			valueStrs[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(value, "'", "''"))
		}

		values = append(values, fmt.Sprintf("(%s)", strings.Join(valueStrs, ",")))
	}

	fmt.Printf("INSERT INTO %s (%s) VALUES\n%s;\n", tableName, strings.Join(columns, ","), strings.Join(values, ",\n"))
}
func Run(r io.Reader, w io.Writer, tableName string, columns string) error {
	scanner := bufio.NewScanner(r)
	var values []string
	cols := strings.Split(columns, ",")
	for scanner.Scan() {
		var id, title string
		line := scanner.Text()
		re := regexp.MustCompile(`\s*(\d+)\s*\|\s*(.+)`)
		match := re.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}
		id, title = match[1], match[2]

		var value strings.Builder
		value.WriteString("(")
		for i, col := range cols {
			if col == "id" {
				value.WriteString(fmt.Sprintf("'%s'", id))
			} else if col == "title" {
				value.WriteString(fmt.Sprintf("'%s'", strings.ReplaceAll(title, "'", "''")))
			}
			if i != len(cols)-1 {
				value.WriteString(",")
			}
		}
		value.WriteString(")")
		values = append(values, value.String())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprintf(w, "INSERT INTO %s (%s) VALUES\n%s;\n", tableName, columns, strings.Join(values, ",\n"))
	return nil
}
