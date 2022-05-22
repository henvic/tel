package tel

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

// Encoder is a mechanism for serializing an attribute set into a specific
// string representation that supports caching, to avoid repeated
// serialization. An example could be an exporter encoding the attribute
// set into a wire representation.
type Encoder = attribute.Encoder

// EncoderID is used to identify distinct Encoder
// implementations, for caching encoded results.
type EncoderID = attribute.EncoderID

// NewEncoderID returns a unique attribute encoder ID. It should be called
// once per each type of attribute encoder. Preferably in init() or in var
// definition.
func NewEncoderID() EncoderID {
	return attribute.NewEncoderID()
}

// DefaultEncoder returns an attribute encoder that encodes attributes in such
// a way that each escaped attribute's key is followed by an equal sign and
// then by an escaped attribute's value. All key-value pairs are separated by
// a comma.
//
// Escaping is done by prepending a backslash before either a backslash, equal
// sign or a comma.
func DefaultEncoder() Encoder {
	return attribute.DefaultEncoder()
}

// AttributeIterator allows iterating over the set of attributes in order, sorted by
// key.
type AttributeIterator = attribute.Iterator

// AttributeMergeIterator supports iterating over two sets of attributes while
// eliminating duplicate values from the combined set. The first iterator
// value takes precedence.
type AttributeMergeIterator = attribute.MergeIterator

// AttributeNewMergeIterator returns a MergeIterator for merging two attribute sets.
// Duplicates are resolved by taking the value from the first set.
func AttributeNewMergeIterator(s1, s2 *Set) AttributeMergeIterator {
	return attribute.NewMergeIterator(s1, s2)
}

// Key represents the key part in key-value pairs. It's a string. The
// allowed character set in the key depends on the use of the key.
type Key = attribute.Key

// KeyValue holds a key and value pair.
type KeyValue = attribute.KeyValue

// AttributeBool creates a KeyValue with a BOOL Value type.
func AttributeBool(k string, v bool) KeyValue {
	return attribute.Key(k).Bool(v)
}

// AttributeBoolSlice creates a KeyValue with a BOOLSLICE Value type.
func AttributeBoolSlice(k string, v []bool) KeyValue {
	return attribute.Key(k).BoolSlice(v)
}

// AttributeInt creates a KeyValue with an INT64 Value type.
func AttributeInt(k string, v int) KeyValue {
	return attribute.Key(k).Int(v)
}

// AttributeIntSlice creates a KeyValue with an INT64SLICE Value type.
func AttributeIntSlice(k string, v []int) KeyValue {
	return attribute.Key(k).IntSlice(v)
}

// AttributeInt64 creates a KeyValue with an INT64 Value type.
func AttributeInt64(k string, v int64) KeyValue {
	return attribute.Key(k).Int64(v)
}

// AttributeInt64Slice creates a KeyValue with an INT64SLICE Value type.
func AttributeInt64Slice(k string, v []int64) KeyValue {
	return attribute.Key(k).Int64Slice(v)
}

// AttributeFloat64 creates a KeyValue with a FLOAT64 Value type.
func AttributeFloat64(k string, v float64) KeyValue {
	return attribute.Key(k).Float64(v)
}

// AttributeFloat64Slice creates a KeyValue with a FLOAT64SLICE Value type.
func AttributeFloat64Slice(k string, v []float64) KeyValue {
	return attribute.Key(k).Float64Slice(v)
}

// AttributeString creates a KeyValue with a STRING Value type.
func AttributeString(k, v string) KeyValue {
	return attribute.Key(k).String(v)
}

// AttributeStringSlice creates a KeyValue with a STRINGSLICE Value type.
func AttributeStringSlice(k string, v []string) KeyValue {
	return attribute.Key(k).StringSlice(v)
}

// Stringer creates a new key-value pair with a passed name and a string
// Attributevalue generated by the passed Stringer interface.
func AttributeStringer(k string, v fmt.Stringer) KeyValue {
	return attribute.Key(k).String(v.String())
}

// Set is the representation for a distinct attribute set. It manages an
// immutable set of attributes, with an internal cache for storing
// attribute encodings.
//
// This type supports the Equivalent method of comparison using values of
// type Distinct.
type Set = attribute.Set

// Distinct wraps a variable-size array of KeyValue, constructed with keys
// in sorted order. This can be used as a map key or for equality checking
// between Sets.
type Distinct = attribute.Distinct

// Filter supports removing certain attributes from attribute sets. When
// the filter returns true, the attribute will be kept in the filtered
// attribute set. When the filter returns false, the attribute is excluded
// from the filtered attribute set, and the attribute instead appears in
// the removed list of excluded attributes.
type Filter = attribute.Filter

