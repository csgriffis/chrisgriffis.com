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

package controller

import (
	"fmt"

	"chrisgriffis.com/internal/quotes"
)

type Quote struct {
	quote *quotes.Service
}

func MustQuote(q *Quote, err error) *Quote {
	if err != nil {
		panic(fmt.Errorf("error creating quote: %w", err))
	}

	return q
}

func NewQuote(q *quotes.Service) (*Quote, error) {
	if q == nil {
		return nil, fmt.Errorf("quotes service required")
	}

	return &Quote{
		quote: q,
	}, nil
}

func (q *Quote) GetQuote() (quotes.Quote, error) {
	res, err := q.quote.GetQuote()
	if err != nil {
		return quotes.Quote{}, fmt.Errorf("[quote.GetQuote] %w", err)
	}

	if len(res) == 0 {
		return quotes.Quote{}, fmt.Errorf("no quotes returned")
	}

	return res[0], nil
}
