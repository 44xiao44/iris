package macro

import (
	"strconv"
	"strings"

	"github.com/kataras/iris/core/router/macro/interpreter/ast"
)

var (
	// String type
	// Allows anything (single path segment, as everything except the `Path`).
	String = NewMacro("string", "", true, false, func(string) bool { return true }).
		RegisterFunc("regexp", func(expr string) EvaluatorFunc {
			return MustNewEvaluatorFromRegexp(expr)
		}).
		// checks if param value starts with the 'prefix' arg
		RegisterFunc("prefix", func(prefix string) EvaluatorFunc {
			return func(paramValue string) bool {
				return strings.HasPrefix(paramValue, prefix)
			}
		}).
		// checks if param value ends with the 'suffix' arg
		RegisterFunc("suffix", func(suffix string) EvaluatorFunc {
			return func(paramValue string) bool {
				return strings.HasSuffix(paramValue, suffix)
			}
		}).
		// checks if param value contains the 's' arg
		RegisterFunc("contains", func(s string) EvaluatorFunc {
			return func(paramValue string) bool {
				return strings.Contains(paramValue, s)
			}
		}).
		// checks if param value's length is at least 'min'
		RegisterFunc("min", func(min int) EvaluatorFunc {
			return func(paramValue string) bool {
				return len(paramValue) >= min
			}
		}).
		// checks if param value's length is not bigger than 'max'
		RegisterFunc("max", func(max int) EvaluatorFunc {
			return func(paramValue string) bool {
				return max >= len(paramValue)
			}
		})

	simpleNumberEvalutator = MustNewEvaluatorFromRegexp("^-?[0-9]+$")
	// Number or int type
	// both positive and negative numbers, any number of digits.
	Number = NewMacro("number", "int", false, false, simpleNumberEvalutator).
		// checks if the param value's int representation is
		// bigger or equal than 'min'
		RegisterFunc("min", func(min int) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.Atoi(paramValue)
				if err != nil {
					return false
				}
				return n >= min
			}
		}).
		// checks if the param value's int representation is
		// smaller or equal than 'max'.
		RegisterFunc("max", func(max int) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.Atoi(paramValue)
				if err != nil {
					return false
				}
				return n <= max
			}
		}).
		// checks if the param value's int representation is
		// between min and max, including 'min' and 'max'.
		RegisterFunc("range", func(min, max int) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.Atoi(paramValue)
				if err != nil {
					return false
				}

				if n < min || n > max {
					return false
				}
				return true
			}
		})

	// Int64 as int64 type
	// -9223372036854775808 to 9223372036854775807.
	Int64 = NewMacro("int64", "long", false, false, func(paramValue string) bool {
		if !simpleNumberEvalutator(paramValue) {
			return false
		}
		_, err := strconv.ParseInt(paramValue, 10, 64)
		// if err == strconv.ErrRange...
		return err == nil
	}).
		// checks if the param value's int64 representation is
		// bigger or equal than 'min'.
		RegisterFunc("min", func(min int64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseInt(paramValue, 10, 64)
				if err != nil {
					return false
				}
				return n >= min
			}
		}).
		// checks if the param value's int64 representation is
		// smaller or equal than 'max'.
		RegisterFunc("max", func(max int64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseInt(paramValue, 10, 64)
				if err != nil {
					return false
				}
				return n <= max
			}
		}).
		// checks if the param value's int64 representation is
		// between min and max, including 'min' and 'max'.
		RegisterFunc("range", func(min, max int64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseInt(paramValue, 10, 64)
				if err != nil {
					return false
				}

				if n < min || n > max {
					return false
				}
				return true
			}
		})

	// Uint8 as uint8 type
	// 0 to 255.
	Uint8 = NewMacro("uint8", "", false, false, MustNewEvaluatorFromRegexp("^([0-9]|[1-8][0-9]|9[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")).
		// checks if the param value's uint8 representation is
		// bigger or equal than 'min'.
		RegisterFunc("min", func(min uint8) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 8)
				if err != nil {
					return false
				}

				return uint8(n) >= min
			}
		}).
		// checks if the param value's uint8 representation is
		// smaller or equal than 'max'.
		RegisterFunc("max", func(max uint8) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 8)
				if err != nil {
					return false
				}
				return uint8(n) <= max
			}
		}).
		// checks if the param value's uint8 representation is
		// between min and max, including 'min' and 'max'.
		RegisterFunc("range", func(min, max uint8) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 8)
				if err != nil {
					return false
				}

				if v := uint8(n); v < min || v > max {
					return false
				}
				return true
			}
		})

	// Uint64 as uint64 type
	// 0 to 18446744073709551615.
	Uint64 = NewMacro("uint64", "", false, false, func(paramValue string) bool {
		if !simpleNumberEvalutator(paramValue) {
			return false
		}
		_, err := strconv.ParseUint(paramValue, 10, 64)
		return err == nil
	}).
		// checks if the param value's uint64 representation is
		// bigger or equal than 'min'.
		RegisterFunc("min", func(min uint64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 64)
				if err != nil {
					return false
				}
				return n >= min
			}
		}).
		// checks if the param value's uint64 representation is
		// smaller or equal than 'max'.
		RegisterFunc("max", func(max uint64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 64)
				if err != nil {
					return false
				}
				return n <= max
			}
		}).
		// checks if the param value's uint64 representation is
		// between min and max, including 'min' and 'max'.
		RegisterFunc("range", func(min, max uint64) EvaluatorFunc {
			return func(paramValue string) bool {
				n, err := strconv.ParseUint(paramValue, 10, 64)
				if err != nil {
					return false
				}

				if n < min || n > max {
					return false
				}
				return true
			}
		})

	// Bool or boolean as bool type
	// a string which is "1" or "t" or "T" or "TRUE" or "true" or "True"
	// or "0" or "f" or "F" or "FALSE" or "false" or "False".
	Bool = NewMacro("bool", "boolean", false, false, func(paramValue string) bool {
		// a simple if statement is faster than regex ^(true|false|True|False|t|0|f|FALSE|TRUE)$
		// in this case.
		_, err := strconv.ParseBool(paramValue)
		return err == nil
	})

	// Alphabetical letter type
	// letters only (upper or lowercase)
	Alphabetical = NewMacro("alphabetical", "", false, false, MustNewEvaluatorFromRegexp("^[a-zA-Z ]+$"))
	// File type
	// letters (upper or lowercase)
	// numbers (0-9)
	// underscore (_)
	// dash (-)
	// point (.)
	// no spaces! or other character
	File = NewMacro("file", "", false, false, MustNewEvaluatorFromRegexp("^[a-zA-Z0-9_.-]*$"))
	// Path type
	// anything, should be the last part
	//
	// It allows everything, we have String and Path as different
	// types because I want to give the opportunity to the user
	// to organise the macro functions based on wildcard or single dynamic named path parameter.
	// Should be living in the latest path segment of a route path.
	Path = NewMacro("path", "", false, true, func(string) bool { return true })

	Defaults = &Macros{
		String,
		Number,
		Int64,
		Uint8,
		Uint64,
		Bool,
		Alphabetical,
		Path,
	}
)