// Sortable implements sort.Interface, used for sorting KeyValue. This is
// an exported type to support a memory optimization. A pointer to one of
// these is needed for the call to sort.Stable(), which the caller may
// provide in order to avoid an allocation. See NewSetWithSortable().
type Sortable = attribute.Sortable

// EmptySet returns a reference to a Set with no elements.
//
// This is a convenience provided for optimized calling utility.
func EmptySet() *Set {
	return attribute.EmptySet()
}

// NewSet returns a new Set. See the documentation for
// NewSetWithSortableFiltered for more details.
//
// Except for empty sets, this method adds an additional allocation compared
// with calls that include a Sortable.
func NewSet(kvs ...KeyValue) Set {
	return attribute.NewSet(kvs...)
}

// NewSetWithSortable returns a new Set. See the documentation for
// NewSetWithSortableFiltered for more details.
//
// This call includes a Sortable option as a memory optimization.
func NewSetWithSortable(kvs []KeyValue, tmp *Sortable) Set {
	return attribute.NewSetWithSortable(kvs, tmp)
}

// NewSetWithFiltered returns a new Set. See the documentation for
// NewSetWithSortableFiltered for more details.
//
// This call includes a Filter to include/exclude attribute keys from the
// return value. Excluded keys are returned as a slice of attribute values.
func NewSetWithFiltered(kvs []KeyValue, filter Filter) (Set, []KeyValue) {
	return attribute.NewSetWithFiltered(kvs, filter)
}

// NewSetWithSortableFiltered returns a new Set.
//
// Duplicate keys are eliminated by taking the last value.  This
// re-orders the input slice so that unique last-values are contiguous
// at the end of the slice.
//
// This ensures the following:
//
// - Last-value-wins semantics
// - Caller sees the reordering, but doesn't lose values
// - Repeated call preserve last-value wins.
//
// Note that methods are defined on Set, although this returns Set. Callers
// can avoid memory allocations by:
//
// - allocating a Sortable for use as a temporary in this method
// - allocating a Set for storing the return value of this constructor.
//
// The result maintains a cache of encoded attributes, by attribute.EncoderID.
// This value should not be copied after its first use.
//
// The second []KeyValue return value is a list of attributes that were
// excluded by the Filter (if non-nil).
func NewSetWithSortableFiltered(kvs []KeyValue, tmp *Sortable, filter Filter) (Set, []KeyValue) {
	return attribute.NewSetWithSortableFiltered(kvs, tmp, filter)
}

// Type describes the type of the data Value holds.
type Type = attribute.Type

// Value represents the value part in key-value pairs.
type Value = attribute.Value

const (
	// INVALID is used for a Value with no value set.
	INVALID = attribute.INVALID
	// BOOL is a boolean Type Value.
	BOOL = attribute.BOOL
	// INT64 is a 64-bit signed integral Type Value.
	INT64 = attribute.INT64
	// FLOAT64 is a 64-bit floating point Type Value.
	FLOAT64 = attribute.FLOAT64
	// STRING is a string Type Value.
	STRING = attribute.STRING
	// BOOLSLICE is a slice of booleans Type Value.
	BOOLSLICE = attribute.BOOLSLICE
	// INT64SLICE is a slice of 64-bit signed integral numbers Type Value.
	INT64SLICE = attribute.INT64SLICE
	// FLOAT64SLICE is a slice of 64-bit floating point numbers Type Value.
	FLOAT64SLICE = attribute.FLOAT64SLICE
	// STRINGSLICE is a slice of strings Type Value.
	STRINGSLICE = attribute.STRINGSLICE
)

// BoolValue creates a BOOL Value.
func BoolValue(v bool) Value {
	return attribute.BoolValue(v)
}

// BoolSliceValue creates a BOOLSLICE Value.
func BoolSliceValue(v []bool) Value {
	return attribute.BoolSliceValue(v)
}

// IntValue creates an INT64 Value.
func IntValue(v int) Value {
	return attribute.IntValue(v)
}

// IntSliceValue creates an INTSLICE Value.
func IntSliceValue(v []int) Value {
	return attribute.IntSliceValue(v)
}

// Int64Value creates an INT64 Value.
func Int64Value(v int64) Value {
	return attribute.Int64Value(v)
}

// Int64SliceValue creates an INT64SLICE Value.
func Int64SliceValue(v []int64) Value {
	return attribute.Int64SliceValue(v)
}

// Float64Value creates a FLOAT64 Value.
func Float64Value(v float64) Value {
	return attribute.Float64Value(v)
}

// Float64SliceValue creates a FLOAT64SLICE Value.
func Float64SliceValue(v []float64) Value {
	return attribute.Float64SliceValue(v)
}

// StringValue creates a STRING Value.
func StringValue(v string) Value {
	return attribute.StringValue(v)
}

// StringSliceValue creates a STRINGSLICE Value.
func StringSliceValue(v []string) Value {
	return attribute.StringSliceValue(v)
}
