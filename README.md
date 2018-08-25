## Usage

```golang
uesrname := "blars"
password := "tacoman"
url := "https://api.stormondemand.com"

api := storm.NewAPI(username, password, url)

// Get storm servers
api.StormServer.List()

// Create a storm server
stormServerParams := &storm.StormServerParams{
  configID: 123,
  hostname: "blars.tacoman.com"
  zoneID: 123,
  password: "123",
  publicKey: "yourkey"
}
api.StormServer.Create(stormServerParams)
```
