package protocol

import (
    "errors"
    "encoding/binary"
)

const (
    MESSAGE_MAX = 4096
    MESSAGE_LENGTH_SIZE = 4
)

/**
    record is a single message transerred using connection
    record will be construction as |nstr|len|message|...|lenn|strn|
*/
func Deserialize(record []byte) ([]string, error) {
    if(4 > len(record)) {
        return nil, errors.New("record is too little")
    }

    if(MESSAGE_MAX < len(record)) {
        return nil, errors.New("record is too large")
    }

    messages := []string{}

    // first 4 bytes will be total number of string
    nMessage := int(binary.BigEndian.Uint32(record[0:MESSAGE_LENGTH_SIZE]))

    pos := MESSAGE_LENGTH_SIZE

    for n := 1; n <= nMessage; n++ {
        lenEndPos := pos + MESSAGE_LENGTH_SIZE
        len := int(binary.BigEndian.Uint32(record[pos:lenEndPos]))


        messageEndPos := lenEndPos + len
        message := string(record[lenEndPos: messageEndPos])

        messages = append(messages, message)
        pos = messageEndPos
    }

    return messages, nil
}

func Serialize(messages []string) ([]byte, error) {
    record := make([]byte, 4)
    nMessage := len(messages)
    
    binary.BigEndian.PutUint32(record, uint32(nMessage))

    for n := 1; n <= nMessage; n++ {
        message := messages[n-1]
        mLen := make([]byte, 4)

        binary.BigEndian.PutUint32(mLen, uint32(len(message)))

        record = append(record, mLen...)
        record = append(record, []byte(message)...)
    }

    return record, nil
}
