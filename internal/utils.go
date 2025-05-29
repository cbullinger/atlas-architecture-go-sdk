package internal

import (
	"io"
	"log"
)

// SafeClose closes c and logs a warning on error.
func SafeClose(c io.Closer) {
	if c != nil {
		if err := c.Close(); err != nil {
			log.Printf("warning: close failed: %v", err)
		}
	}
}

// SafeCopy copies source to destination, and propagates any error after logging.
func SafeCopy(dst io.Writer, src io.Reader) error {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("warning: copy failed: %v", err)
		return err
	}
	return nil
}
