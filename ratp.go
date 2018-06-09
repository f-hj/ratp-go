package main

import (
	"fmt"
	"strings"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://wsiv.ratp.fr"

// NewWsivPortType creates an initializes a WsivPortType.
func NewWsivPortType(cli *soap.Client) WsivPortType {
	return &wsivPortType{cli}
}

// WsivPortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type WsivPortType interface {
	// GetDirections was auto-generated from WSDL.
	GetDirections(parameters *GetDirections) (*GetDirectionsResponse, error)

	// GetGeoPoints was auto-generated from WSDL.
	GetGeoPoints(parameters *GetGeoPoints) (*GetGeoPointsResponse, error)

	// GetLines was auto-generated from WSDL.
	GetLines(parameters *GetLines) (*GetLinesResponse, error)

	// GetMission was auto-generated from WSDL.
	GetMission(parameters *GetMission) (*GetMissionResponse, error)

	// GetMissionsFirstLast was auto-generated from WSDL.
	GetMissionsFirstLast(parameters *GetMissionsFirstLast) (*GetMissionsFirstLastResponse, error)

	// GetMissionsFrequency was auto-generated from WSDL.
	GetMissionsFrequency(parameters *GetMissionsFrequency) (*GetMissionsFrequencyResponse, error)

	// GetMissionsNext was auto-generated from WSDL.
	GetMissionsNext(parameters *GetMissionsNext) (*GetMissionsNextResponse, error)

	// GetPerturbations was auto-generated from WSDL.
	GetPerturbations(parameters *GetPerturbations) (*GetPerturbationsResponse, error)

	// GetStations was auto-generated from WSDL.
	GetStations(parameters *GetStations) (*GetStationsResponse, error)

	// GetVersion was auto-generated from WSDL.
	GetVersion() (*GetVersionResponse, error)
}

// Direction was auto-generated from WSDL.
type Direction struct {
	Line            *Line      `xml:"http://wsiv.ratp.fr/xsd line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
	Name            *string    `xml:"http://wsiv.ratp.fr/xsd name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Sens            *string    `xml:"http://wsiv.ratp.fr/xsd sens,omitempty" json:"sens,omitempty" yaml:"sens,omitempty"`
	StationsEndLine []*Station `xml:"http://wsiv.ratp.fr/xsd stationsEndLine,omitempty" json:"stationsEndLine,omitempty" yaml:"stationsEndLine,omitempty"`
}

// GeoPoint was auto-generated from WSDL.
type GeoPoint struct {
	Id         *string  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name       *string  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	NameSuffix *string  `xml:"nameSuffix,omitempty" json:"nameSuffix,omitempty" yaml:"nameSuffix,omitempty"`
	Type       *string  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	X          *float64 `xml:"x,omitempty" json:"x,omitempty" yaml:"x,omitempty"`
	Y          *float64 `xml:"y,omitempty" json:"y,omitempty" yaml:"y,omitempty"`
}

// Itinerary was auto-generated from WSDL.
type Itinerary struct {
	DateEnd          *string    `xml:"dateEnd,omitempty" json:"dateEnd,omitempty" yaml:"dateEnd,omitempty"`
	DateStart        *string    `xml:"dateStart,omitempty" json:"dateStart,omitempty" yaml:"dateStart,omitempty"`
	DurationsTransit []*int     `xml:"durationsTransit,omitempty" json:"durationsTransit,omitempty" yaml:"durationsTransit,omitempty"`
	GeoPointEnd      *GeoPoint  `xml:"geoPointEnd,omitempty" json:"geoPointEnd,omitempty" yaml:"geoPointEnd,omitempty"`
	GeoPointStart    *GeoPoint  `xml:"geoPointStart,omitempty" json:"geoPointStart,omitempty" yaml:"geoPointStart,omitempty"`
	Missions         []*Mission `xml:"missions,omitempty" json:"missions,omitempty" yaml:"missions,omitempty"`
	Tarif            *Tarif     `xml:"tarif,omitempty" json:"tarif,omitempty" yaml:"tarif,omitempty"`
}

