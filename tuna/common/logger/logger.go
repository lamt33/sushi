package logger

import (
	"encoding/json"
	"fmt"
	"log"
)

type Entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity,omitempty"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`
}

// String renders an entry structure to the JSON format expected by Cloud Logging.
func (e Entry) String() string {
	if e.Severity == "" {
		e.Severity = "INFO"
	}
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}

func Info(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	out(fmt.Sprintf("ℹ️: %s", msg), "INFO")

}

func Error(m string, a ...interface{}) error {
	msg := fmt.Sprintf(m, a...)
	out(fmt.Sprintf("❌: %s", msg), "ERROR")
	return fmt.Errorf(msg)
}

func out(msg, severity string) {
	log.Println(Entry{
		Severity: severity,
		Message:  msg,
	})
}
