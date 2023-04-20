package wowapi

import (
	"encoding/json"
	"fmt"
)

type Equipment struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	EquippedItems []struct {
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		ModifiedAppearanceId float64 `json:"modified_appearance_id"`
		Stats                []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value   float64 `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
			} `json:"display"`
		} `json:"stats"`
		Requirements struct {
			Level struct {
				Value         float64 `json:"value"`
				DisplayString string  `json:"display_string"`
			} `json:"level"`
			PlayableClasses struct {
				Links []struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"links"`
				DisplayString string `json:"display_string"`
			} `json:"playable_classes"`
		} `json:"requirements"`
		NameDescription struct {
			DisplayString string `json:"display_string"`
			Color         struct {
				B float64 `json:"b"`
				A float64 `json:"a"`
				R float64 `json:"r"`
				G float64 `json:"g"`
			} `json:"color"`
		} `json:"name_description"`
		Context float64 `json:"context"`
		Media   struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"media"`
		Binding struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		Transmog struct {
			ItemModifiedAppearanceId float64 `json:"item_modified_appearance_id"`
			Item                     struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item"`
			DisplayString string `json:"display_string"`
		} `json:"transmog"`
		Durability struct {
			Value         float64 `json:"value"`
			DisplayString string  `json:"display_string"`
		} `json:"durability"`
		Level struct {
			Value         float64 `json:"value"`
			DisplayString string  `json:"display_string"`
		} `json:"level"`
		Quantity  float64 `json:"quantity"`
		BonusList []struct {
			BonusList float64 `json:"bonus_list"`
		} `json:"bonus_list"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_class"`
		InventoryType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Armor struct {
			Value   float64 `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					B float64 `json:"b"`
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
				} `json:"color"`
			} `json:"display"`
		} `json:"armor"`
		Set struct {
			ItemSet struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item_set"`
			Items []struct {
				IsEquipped bool `json:"is_equipped"`
				Item       struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"item"`
			} `json:"items"`
			Effects []struct {
				DisplayString string  `json:"display_string"`
				RequiredCount float64 `json:"required_count"`
				IsActive      bool    `json:"is_active"`
			} `json:"effects"`
			DisplayString string `json:"display_string"`
		} `json:"set"`
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"item"`
		Sockets []struct {
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Id float64 `json:"id"`
			} `json:"media"`
			SocketType struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"socket_type"`
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item"`
			DisplayString string `json:"display_string"`
		} `json:"sockets"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Name         string `json:"name"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_subclass"`
		SellPrice struct {
			DisplayStrings struct {
				Header string `json:"header"`
				Gold   string `json:"gold"`
				Silver string `json:"silver"`
				Copper string `json:"copper"`
			} `json:"display_strings"`
			Value float64 `json:"value"`
		} `json:"sell_price"`
	} `json:"equipped_items"`
	EquippedItemSets []struct {
		ItemSet struct {
			Id  float64 `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
		} `json:"item_set"`
		Items []struct {
			IsEquipped bool `json:"is_equipped"`
			Item       struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item"`
		} `json:"items"`
		Effects []struct {
			IsActive      bool    `json:"is_active"`
			DisplayString string  `json:"display_string"`
			RequiredCount float64 `json:"required_count"`
		} `json:"effects"`
		DisplayString string `json:"display_string"`
	} `json:"equipped_item_sets"`
}

func (req RequestFunc) CharacterEquipment(realm string, name string) (s Equipment, err error) {
	url := fmt.Sprintf("/profile/wow/character/%s/%s/equipment", realm, name)
	body, err := req(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return
	}
	return
}
