package restyclient

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func FetchUser(userID uint) (map[string]interface{}, error) {
	client := resty.New()
	resp, err := client.R().
		SetResult(map[string]interface{}{}).
		Get(os.Getenv("USER_SERVICE_URL") + "/users/" + string(rune(userID)))

	if err != nil {
		log.Println("Fehler beim Abrufen des Benutzers:", err)
		return nil, err
	}

	return *resp.Result().(*map[string]interface{}), nil
}
