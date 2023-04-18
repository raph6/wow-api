package wowapi

type Equipment struct {
	EquippedItems []struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"item"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"media"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_class"`
		Durability struct {
			Value         float64 `json:"value"`
			DisplayString string  `json:"display_string"`
		} `json:"durability"`
		Sockets []struct {
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
			Media         struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Id float64 `json:"id"`
			} `json:"media"`
		} `json:"sockets"`
		Quantity  float64 `json:"quantity"`
		BonusList []struct {
			BonusList float64 `json:"bonus_list"`
		} `json:"bonus_list"`
		Name         string `json:"name"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_subclass"`
		InventoryType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value   float64 `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
				} `json:"color"`
			} `json:"display"`
		} `json:"stats"`
		Armor struct {
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
				} `json:"color"`
			} `json:"display"`
			Value float64 `json:"value"`
		} `json:"armor"`
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
		NameDescription struct {
			DisplayString string `json:"display_string"`
			Color         struct {
				B float64 `json:"b"`
				A float64 `json:"a"`
				R float64 `json:"r"`
				G float64 `json:"g"`
			} `json:"color"`
		} `json:"name_description"`
		Set struct {
			ItemSet struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item_set"`
			Items []struct {
				Item struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"item"`
				IsEquipped bool `json:"is_equipped"`
			} `json:"items"`
			Effects []struct {
				DisplayString string  `json:"display_string"`
				RequiredCount float64 `json:"required_count"`
				IsActive      bool    `json:"is_active"`
			} `json:"effects"`
			DisplayString string `json:"display_string"`
		} `json:"set"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Context float64 `json:"context"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		ModifiedAppearanceId float64 `json:"modified_appearance_id"`
		Binding              struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		SellPrice struct {
			Value          float64 `json:"value"`
			DisplayStrings struct {
				Silver string `json:"silver"`
				Copper string `json:"copper"`
				Header string `json:"header"`
				Gold   string `json:"gold"`
			} `json:"display_strings"`
		} `json:"sell_price"`
		Requirements struct {
			Level struct {
				Value         float64 `json:"value"`
				DisplayString string  `json:"display_string"`
			} `json:"level"`
			PlayableClasses struct {
				Links []struct {
					Name string  `json:"name"`
					Id   float64 `json:"id"`
					Key  struct {
						Href string `json:"href"`
					} `json:"key"`
				} `json:"links"`
				DisplayString string `json:"display_string"`
			} `json:"playable_classes"`
		} `json:"requirements"`
		Level struct {
			Value         float64 `json:"value"`
			DisplayString string  `json:"display_string"`
		} `json:"level"`
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
			Item struct {
				Name string  `json:"name"`
				Id   float64 `json:"id"`
				Key  struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"item"`
		} `json:"items"`
		Effects []struct {
			DisplayString string  `json:"display_string"`
			RequiredCount float64 `json:"required_count"`
		} `json:"effects"`
		DisplayString string `json:"display_string"`
	} `json:"equipped_item_sets"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Name  string  `json:"name"`
		Id    float64 `json:"id"`
		Realm struct {
			Name string  `json:"name"`
			Id   float64 `json:"id"`
			Slug string  `json:"slug"`
			Key  struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"realm"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"character"`
}
