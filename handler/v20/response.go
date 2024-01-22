// Copyright 2023 Martin Zimandl <martin.zimandl@gmail.com>
// Copyright 2023 Institute of the Czech National Corpus,
//                Faculty of Arts, Charles University
//   This file is part of MQUERY.
//
//  MQUERY is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  MQUERY is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with MQUERY.  If not, see <https://www.gnu.org/licenses/>.

package v20

import (
	"fcs/corpus"
	"fcs/general"
)

type FCSResourceInfo struct {
	PID         string
	Title       map[string]string
	Description map[string]string
	URI         string
	Languages   []string

	// AvailableLayers is a list of values separated
	// by single space (as required in CLARIN-FCS
	// Interface Specification)
	AvailableLayers string
}

type Segment struct {
	ID    string
	Start int
	End   int
}

type Token struct {
	Text    string
	Hit     bool
	Segment Segment
	Layers  map[corpus.LayerType]string
}

type FCSSearchRow struct {
	Position   int
	PID        string
	Web        string
	Ref        string
	LayerAttrs []corpus.LayerType
	Tokens     []Token
}

type FCSExplain struct {
	ServerName          string
	ServerPort          string
	Database            string
	DatabaseTitle       string
	DatabaseDescription string
	PosAttrs            []corpus.PosAttr
	Resources           []FCSResourceInfo
	ExtraResponseData   bool
}

type FCSSearchRetrieve struct {
	QueryType QueryType
	Results   []FCSSearchRow
}

type FCSResponse struct {
	General       general.FCSGeneralResponse
	RecordPacking RecordPacking
	Operation     Operation

	Explain        FCSExplain
	SearchRetrieve FCSSearchRetrieve
}