// Line was auto-generated from WSDL.
type Line struct {
	Code     *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	CodeStif *string `xml:"codeStif,omitempty" json:"codeStif,omitempty" yaml:"codeStif,omitempty"`
	Id       *string `xml:"http://wsiv.ratp.fr/xsd id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Image    *string `xml:"image,omitempty" json:"image,omitempty" yaml:"image,omitempty"`
	Name     *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Realm    *string `xml:"realm,omitempty" json:"realm,omitempty" yaml:"realm,omitempty"`
	Reseau   *Reseau `xml:"reseau,omitempty" json:"reseau,omitempty" yaml:"reseau,omitempty"`

	ComputedCode string `xml:"computedCode,omitempty" json:"computedCode,omitempty"`
}

func (line *Line) GetComputedCode() string {
	switch *line.Reseau.Code {
	case "metro":
		return "M" + strings.ToUpper(*line.Code)
	case "rer":
		return "R" + *line.Code
	case "tram":
		return "B" + *line.Code
	default:
		return *line.Code
	}
}

// Mission was auto-generated from WSDL.
type Mission struct {
	Code              *string         `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Direction         *Direction      `xml:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	Id                *string         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Line              *Line           `xml:"line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
	Perturbations     []*Perturbation `xml:"perturbations,omitempty" json:"perturbations,omitempty" yaml:"perturbations,omitempty"`
	StationEndLine    *Station        `xml:"stationEndLine,omitempty" json:"stationEndLine,omitempty" yaml:"stationEndLine,omitempty"`
	Stations          []*Station      `xml:"stations,omitempty" json:"stations,omitempty" yaml:"stations,omitempty"`
	StationsDates     []*string       `xml:"stationsDates,omitempty" json:"stationsDates,omitempty" yaml:"stationsDates,omitempty"`
	StationsMessages  []*string       `xml:"stationsMessages,omitempty" json:"stationsMessages,omitempty" yaml:"stationsMessages,omitempty"`
	StationsPlatforms []*string       `xml:"stationsPlatforms,omitempty" json:"stationsPlatforms,omitempty" yaml:"stationsPlatforms,omitempty"`
	StationsStops     []*bool         `xml:"stationsStops,omitempty" json:"stationsStops,omitempty" yaml:"stationsStops,omitempty"`
}

