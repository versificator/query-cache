package main

import (
	"database/sql"
	"bufio"
	"io"
	"crypto/sha1"
	"hash"
	"log"
)

type cache struct{
	hashCode hash.Hash
	startTimestamp int8
	endTimestamp int8
}

type Writer struct {
	Delimiter  string // Field delimiter (set to ',' by NewWriter)
	Quote      string // Quote character
	Escape     string // Escape character
	Terminator string // Character to end each line
	w          *bufio.Writer
}


// NewWriter returns a new Writer that writes to w.
func TSVWriter(w io.Writer) *Writer {
	return &Writer{
		Delimiter:  ",",
		Quote:      "\"",
		Escape:     "\\",
		Terminator: "\n",
		w:          bufio.NewWriter(w),
	}
}

func checkIfCacheExists(jsonRaw *string) bool{
	return false
}

func checkQuery(jsonRaw *string) cache {
	s := cleanJSON(jsonRaw)
	log.Printf("Cleaned json %s", s)

	hashCode := getHash(s)
	log.Printf("Hash %s", hashCode)
	return cache{hashCode, 0, 0}
}



func getHash(jsonRaw []byte) hash.Hash{
	h := sha1.New()
	h.Write(jsonRaw)
	log.Printf("getHash: %s",h)
	return h
}
// Writer writes a single CSV record to w along with any necessary quoting.
func (w *Writer) Write(record []sql.RawBytes) (buf int, err error) {
	for n, field := range record {
		// Shortcut exit for empty strings
		if n > 0 {
			if _, err = w.w.WriteString(w.Delimiter); err != nil {
				return
			}
		}

		// Check if and escape/translate if field is NULL
		if field == nil {
			_, err = w.w.WriteString(w.Escape)
			_, err = w.w.WriteString("N")
			continue
		}

		// Write quote character if set
		if w.Quote != "" {
			if _, err = w.w.WriteString(w.Quote); err != nil {
				return
			}
		}

		// We need to examine each byte to determine if special characters need to be escaped
		for _, f := range field {
			switch string(f) {
			case w.Delimiter:
				if w.Quote == "" {
					_, err = w.w.WriteString(w.Escape)
					_, err = w.w.WriteString(w.Delimiter)
				} else {
					_, err = w.w.WriteString(w.Delimiter)
				}
			case w.Quote:
				_, err = w.w.WriteString(w.Escape)
				_, err = w.w.WriteString(w.Quote)
			case w.Escape:
				_, err = w.w.WriteString(w.Escape)
				_, err = w.w.WriteString(w.Escape)
			case string(0x00):
				_, err = w.w.WriteString(w.Escape)
				_, err = w.w.WriteRune('0')
			case string(0x0A):
				_, err = w.w.WriteString(w.Escape)
				err = w.w.WriteByte(f)
			default:
				err = w.w.WriteByte(f)

			}
			if err != nil {
				return
			}
		}

		// Write quote character if set
		if w.Quote != "" {
			if _, err = w.w.WriteString(w.Quote); err != nil {
				return
			}
		}
	}

	// Write line terminator
	_, err = w.w.WriteString(w.Terminator)

	// Return the number of bytes written to the current buffer
	buf = w.w.Buffered()

	return buf, err
}


// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

// Error reports any error that has occurred during a previous Write or Flush.
func (w *Writer) Error() error {
	_, err := w.w.Write(nil)
	return err
}

// WriteAll writes multiple CSV records to w using Write and then calls Flush.
func (w *Writer) WriteAll(records [][]sql.RawBytes) (err error) {
	for _, record := range records {
		_, err = w.Write(record)
		if err != nil {
			return err
		}
	}
	return w.w.Flush()
}