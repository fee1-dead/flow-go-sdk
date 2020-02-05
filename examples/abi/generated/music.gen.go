// Code generated by Flow Go SDK. DO NOT EDIT.

package generated

import (
	"bytes"
	"fmt"
	"github.com/dapperlabs/flow-go/language"
	"github.com/dapperlabs/flow-go/language/encoding"
)

type AlbumView interface {
	Artist() ArtistView
	Name() string
	Rating() *uint8
	Year() uint16
}
type albumView struct {
	_artist ArtistView
	_name   string
	_rating *uint8
	_year   uint16
}

func (t *albumView) Artist() ArtistView {
	return t._artist
}
func (t *albumView) Name() string {
	return t._name
}
func (t *albumView) Rating() *uint8 {
	return t._rating
}
func (t *albumView) Year() uint16 {
	return t._year
}
func AlbumViewfromValue(value language.Value) (AlbumView, error) {
	composite, err := language.CastToComposite(value)
	if err != nil {
		return nil, err
	}

	_artist, err := ArtistViewfromValue(composite.Fields[0])
	if err != nil {
		return nil, err
	}

	_name, err := language.CastToString(composite.Fields[1])
	if err != nil {
		return nil, err
	}

	_rating, err := __converter0(composite.Fields[2])
	if err != nil {
		return nil, err
	}

	_year, err := language.CastToUInt16(composite.Fields[3])
	if err != nil {
		return nil, err
	}

	return &albumView{
		_artist: _artist,
		_name:   _name,
		_rating: _rating,
		_year:   _year,
	}, nil
}
func DecodeAlbumView(b []byte) (AlbumView, error) {
	r := bytes.NewReader(b)
	dec := encoding.NewDecoder(r)
	v, err := dec.DecodeComposite(albumType)
	if err != nil {
		return nil, err
	}

	return AlbumViewfromValue(v)
}
func DecodeAlbumViewVariableSizedArray(b []byte) ([]AlbumView, error) {
	r := bytes.NewReader(b)
	dec := encoding.NewDecoder(r)
	v, err := dec.DecodeVariableSizedArray(language.VariableSizedArrayType{ElementType: albumType})
	if err != nil {
		return nil, err
	}

	array := make([]AlbumView, len(v.Values))
	for i, t := range v.Values {
		array[i], err = AlbumViewfromValue(t.(language.Composite))
		if err != nil {
			return nil, err
		}

	}
	return array, nil
}

var albumType = language.CompositeType{
	Fields: []language.Field{language.Field{
		Identifier: "artist",
		Type:       artistType,
	}, language.Field{
		Identifier: "name",
		Type:       language.StringType{},
	}, language.Field{
		Identifier: "rating",
		Type:       language.OptionalType{Type: language.UInt8Type{}},
	}, language.Field{
		Identifier: "year",
		Type:       language.UInt16Type{},
	}},
	Identifier: "Album",
	Initializers: [][]language.Parameter{{language.Parameter{
		Identifier: "artist",
		Label:      "",
		Type:       artistType,
	}, language.Parameter{
		Identifier: "name",
		Label:      "",
		Type:       language.StringType{},
	}, language.Parameter{
		Identifier: "year",
		Label:      "",
		Type:       language.UInt16Type{},
	}, language.Parameter{
		Identifier: "rating",
		Label:      "",
		Type:       language.OptionalType{Type: language.UInt8Type{}},
	}}},
}

type AlbumConstructor interface {
	Encode() ([]byte, error)
}
type albumConstructor struct {
	artist ArtistView
	name   string
	year   uint16
	rating *uint8
}

