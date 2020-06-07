package dbr

import (
	"log"
	"time"
)

type sqlEventReceiver struct {
	costThreshold int64
}

// Event receives a simple notification when various events occur
func (s *sqlEventReceiver) Event(eventName string) {
	log.Println("DB Event", "name", eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data
func (s *sqlEventReceiver) EventKv(eventName string, kvs map[string]string) {
	log.Println("DB EventKv", "name", eventName, "kv", kvs)
}

// EventErr receives a notification of an error if one occurs
func (s *sqlEventReceiver) EventErr(eventName string, err error) error {
	log.Println("DB EventErr", "name", eventName, "err", err)
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data
func (s *sqlEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	log.Println("DB EventErrKv", "name", eventName, "kv", kvs, "err", err)
	return err
}

// Timing receives the time an event took to happen
func (s *sqlEventReceiver) Timing(eventName string, nanoseconds int64) {
	t := int64(time.Duration(nanoseconds) / time.Millisecond)
	if t > s.costThreshold {
		log.Println("DB Timing", "name", eventName, "cost", time.Duration(nanoseconds).String())
	}
}

// TimingKv receives the time an event took to happen along with optional key/value data
func (s *sqlEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	t := int64(time.Duration(nanoseconds) / time.Millisecond)
	if t > s.costThreshold {
		log.Println("DB TimingKv", "name", eventName, "kv", kvs, "cost", time.Duration(nanoseconds).String())
	}
}
