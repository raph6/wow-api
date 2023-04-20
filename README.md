# World of Warcraft API Golang

This is a simple wrapper for the World of Warcraft API written in Go.

## Usage

Installation
```sh
go get -u github.com/raph6/wowapi
```

Usage
```go
import "github.com/raph6/wowapi"

func main() {
    // API_CLIENT_ID, API_SECRET, region, lang
    client := wowapi.NewClient("xx_API_CLIENT_ID_xx", "xx_API_SECRET_xx", "eu", "fr_FR")
    // accepted region ->  eu | us | kr | tw
    // accepted lang ->  en_US | es_MX | pt_BR | en_GB | es_ES | fr_FR | ru_RU | de_DE | pt_PT | it_IT | zh_TW | ko_KR

    // Realm, Name
    titles, err := client.CharacterTitles("Kirin-Tor", "Vimdiesel")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(titles)
}
```

Available methods
```go
client.CharacterAchievementsStatistics(realm string, name string)
client.CharacterAchievements(realm string, name string)
client.CharacterAppearance(realm string, name string)
client.CharacterCharacterMedia(realm string, name string)
client.CharacterCharacterStatus(realm string, name string)
client.CharacterCollectionsHeirlooms(realm string, name string)
client.CharacterCollectionsMounts(realm string, name string)
client.CharacterCollectionsPets(realm string, name string)
client.CharacterCollectionsToys(realm string, name string)
client.CharacterCollections(realm string, name string)
client.CharacterEncountersDungeons(realm string, name string)
client.CharacterEncountersRaids(realm string, name string)
client.CharacterEncounters(realm string, name string)
client.CharacterEquipment(realm string, name string)
client.CharacterMythicKeystoneProfile(realm string, name string)
client.CharacterPvpSummary(realm string, name string)
client.CharacterQuestsCompleted(realm string, name string)
client.CharacterQuests(realm string, name string)
client.CharacterSoulbinds(realm string, name string)
```

Todo
- [ ] All tests
- [ ] Raider.io
- [ ] Hunter pets methods


## Tests

```sh
API_CLIENT_ID=xxxxx API_SECRET=xxxxx go test
```