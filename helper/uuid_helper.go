package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

// Convert AD177BB820016499A15A54A9 to 00000000-AD17-7BB8-2001-6499A15A54A9
func UuidFromObjectId(objectId string) string {
	s0 := "00000000"
	s1 := objectId[:4]
	s2 := objectId[4:8]
	s3 := objectId[8:12]
	s4 := objectId[12:]

	return fmt.Sprintf("%s-%s-%s-%s-%s", s0, s1, s2, s3, s4)
}

// Convert 00000000-AD17-7BB8-2001-6499A15A54A9 to AD177BB820016499A15A54A9
func ParseUUIDToObjectID(id string) string {
	return fmt.Sprintf("%s%s%s%s",
		id[9:13],
		id[14:18],
		id[19:23],
		id[24:],
	)
}

// Save value to json and save to file with given path
func SaveJsonToFile(path string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	log.Println("Saved to", path)

	return nil
}

func StringToUUID(id string) uuid.UUID {
	return uuid.Must(uuid.Parse((id)))
}
