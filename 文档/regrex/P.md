#### 正则\P{name}

```text
\p{L} or \p{Letter}: any kind of letter from any language.
\p{Ll} or \p{Lowercase_Letter}: a lowercase letter that has an uppercase variant.
\p{Lu} or \p{Uppercase_Letter}: an uppercase letter that has a lowercase variant.
\p{Lt} or \p{Titlecase_Letter}: a letter that appears at the start of a word when only the first letter of the word is capitalized.
\p{L&} or \p{Letter&}: a letter that exists in lowercase and uppercase variants (combination of Ll, Lu and Lt).
\p{Lm} or \p{Modifier_Letter}: a special character that is used like a letter.
\p{Lo} or \p{Other_Letter}: a letter or ideograph that does not have lowercase and uppercase variants.
\p{M} or \p{Mark}: a character intended to be combined with another character (e.g. accents, umlauts, enclosing boxes, etc.).
\p{Mn} or \p{Non_Spacing_Mark}: a character intended to be combined with another character without taking up extra space (e.g. accents, umlauts, etc.).
\p{Mc} or \p{Spacing_Combining_Mark}: a character intended to be combined with another character that takes up extra space (vowel signs in many Eastern languages).
\p{Me} or \p{Enclosing_Mark}: a character that encloses the character is is combined with (circle, square, keycap, etc.).
\p{Z} or \p{Separator}: any kind of whitespace or invisible separator.
\p{Zs} or \p{Space_Separator}: a whitespace character that is invisible, but does take up space.
\p{Zl} or \p{Line_Separator}: line separator character U+2028.
\p{Zp} or \p{Paragraph_Separator}: paragraph separator character U+2029.
\p{S} or \p{Symbol}: math symbols, currency signs, dingbats, box-drawing characters, etc..
\p{Sm} or \p{Math_Symbol}: any mathematical symbol.
\p{Sc} or \p{Currency_Symbol}: any currency sign.
\p{Sk} or \p{Modifier_Symbol}: a combining character (mark) as a full character on its own.
\p{So} or \p{Other_Symbol}: various symbols that are not math symbols, currency signs, or combining characters.
\p{N} or \p{Number}: any kind of numeric character in any script.
\p{Nd} or \p{Decimal_Digit_Number}: a digit zero through nine in any script except ideographic scripts.
\p{Nl} or \p{Letter_Number}: a number that looks like a letter, such as a Roman numeral.
\p{No} or \p{Other_Number}: a superscript or subscript digit, or a number that is not a digit 0..9 (excluding numbers from ideographic scripts).
\p{P} or \p{Punctuation}: any kind of punctuation character.
\p{Pd} or \p{Dash_Punctuation}: any kind of hyphen or dash.
\p{Ps} or \p{Open_Punctuation}: any kind of opening bracket.
\p{Pe} or \p{Close_Punctuation}: any kind of closing bracket.
\p{Pi} or \p{Initial_Punctuation}: any kind of opening quote.
\p{Pf} or \p{Final_Punctuation}: any kind of closing quote.
\p{Pc} or \p{Connector_Punctuation}: a punctuation character such as an underscore that connects words.
\p{Po} or \p{Other_Punctuation}: any kind of punctuation character that is not a dash, bracket, quote or connector.
\p{C} or \p{Other}: invisible control characters and unused code points.
\p{Cc} or \p{Control}: an ASCII 0x00..0x1F or Latin-1 0x80..0x9F control character.
\p{Cf} or \p{Format}: invisible formatting indicator.
\p{Co} or \p{Private_Use}: any code point reserved for private use.
\p{Cs} or \p{Surrogate}: one half of a surrogate pair in UTF-16 encoding.
\p{Cn} or \p{Unassigned}: any code point to which no character has been assigned.
```

##### 匹配语言

```text
\p{Common}
\p{Arabic}
\p{Armenian}
\p{Bengali}
\p{Bopomofo}
\p{Braille}
\p{Buhid}
\p{CanadianAboriginal}
\p{Cherokee}
\p{Cyrillic}
\p{Devanagari}
\p{Ethiopic}
\p{Georgian}
\p{Greek}
\p{Gujarati}
\p{Gurmukhi}
\p{Han}   //匹配汉字
\p{Hangul}
\p{Hanunoo}
\p{Hebrew}
\p{Hiragana}
\p{Inherited}
\p{Kannada}
\p{Katakana}
\p{Khmer}
\p{Lao}
\p{Latin}
\p{Limbu}
\p{Malayalam}
\p{Mongolian}
\p{Myanmar}
\p{Ogham}
\p{Oriya}
\p{Runic}
\p{Sinhala}
\p{Syriac}
\p{Tagalog}
\p{Tagbanwa}
\p{TaiLe}
\p{Tamil}
\p{Telugu}
\p{Thaana}
\p{Thai}
\p{Tibetan}
\p{Yi}
```