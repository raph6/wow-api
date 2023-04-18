package wowapi

type Equipment struct {
	_links struct {
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
	Equipped_items []struct {
		Sockets []struct {
			Socket_type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"socket_type"`
			Item struct {
				Id  float64 `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"item"`
			Display_string string `json:"display_string"`
			Media          struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Id float64 `json:"id"`
			} `json:"media"`
		} `json:"sockets"`
		Quality struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"quality"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"media"`
		Armor struct {
			Value   float64 `json:"value"`
			Display struct {
				Display_string string `json:"display_string"`
				Color          struct {
					A float64 `json:"a"`
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
				} `json:"color"`
			} `json:"display"`
		} `json:"armor"`
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id float64 `json:"id"`
		} `json:"item"`
		Inventory_type struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Binding struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		Sell_price struct {
			Value           float64 `json:"value"`
			Display_strings struct {
				Header string `json:"header"`
				Gold   string `json:"gold"`
				Silver string `json:"silver"`
				Copper string `json:"copper"`
			} `json:"display_strings"`
		} `json:"sell_price"`
		Set struct {
			Item_set struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string  `json:"name"`
				Id   float64 `json:"id"`
			} `json:"item_set"`
			Items []struct {
				Is_equipped bool `json:"is_equipped"`
				Item        struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"item"`
			} `json:"items"`
			Effects []struct {
				Display_string string  `json:"display_string"`
				Required_count float64 `json:"required_count"`
				Is_active      bool    `json:"is_active"`
			} `json:"effects"`
			Display_string string `json:"display_string"`
		} `json:"set"`
		Transmog struct {
			Item struct {
				Id  float64 `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
			} `json:"item"`
			Display_string              string  `json:"display_string"`
			Item_modified_appearance_id float64 `json:"item_modified_appearance_id"`
		} `json:"transmog"`
		Name_description struct {
			Display_string string `json:"display_string"`
			Color          struct {
				R float64 `json:"r"`
				G float64 `json:"g"`
				B float64 `json:"b"`
				A float64 `json:"a"`
			} `json:"color"`
		} `json:"name_description"`
		Item_subclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_subclass"`
		Quantity               float64 `json:"quantity"`
		Context                float64 `json:"context"`
		Name                   string  `json:"name"`
		Modified_appearance_id float64 `json:"modified_appearance_id"`
		Item_class             struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string  `json:"name"`
			Id   float64 `json:"id"`
		} `json:"item_class"`
		Requirements struct {
			Level struct {
				Display_string string  `json:"display_string"`
				Value          float64 `json:"value"`
			} `json:"level"`
			Playable_classes struct {
				Links []struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string  `json:"name"`
					Id   float64 `json:"id"`
				} `json:"links"`
				Display_string string `json:"display_string"`
			} `json:"playable_classes"`
		} `json:"requirements"`
		Level struct {
			Value          float64 `json:"value"`
			Display_string string  `json:"display_string"`
		} `json:"level"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value   float64 `json:"value"`
			Display struct {
				Display_string string `json:"display_string"`
				Color          struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
			} `json:"display"`
		} `json:"stats"`
		Durability struct {
			Value          float64 `json:"value"`
			Display_string string  `json:"display_string"`
		} `json:"durability"`
		Bonus_list []struct {
			Bonus_list float64 `json:"bonus_list"`
		} `json:"bonus_list"`
	} `json:"equipped_items"`
	Equipped_item_sets []struct {
		Display_string string `json:"display_string"`
		Item_set       struct {
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
		} `json:"items"`
		Effects []struct {
			Required_count float64 `json:"required_count"`
			Display_string string  `json:"display_string"`
		} `json:"effects"`
	} `json:"equipped_item_sets"`
}