func (p albumConstructor) toValue() language.ConstantSizedArray {
	return language.NewConstantSizedArray([]language.Value{language.MustConvertValue(p.artist), language.MustConvertValue(p.name), language.MustConvertValue(p.year), language.MustConvertValue(p.rating)})
}
func (p albumConstructor) Encode() ([]byte, error) {
	var w bytes.Buffer
	encoder := encoding.NewEncoder(&w)
	err := encoder.EncodeConstantSizedArray(p.toValue())
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
func NewAlbumConstructor(artist ArtistView, name string, year uint16, rating *uint8) (AlbumConstructor, error) {
	return albumConstructor{
		artist: artist,
		name:   name,
		rating: rating,
		year:   year,
	}, nil
}

type ArtistView interface {
	Country() string
	Members() *[]string
	Name() string
}
type artistView struct {
	_country string
	_members *[]string
	_name    string
}

func (t *artistView) Country() string {
	return t._country
}
func (t *artistView) Members() *[]string {
	return t._members
}
func (t *artistView) Name() string {
	return t._name
}
func ArtistViewfromValue(value language.Value) (ArtistView, error) {
	composite, err := language.CastToComposite(value)
	if err != nil {
		return nil, err
	}

	_country, err := language.CastToString(composite.Fields[0])
	if err != nil {
		return nil, err
	}

	_members, err := __converter1(composite.Fields[1])
	if err != nil {
		return nil, err
	}

	_name, err := language.CastToString(composite.Fields[2])
	if err != nil {
		return nil, err
	}

	return &artistView{
		_country: _country,
		_members: _members,
		_name:    _name,
	}, nil
}
func DecodeArtistView(b []byte) (ArtistView, error) {
	r := bytes.NewReader(b)
	dec := encoding.NewDecoder(r)
	v, err := dec.DecodeComposite(artistType)
	if err != nil {
		return nil, err
	}

	return ArtistViewfromValue(v)
}
func DecodeArtistViewVariableSizedArray(b []byte) ([]ArtistView, error) {
	r := bytes.NewReader(b)
	dec := encoding.NewDecoder(r)
	v, err := dec.DecodeVariableSizedArray(language.VariableSizedArrayType{ElementType: artistType})
	if err != nil {
		return nil, err
	}

	array := make([]ArtistView, len(v.Values))
	for i, t := range v.Values {
		array[i], err = ArtistViewfromValue(t.(language.Composite))
		if err != nil {
			return nil, err
		}

	}
	return array, nil
}

var artistType = language.CompositeType{
	Fields: []language.Field{language.Field{
		Identifier: "country",
		Type:       language.StringType{},
	}, language.Field{
		Identifier: "members",
		Type:       language.OptionalType{Type: language.VariableSizedArrayType{ElementType: language.StringType{}}},
	}, language.Field{
		Identifier: "name",
		Type:       language.StringType{},
	}},
	Identifier: "Artist",
	Initializers: [][]language.Parameter{{language.Parameter{
		Identifier: "name",
		Label:      "",
		Type:       language.StringType{},
	}, language.Parameter{
		Identifier: "members",
		Label:      "",
		Type:       language.OptionalType{Type: language.VariableSizedArrayType{ElementType: language.StringType{}}},
	}, language.Parameter{
		Identifier: "country",
		Label:      "",
		Type:       language.StringType{},
	}}},
}

type ArtistConstructor interface {
	Encode() ([]byte, error)
}
type artistConstructor struct {
	name    string
	members *[]string
	country string
}

func (p artistConstructor) toValue() language.ConstantSizedArray {
	return language.NewConstantSizedArray([]language.Value{language.MustConvertValue(p.name), language.MustConvertValue(p.members), language.MustConvertValue(p.country)})
}
func (p artistConstructor) Encode() ([]byte, error) {
	var w bytes.Buffer
	encoder := encoding.NewEncoder(&w)
	err := encoder.EncodeConstantSizedArray(p.toValue())
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
func NewArtistConstructor(name string, members *[]string, country string) (ArtistConstructor, error) {
	return artistConstructor{
		country: country,
		members: members,
		name:    name,
	}, nil
}
func __converter0(p language.Value) (*uint8, error) {
	var ret0 uint8
	var go0 interface{}
	cast0, ok := p.(language.Optional)
	if !ok {
		return nil, fmt.Errorf("cannot cast %T", p)

	}
	go0 = cast0.ToGoValue()

	var err error
	if go0 == nil {
		return nil, nil
	} else {
		cast1, ok := cast0.Value.(language.UInt8)
		if !ok {
			return nil, fmt.Errorf("cannot cast %T", cast0.Value)

		}

		ret0, err = language.CastToUInt8(cast1)
		if err != nil {
			return nil, err
		}

	}
	return &ret0, nil

}
func __converter1(p language.Value) (*[]string, error) {
	var ret0 []string
	var go0 interface{}
	cast0, ok := p.(language.Optional)
	if !ok {
		return nil, fmt.Errorf("cannot cast %T", p)

	}
	go0 = cast0.ToGoValue()

	var err error
	if go0 == nil {
		return nil, nil
	} else {
		var ret1 []string
		cast1, ok := cast0.Value.(language.VariableSizedArray)
		if !ok {
			return nil, fmt.Errorf("cannot cast %T", cast0.Value)

		}

		if err != nil {
			return nil, err
		}
		ret1 = make([]string, len(cast1.Values))
		for i1, elem1 := range cast1.Values {
			cast2, ok := elem1.(language.String)
			if !ok {
				return nil, fmt.Errorf("cannot cast %T", elem1)

			}

			ret1[i1], err = language.CastToString(cast2)
			if err != nil {
				return nil, err
			}

		}
		ret0 = ret1

	}
	return &ret0, nil

}
