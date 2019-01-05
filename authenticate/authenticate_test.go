package authenticate

import "testing"
import "encoding/json"

func TestEncodingPayload(t *testing.T) {
	j, err := json.Marshal(Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		UserName:    "mojang account name",
		Password:    "mojang account password",
		ClientToken: "client identifier",
		RequestUser: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(j))
}

func TestAuthenticate(t *testing.T) {
	resp, err := Authenticate("", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
