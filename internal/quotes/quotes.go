/*
Copyright © 2025 Chris Griffis <dev@chrisgriffis.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package quotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Quote              string `json:"quote,omitempty"`
	Character          string `json:"character,omitempty"`
	Image              string `json:"image,omitempty"`
	CharacterDirection string `json:"characterDirection,omitempty"`
}

type Service struct {
	baseUrl string
}

func NewQuoteService() *Service {
	return &Service{
		baseUrl: "https://thesimpsonsquoteapi.glitch.me/quotes",
	}
}

func (q *Service) GetQuote() ([]Quote, error) {
	req, err := http.NewRequest("GET", q.baseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var quotes []Quote
	err = json.NewDecoder(resp.Body).Decode(&quotes)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return quotes, nil
}