// Perturbation was auto-generated from WSDL.
type Perturbation struct {
	Cause       *PerturbationCause       `xml:"cause,omitempty" json:"cause,omitempty" yaml:"cause,omitempty"`
	Consequence *PerturbationConsequence `xml:"consequence,omitempty" json:"consequence,omitempty" yaml:"consequence,omitempty"`
	DateEnd     *string                  `xml:"dateEnd,omitempty" json:"dateEnd,omitempty" yaml:"dateEnd,omitempty"`
	DateStart   *string                  `xml:"dateStart,omitempty" json:"dateStart,omitempty" yaml:"dateStart,omitempty"`
	Id          *string                  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Incidents   []*PerturbationIncident  `xml:"incidents,omitempty" json:"incidents,omitempty" yaml:"incidents,omitempty"`
	Level       *string                  `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Line        *Line                    `xml:"http://wsiv.ratp.fr/xsd line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
	Media       *string                  `xml:"media,omitempty" json:"media,omitempty" yaml:"media,omitempty"`
	Message     *PerturbationMessage     `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	Source      *string                  `xml:"source,omitempty" json:"source,omitempty" yaml:"source,omitempty"`
	Station     *Station                 `xml:"station,omitempty" json:"station,omitempty" yaml:"station,omitempty"`
	Status      *string                  `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	TimeEnd     *string                  `xml:"timeEnd,omitempty" json:"timeEnd,omitempty" yaml:"timeEnd,omitempty"`
	TimeStart   *string                  `xml:"timeStart,omitempty" json:"timeStart,omitempty" yaml:"timeStart,omitempty"`
	Title       *string                  `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// PerturbationCause was auto-generated from WSDL.
type PerturbationCause struct {
	Id         *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name       *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	ParentId   *string `xml:"parentId,omitempty" json:"parentId,omitempty" yaml:"parentId,omitempty"`
	ParentName *string `xml:"parentName,omitempty" json:"parentName,omitempty" yaml:"parentName,omitempty"`
}

// PerturbationConsequence was auto-generated from WSDL.
type PerturbationConsequence struct {
	Code  *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Id    *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Level *string `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Name  *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// PerturbationIncident was auto-generated from WSDL.
type PerturbationIncident struct {
	Date          *string                     `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Id            *string                     `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IncidentLines []*PerturbationIncidentLine `xml:"incidentLines,omitempty" json:"incidentLines,omitempty" yaml:"incidentLines,omitempty"`
	MessageGlobal *PerturbationMessage        `xml:"messageGlobal,omitempty" json:"messageGlobal,omitempty" yaml:"messageGlobal,omitempty"`
	Status        *string                     `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// PerturbationIncidentLine was auto-generated from WSDL.
type PerturbationIncidentLine struct {
	Consequence   *PerturbationConsequence `xml:"consequence,omitempty" json:"consequence,omitempty" yaml:"consequence,omitempty"`
	Line          *Line                    `xml:"line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
	MessageLarge  *PerturbationMessage     `xml:"messageLarge,omitempty" json:"messageLarge,omitempty" yaml:"messageLarge,omitempty"`
	MessageMedium *PerturbationMessage     `xml:"messageMedium,omitempty" json:"messageMedium,omitempty" yaml:"messageMedium,omitempty"`
	MessageShort  *PerturbationMessage     `xml:"messageShort,omitempty" json:"messageShort,omitempty" yaml:"messageShort,omitempty"`
	Stations      []*Station               `xml:"stations,omitempty" json:"stations,omitempty" yaml:"stations,omitempty"`
}

// PerturbationMessage was auto-generated from WSDL.
type PerturbationMessage struct {
	MediaSpecific *bool   `xml:"mediaSpecific,omitempty" json:"mediaSpecific,omitempty" yaml:"mediaSpecific,omitempty"`
	Text          *string `xml:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	Type          *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Updated       *bool   `xml:"updated,omitempty" json:"updated,omitempty" yaml:"updated,omitempty"`
}

// Reseau was auto-generated from WSDL.
type Reseau struct {
	Code  *string `xml:"http://wsiv.ratp.fr/xsd code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Id    *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Image *string `xml:"image,omitempty" json:"image,omitempty" yaml:"image,omitempty"`
	Name  *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Station was auto-generated from WSDL.
type Station struct {
	Direction   *Direction   `xml:"http://wsiv.ratp.fr/xsd direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	GeoPointA   *GeoPoint    `xml:"geoPointA,omitempty" json:"geoPointA,omitempty" yaml:"geoPointA,omitempty"`
	GeoPointR   *GeoPoint    `xml:"geoPointR,omitempty" json:"geoPointR,omitempty" yaml:"geoPointR,omitempty"`
	Id          *string      `xml:"http://wsiv.ratp.fr/xsd id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	IdsNextA    []*string    `xml:"idsNextA,omitempty" json:"idsNextA,omitempty" yaml:"idsNextA,omitempty"`
	IdsNextR    []*string    `xml:"idsNextR,omitempty" json:"idsNextR,omitempty" yaml:"idsNextR,omitempty"`
	Line        *Line        `xml:"http://wsiv.ratp.fr/xsd line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
	Name        *string      `xml:"http://wsiv.ratp.fr/xsd name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	StationArea *StationArea `xml:"stationArea,omitempty" json:"stationArea,omitempty" yaml:"stationArea,omitempty"`
}

// StationAcces was auto-generated from WSDL.
type StationAcces struct {
	Address        *string  `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Id             *string  `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Index          *string  `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Name           *string  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	TimeDaysLabel  *string  `xml:"timeDaysLabel,omitempty" json:"timeDaysLabel,omitempty" yaml:"timeDaysLabel,omitempty"`
	TimeDaysStatus *string  `xml:"timeDaysStatus,omitempty" json:"timeDaysStatus,omitempty" yaml:"timeDaysStatus,omitempty"`
	TimeEnd        *string  `xml:"timeEnd,omitempty" json:"timeEnd,omitempty" yaml:"timeEnd,omitempty"`
	TimeStart      *string  `xml:"timeStart,omitempty" json:"timeStart,omitempty" yaml:"timeStart,omitempty"`
	X              *float64 `xml:"x,omitempty" json:"x,omitempty" yaml:"x,omitempty"`
	Y              *float64 `xml:"y,omitempty" json:"y,omitempty" yaml:"y,omitempty"`
}

// StationArea was auto-generated from WSDL.
type StationArea struct {
	Access          []*StationAcces `xml:"access,omitempty" json:"access,omitempty" yaml:"access,omitempty"`
	Id              *string         `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Name            *string         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Stations        []*Station      `xml:"stations,omitempty" json:"stations,omitempty" yaml:"stations,omitempty"`
	TarifsToParis   []*Tarif        `xml:"tarifsToParis,omitempty" json:"tarifsToParis,omitempty" yaml:"tarifsToParis,omitempty"`
	ZoneCarteOrange *string         `xml:"zoneCarteOrange,omitempty" json:"zoneCarteOrange,omitempty" yaml:"zoneCarteOrange,omitempty"`
}

// Tarif was auto-generated from WSDL.
type Tarif struct {
	DemiTarif  *float64 `xml:"demiTarif,omitempty" json:"demiTarif,omitempty" yaml:"demiTarif,omitempty"`
	PleinTarif *float64 `xml:"pleinTarif,omitempty" json:"pleinTarif,omitempty" yaml:"pleinTarif,omitempty"`
	ViaLine    *Line    `xml:"viaLine,omitempty" json:"viaLine,omitempty" yaml:"viaLine,omitempty"`
	ViaReseau  *Reseau  `xml:"viaReseau,omitempty" json:"viaReseau,omitempty" yaml:"viaReseau,omitempty"`
}

// WrDirections was auto-generated from WSDL.
type WrDirections struct {
	AmbiguityMessage *string      `xml:"ambiguityMessage,omitempty" json:"ambiguityMessage,omitempty" yaml:"ambiguityMessage,omitempty"`
	AmbiguousLines   []*Line      `xml:"ambiguousLines,omitempty" json:"ambiguousLines,omitempty" yaml:"ambiguousLines,omitempty"`
	ArgumentLine     *Line        `xml:"argumentLine,omitempty" json:"argumentLine,omitempty" yaml:"argumentLine,omitempty"`
	Directions       []*Direction `xml:"directions,omitempty" json:"directions,omitempty" yaml:"directions,omitempty"`
}

// WrItineraries was auto-generated from WSDL.
type WrItineraries struct {
	AmbiguityMessage        *string      `xml:"ambiguityMessage,omitempty" json:"ambiguityMessage,omitempty" yaml:"ambiguityMessage,omitempty"`
	AmbiguousGeoPointsEnd   []*GeoPoint  `xml:"ambiguousGeoPointsEnd,omitempty" json:"ambiguousGeoPointsEnd,omitempty" yaml:"ambiguousGeoPointsEnd,omitempty"`
	AmbiguousGeoPointsStart []*GeoPoint  `xml:"ambiguousGeoPointsStart,omitempty" json:"ambiguousGeoPointsStart,omitempty" yaml:"ambiguousGeoPointsStart,omitempty"`
	ArgumentDate            *string      `xml:"argumentDate,omitempty" json:"argumentDate,omitempty" yaml:"argumentDate,omitempty"`
	Itineraries             []*Itinerary `xml:"itineraries,omitempty" json:"itineraries,omitempty" yaml:"itineraries,omitempty"`
}

// WrMission was auto-generated from WSDL.
type WrMission struct {
	AmbiguityMessage *string  `xml:"ambiguityMessage,omitempty" json:"ambiguityMessage,omitempty" yaml:"ambiguityMessage,omitempty"`
	AmbiguousLines   []*Line  `xml:"ambiguousLines,omitempty" json:"ambiguousLines,omitempty" yaml:"ambiguousLines,omitempty"`
	ArgumentDate     *string  `xml:"argumentDate,omitempty" json:"argumentDate,omitempty" yaml:"argumentDate,omitempty"`
	ArgumentLine     *Line    `xml:"argumentLine,omitempty" json:"argumentLine,omitempty" yaml:"argumentLine,omitempty"`
	Mission          *Mission `xml:"mission,omitempty" json:"mission,omitempty" yaml:"mission,omitempty"`
}

// WrMissions was auto-generated from WSDL.
type WrMissions struct {
	AmbiguityMessage    *string         `xml:"ambiguityMessage,omitempty" json:"ambiguityMessage,omitempty" yaml:"ambiguityMessage,omitempty"`
	AmbiguousDirections []*Direction    `xml:"ambiguousDirections,omitempty" json:"ambiguousDirections,omitempty" yaml:"ambiguousDirections,omitempty"`
	AmbiguousLines      []*Line         `xml:"ambiguousLines,omitempty" json:"ambiguousLines,omitempty" yaml:"ambiguousLines,omitempty"`
	AmbiguousStations   []*Station      `xml:"ambiguousStations,omitempty" json:"ambiguousStations,omitempty" yaml:"ambiguousStations,omitempty"`
	ArgumentDate        *string         `xml:"argumentDate,omitempty" json:"argumentDate,omitempty" yaml:"argumentDate,omitempty"`
	ArgumentDirection   *Direction      `xml:"argumentDirection,omitempty" json:"argumentDirection,omitempty" yaml:"argumentDirection,omitempty"`
	ArgumentLine        *Line           `xml:"argumentLine,omitempty" json:"argumentLine,omitempty" yaml:"argumentLine,omitempty"`
	ArgumentStation     *Station        `xml:"argumentStation,omitempty" json:"argumentStation,omitempty" yaml:"argumentStation,omitempty"`
	Missions            []*Mission      `xml:"missions,omitempty" json:"missions,omitempty" yaml:"missions,omitempty"`
	Perturbations       []*Perturbation `xml:"perturbations,omitempty" json:"perturbations,omitempty" yaml:"perturbations,omitempty"`
}

// WrPerturbations was auto-generated from WSDL.
type WrPerturbations struct {
	ArgumentMedia  *string         `xml:"argumentMedia,omitempty" json:"argumentMedia,omitempty" yaml:"argumentMedia,omitempty"`
	ArgumentSource *string         `xml:"argumentSource,omitempty" json:"argumentSource,omitempty" yaml:"argumentSource,omitempty"`
	Perturbations  []*Perturbation `xml:"perturbations,omitempty" json:"perturbations,omitempty" yaml:"perturbations,omitempty"`
}

// WrStations was auto-generated from WSDL.
type WrStations struct {
	AmbiguityMessage   *string     `xml:"ambiguityMessage,omitempty" json:"ambiguityMessage,omitempty" yaml:"ambiguityMessage,omitempty"`
	AmbiguousGeoPoints []*GeoPoint `xml:"ambiguousGeoPoints,omitempty" json:"ambiguousGeoPoints,omitempty" yaml:"ambiguousGeoPoints,omitempty"`
	AmbiguousLines     []*Line     `xml:"ambiguousLines,omitempty" json:"ambiguousLines,omitempty" yaml:"ambiguousLines,omitempty"`
	ArgumentDirection  *Direction  `xml:"argumentDirection,omitempty" json:"argumentDirection,omitempty" yaml:"argumentDirection,omitempty"`
	ArgumentGeoPoint   *GeoPoint   `xml:"argumentGeoPoint,omitempty" json:"argumentGeoPoint,omitempty" yaml:"argumentGeoPoint,omitempty"`
	ArgumentLine       *Line       `xml:"argumentLine,omitempty" json:"argumentLine,omitempty" yaml:"argumentLine,omitempty"`
	Distances          []*int      `xml:"distances,omitempty" json:"distances,omitempty" yaml:"distances,omitempty"`
	Stations           []*Station  `xml:"stations,omitempty" json:"stations,omitempty" yaml:"stations,omitempty"`
}

// GetDirections was auto-generated from WSDL.
type GetDirections struct {
	Line *Line `xml:"http://wsiv.ratp.fr line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
}

// GetDirectionsResponse was auto-generated from WSDL.
type GetDirectionsResponse struct {
	Return *WrDirections `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetGeoPoints was auto-generated from WSDL.
type GetGeoPoints struct {
	Gp    *GeoPoint `xml:"gp,omitempty" json:"gp,omitempty" yaml:"gp,omitempty"`
	Limit *int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetGeoPointsResponse was auto-generated from WSDL.
type GetGeoPointsResponse struct {
	Return []*GeoPoint `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetLines was auto-generated from WSDL.
type GetLines struct {
	Line *Line `xml:"tns:line,omitempty" json:"line,omitempty" yaml:"line,omitempty"`
}

// GetLinesResponse was auto-generated from WSDL.
type GetLinesResponse struct {
	Return []*Line `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMission was auto-generated from WSDL.
type GetMission struct {
	Mission          *Mission `xml:"mission,omitempty" json:"mission,omitempty" yaml:"mission,omitempty"`
	Date             *string  `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	StationAll       *bool    `xml:"stationAll,omitempty" json:"stationAll,omitempty" yaml:"stationAll,omitempty"`
	StationSortAlpha *bool    `xml:"stationSortAlpha,omitempty" json:"stationSortAlpha,omitempty" yaml:"stationSortAlpha,omitempty"`
}

// GetMissionResponse was auto-generated from WSDL.
type GetMissionResponse struct {
	Return *WrMission `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMissionsFirstLast was auto-generated from WSDL.
type GetMissionsFirstLast struct {
	Station   *Station   `xml:"http://wsiv.ratp.fr station,omitempty" json:"station,omitempty" yaml:"station,omitempty"`
	Direction *Direction `xml:"http://wsiv.ratp.fr direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	Date      *string    `xml:"http://wsiv.ratp.fr date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
}

// GetMissionsFirstLastResponse was auto-generated from WSDL.
type GetMissionsFirstLastResponse struct {
	Return *WrMissions `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMissionsFrequency was auto-generated from WSDL.
type GetMissionsFrequency struct {
	Station        *Station   `xml:"http://wsiv.ratp.fr station,omitempty" json:"station,omitempty" yaml:"station,omitempty"`
	Direction      *Direction `xml:"http://wsiv.ratp.fr direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	StationEndLine *Station   `xml:"http://wsiv.ratp.fr stationEndLine,omitempty" json:"stationEndLine,omitempty" yaml:"stationEndLine,omitempty"`
	StationEnd     *Station   `xml:"http://wsiv.ratp.fr stationEnd,omitempty" json:"stationEnd,omitempty" yaml:"stationEnd,omitempty"`
	DatesStart     []*string  `xml:"http://wsiv.ratp.fr datesStart,omitempty" json:"datesStart,omitempty" yaml:"datesStart,omitempty"`
	DatesEnd       []*string  `xml:"http://wsiv.ratp.fr datesEnd,omitempty" json:"datesEnd,omitempty" yaml:"datesEnd,omitempty"`
}

// GetMissionsFrequencyResponse was auto-generated from WSDL.
type GetMissionsFrequencyResponse struct {
	Return *WrMissions `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetMissionsNext was auto-generated from WSDL.
type GetMissionsNext struct {
	Station   *Station   `xml:"http://wsiv.ratp.fr station,omitempty" json:"station,omitempty" yaml:"station,omitempty"`
	Direction *Direction `xml:"http://wsiv.ratp.fr direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	DateStart *string    `xml:"http://wsiv.ratp.fr dateStart,omitempty" json:"dateStart,omitempty" yaml:"dateStart,omitempty"`
	Limit     *int       `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetMissionsNextResponse was auto-generated from WSDL.
type GetMissionsNextResponse struct {
	Return *WrMissions `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetPerturbations was auto-generated from WSDL.
type GetPerturbations struct {
	Perturbation *Perturbation `xml:"http://wsiv.ratp.fr perturbation,omitempty" json:"perturbation,omitempty" yaml:"perturbation,omitempty"`
	IsXmlText    *bool         `xml:"http://wsiv.ratp.fr isXmlText,omitempty" json:"isXmlText,omitempty" yaml:"isXmlText,omitempty"`
}

// GetPerturbationsResponse was auto-generated from WSDL.
type GetPerturbationsResponse struct {
	Return *WrPerturbations `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetStations was auto-generated from WSDL.
type GetStations struct {
	Station   *Station  `xml:"http://wsiv.ratp.fr station,omitempty" json:"station,omitempty" yaml:"station,omitempty"`
	Gp        *GeoPoint `xml:"http://wsiv.ratp.fr gp,omitempty" json:"gp,omitempty" yaml:"gp,omitempty"`
	Distances []*int    `xml:"distances,omitempty" json:"distances,omitempty" yaml:"distances,omitempty"`
	Limit     *int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
	SortAlpha *bool     `xml:"sortAlpha,omitempty" json:"sortAlpha,omitempty" yaml:"sortAlpha,omitempty"`
}

// GetStationsResponse was auto-generated from WSDL.
type GetStationsResponse struct {
	Return *WrStations `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetVersionResponse was auto-generated from WSDL.
type GetVersionResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for GetDirections.
// OperationGetDirectionsRequest was auto-generated from WSDL.
type OperationGetDirectionsRequest struct {
	Parameters *GetDirections `xml:"ns:getDirections,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetDirections.
// OperationGetDirectionsResponse was auto-generated from WSDL.
type OperationGetDirectionsResponse struct {
	Parameters *GetDirectionsResponse `xml:"getDirectionsResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetGeoPoints.
// OperationGetGeoPointsRequest was auto-generated from WSDL.
type OperationGetGeoPointsRequest struct {
	Parameters *GetGeoPoints `xml:"ns:getGeoPoints,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetGeoPoints.
// OperationGetGeoPointsResponse was auto-generated from WSDL.
type OperationGetGeoPointsResponse struct {
	Parameters *GetGeoPointsResponse `xml:"getGeoPointsResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetLines.
// OperationGetLinesRequest was auto-generated from WSDL.
type OperationGetLinesRequest struct {
	Parameters *GetLines `xml:"ns:getLines,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetLines.
// OperationGetLinesResponse was auto-generated from WSDL.
type OperationGetLinesResponse struct {
	Parameters *GetLinesResponse `xml:"getLinesResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMission.
// OperationGetMissionRequest was auto-generated from WSDL.
type OperationGetMissionRequest struct {
	Parameters *GetMission `xml:"ns:getMission,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMission.
// OperationGetMissionResponse was auto-generated from WSDL.
type OperationGetMissionResponse struct {
	Parameters *GetMissionResponse `xml:"getMissionResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsFirstLast.
// OperationGetMissionsFirstLastRequest was auto-generated from
// WSDL.
type OperationGetMissionsFirstLastRequest struct {
	Parameters *GetMissionsFirstLast `xml:"ns:getMissionsFirstLast,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsFirstLast.
// OperationGetMissionsFirstLastResponse was auto-generated from
// WSDL.
type OperationGetMissionsFirstLastResponse struct {
	Parameters *GetMissionsFirstLastResponse `xml:"getMissionsFirstLastResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsFrequency.
// OperationGetMissionsFrequencyRequest was auto-generated from
// WSDL.
type OperationGetMissionsFrequencyRequest struct {
	Parameters *GetMissionsFrequency `xml:"ns:getMissionsFrequency,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsFrequency.
// OperationGetMissionsFrequencyResponse was auto-generated from
// WSDL.
type OperationGetMissionsFrequencyResponse struct {
	Parameters *GetMissionsFrequencyResponse `xml:"getMissionsFrequencyResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsNext.
// OperationGetMissionsNextRequest was auto-generated from WSDL.
type OperationGetMissionsNextRequest struct {
	Parameters *GetMissionsNext `xml:"ns:getMissionsNext,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetMissionsNext.
// OperationGetMissionsNextResponse was auto-generated from WSDL.
type OperationGetMissionsNextResponse struct {
	Parameters *GetMissionsNextResponse `xml:"getMissionsNextResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetPerturbations.
// OperationGetPerturbationsRequest was auto-generated from WSDL.
type OperationGetPerturbationsRequest struct {
	Parameters *GetPerturbations `xml:"ns:getPerturbations,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetPerturbations.
// OperationGetPerturbationsResponse was auto-generated from WSDL.
type OperationGetPerturbationsResponse struct {
	Parameters *GetPerturbationsResponse `xml:"getPerturbationsResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetStations.
// OperationGetStationsRequest was auto-generated from WSDL.
type OperationGetStationsRequest struct {
	Parameters *GetStations `xml:"ns:getStations,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetStations.
// OperationGetStationsResponse was auto-generated from WSDL.
type OperationGetStationsResponse struct {
	Parameters *GetStationsResponse `xml:"getStationsResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type OperationGetVersionRequest struct {
	Parameters *GetVersionResponse `xml:"ns:getVersion,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for GetVersion.
// OperationGetVersionResponse was auto-generated from WSDL.
type OperationGetVersionResponse struct {
	Parameters *GetVersionResponse `xml:"getVersionResponse,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// wsivPortType implements the WsivPortType interface.
type wsivPortType struct {
	cli *soap.Client
}

// GetDirections was auto-generated from WSDL.
func (p *wsivPortType) GetDirections(parameters *GetDirections) (*GetDirectionsResponse, error) {
	α := OperationGetDirectionsRequest{
		parameters,
	}

	γ := OperationGetDirectionsResponse{}
	if err := p.cli.RoundTripWithAction("getDirections", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetGeoPoints was auto-generated from WSDL.
func (p *wsivPortType) GetGeoPoints(parameters *GetGeoPoints) (*GetGeoPointsResponse, error) {
	α := OperationGetGeoPointsRequest{
		parameters,
	}

	γ := OperationGetGeoPointsResponse{}
	if err := p.cli.RoundTripWithAction("getGeoPoints", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetLines was auto-generated from WSDL.
func (p *wsivPortType) GetLines(parameters *GetLines) (*GetLinesResponse, error) {
	α := OperationGetLinesRequest{
		parameters,
	}

	γ := OperationGetLinesResponse{}
	if err := p.cli.RoundTripWithAction("getLines", α, &γ); err != nil {
		return nil, err
	}
	fmt.Println(γ)
	return γ.Parameters, nil
}

// GetMission was auto-generated from WSDL.
func (p *wsivPortType) GetMission(parameters *GetMission) (*GetMissionResponse, error) {
	α := OperationGetMissionRequest{
		parameters,
	}

	γ := OperationGetMissionResponse{}
	if err := p.cli.RoundTripWithAction("getMission", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetMissionsFirstLast was auto-generated from WSDL.
func (p *wsivPortType) GetMissionsFirstLast(parameters *GetMissionsFirstLast) (*GetMissionsFirstLastResponse, error) {
	α := OperationGetMissionsFirstLastRequest{
		parameters,
	}

	γ := OperationGetMissionsFirstLastResponse{}
	if err := p.cli.RoundTripWithAction("getMissionsFirstLast", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetMissionsFrequency was auto-generated from WSDL.
func (p *wsivPortType) GetMissionsFrequency(parameters *GetMissionsFrequency) (*GetMissionsFrequencyResponse, error) {
	α := OperationGetMissionsFrequencyRequest{
		parameters,
	}

	γ := OperationGetMissionsFrequencyResponse{}
	if err := p.cli.RoundTripWithAction("getMissionsFrequency", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetMissionsNext was auto-generated from WSDL.
func (p *wsivPortType) GetMissionsNext(parameters *GetMissionsNext) (*GetMissionsNextResponse, error) {
	α := OperationGetMissionsNextRequest{
		parameters,
	}

	γ := OperationGetMissionsNextResponse{}
	if err := p.cli.RoundTripWithAction("getMissionsNext", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetPerturbations was auto-generated from WSDL.
func (p *wsivPortType) GetPerturbations(parameters *GetPerturbations) (*GetPerturbationsResponse, error) {
	α := OperationGetPerturbationsRequest{
		parameters,
	}

	γ := OperationGetPerturbationsResponse{}
	if err := p.cli.RoundTripWithAction("getPerturbations", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetStations was auto-generated from WSDL.
func (p *wsivPortType) GetStations(parameters *GetStations) (*GetStationsResponse, error) {
	α := OperationGetStationsRequest{
		parameters,
	}

	γ := OperationGetStationsResponse{}
	if err := p.cli.RoundTripWithAction("getStations", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// GetVersion was auto-generated from WSDL.
func (p *wsivPortType) GetVersion() (*GetVersionResponse, error) {
	α := OperationGetVersionRequest{&GetVersionResponse{}}

	γ := OperationGetVersionResponse{}
	if err := p.cli.RoundTripWithAction("getVersion", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}
