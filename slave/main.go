package slave

import (
	"log"
)

// TODO
const modelAddress = ":8081"

func main() {
	m := &model{}
	if err := listenModel(m, modelAddress); err != nil {
		log.Fatal("failed to start model service", err.Error())
	}
}