type Macros []*Macro

func (ms *Macros) Register(indent, alias string, isMaster, isTrailing bool, evaluator EvaluatorFunc) *Macro {
	macro := NewMacro(indent, alias, isMaster, isTrailing, evaluator)
	if ms.register(macro) {
		return macro
	}
	return nil
}

func (ms *Macros) register(macro *Macro) bool {
	if macro.Indent() == "" || macro.Evaluator == nil {
		return false
	}

	cp := *ms

	for _, m := range cp {

		// can't add more than one with the same ast characteristics.
		if macro.Indent() == m.Indent() {
			return false
		}

		if macro.Alias() == m.Alias() || macro.Alias() == m.Indent() {
			return false
		}

		if macro.Master() && m.Master() {
			return false
		}
	}

	cp = append(cp, macro)

	*ms = cp
	return true
}

func (ms *Macros) Unregister(indent string) bool {
	cp := *ms

	for i, m := range cp {
		if m.Indent() == indent {
			copy(cp[i:], cp[i+1:])
			cp[len(cp)-1] = nil
			cp = cp[:len(cp)-1]

			*ms = cp
			return true
		}
	}

	return false
}

func (ms *Macros) Lookup(pt ast.ParamType) *Macro {
	if m := ms.Get(pt.Indent()); m != nil {
		return m
	}

	if alias, has := ast.HasAlias(pt); has {
		if m := ms.Get(alias); m != nil {
			return m
		}
	}

	return nil
}

func (ms *Macros) Get(indentOrAlias string) *Macro {
	if indentOrAlias == "" {
		return nil
	}

	for _, m := range *ms {
		if m.Indent() == indentOrAlias {
			return m
		}

		if m.Alias() == indentOrAlias {
			return m
		}
	}

	return nil
}

func (ms *Macros) GetMaster() *Macro {
	for _, m := range *ms {
		if m.Master() {
			return m
		}
	}

	return nil
}

func (ms *Macros) GetTrailings() (macros []*Macro) {
	for _, m := range *ms {
		if m.Trailing() {
			macros = append(macros, m)
		}
	}

	return
}