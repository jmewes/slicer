// Copyright (C) 2026 Jan Mewes
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package specantlr contains the ANTLR-generated lexer/parser for the minimal
// TypeScript spec-file island grammar used to reverse-engineer Gherkin
// feature files from Jasmine-style *.spec.ts sources.
//
// The generated `.go` files are committed to the repository; end users only
// need Go to build the project. Regenerate them with:
//
//	go generate ./core/antlr/...
//
// This requires Java and the ANTLR jar (see README / AGENTS.md for setup).
package specantlr

//go:generate antlr -Dlanguage=Go -no-listener -visitor -package specantlr -o . SpecGrammar.g4
