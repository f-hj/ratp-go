package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Hello struct {
	RATPVersion string `json:"ratpVersion"`
	APIVersion  string `json:"apiVersion"`
}

type PrettyLines struct {
	Metros   []Line `json:"metros"`
	RER      []Line `json:"rers"`
	Tramways []Line `json:"tramways"`
	Buses    []Line `json:"buses"`
}

type APIError struct {
	Err  string      `json:"err"`
	Info string      `json:"info"`
	Body interface{} `json:"bodyError"`
}

func main() {
	cli := soap.Client{
		URL:           "http://opendata-tr.ratp.fr/wsiv/services/Wsiv",
		Namespace:     "http://wsiv.ratp.fr/xsd",
		ThisNamespace: "http://wsiv.ratp.fr",
		Header:        "",
		Pre: func(req *http.Request) {
			fmt.Println(req)
		},
		Post: func(res *http.Response) {
			fmt.Println(res)
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
			res.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		},
	}

	soapService := NewWsivPortType(&cli)

	// Echo instance
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*soap.HTTPError); ok {
			c.JSON(http.StatusServiceUnavailable, APIError{
				Err:  "RATP error",
				Info: "SOAP from RATP sent an unrecoverable error",
				Body: he,
			})
			return
		}
		if he, ok := err.(*echo.HTTPError); ok {
			c.JSON(he.Code, APIError{
				Err:  he.Message.(string),
				Info: "Normal API error (this is your fault, check the doc!)",
				Body: he,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, APIError{
			Err:  "Internal error",
			Info: "This is logged, sorry",
			Body: err,
		})
		c.Logger().Error(err)
		return
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		version, err := soapService.GetVersion()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &Hello{
			RATPVersion: *version.Return,
			APIVersion:  "git_commit",
		})
	})

	e.GET("/lines", func(c echo.Context) error {
		lines, err := soapService.GetLines(&GetLines{&Line{}})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *lines)
	})

	e.GET("/prettyLines", func(c echo.Context) error {
		lines, err := soapService.GetLines(&GetLines{&Line{}})

		if err != nil {
			return err
		}
		prettyLines := PrettyLines{}
		for _, line := range lines.Return {
			line.ComputedCode = line.GetComputedCode()

			if *line.Reseau.Code == "metro" {
				prettyLines.Metros = append(prettyLines.Metros, *line)
			}
			if *line.Reseau.Code == "rer" {
				prettyLines.RER = append(prettyLines.RER, *line)
			}
			if *line.Reseau.Code == "tram" {
				prettyLines.Tramways = append(prettyLines.Tramways, *line)
			}
			if *line.Reseau.Code == "bus" {
				prettyLines.Buses = append(prettyLines.Buses, *line)
			}
		}
		return c.JSON(http.StatusOK, prettyLines)
	})

	e.GET("/lines/:line/perturbations", func(c echo.Context) error {
		line := c.Param("line")
		perturbations, err := soapService.GetPerturbations(&GetPerturbations{
			Perturbation: &Perturbation{
				Line: &Line{
					Id: &line,
				},
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *perturbations.Return)
	})

	e.GET("/lines/:line/stations", func(c echo.Context) error {
		line := c.Param("line")
		stations, err := soapService.GetStations(&GetStations{
			Station: &Station{
				Line: &Line{
					Id: &line,
				},
				/*Direction: &Direction{
					Sens: &sensA,
				},*/
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *stations.Return)
	})

	e.GET("/lines/:line/stations/way/:direction", func(c echo.Context) error {
		line := c.Param("line")
		direction := c.Param("direction")
		stations, err := soapService.GetStations(&GetStations{
			Station: &Station{
				Line: &Line{
					Id: &line,
				},
				Direction: &Direction{
					Sens: &direction,
				},
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *stations.Return)
	})

	e.GET("/lines/:line/directions", func(c echo.Context) error {
		line := c.Param("line")
		directions, err := soapService.GetDirections(&GetDirections{
			Line: &Line{
				Id: &line,
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *directions.Return)
	})

	e.GET("/lines/:line/stations/:station/next", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		missions, err := soapService.GetMissionsNext(&GetMissionsNext{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *missions.Return)
	})

	e.GET("/lines/:line/stations/:station/way/:direction/next", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		direction := c.Param("direction")
		missions, err := soapService.GetMissionsNext(&GetMissionsNext{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
			Direction: &Direction{
				Sens: &direction,
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *missions.Return)
	})

	e.GET("/lines/:line/stations/:station/direction/:direction/next", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		direction := c.Param("direction")
		missions, err := soapService.GetMissionsNext(&GetMissionsNext{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
			Direction: &Direction{
				Line: &Line{
					Id: &line,
				},
				StationsEndLine: [](*Station){
					&Station{
						Id: &direction,
						Line: &Line{
							Id: &line,
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *missions.Return)
	})

	e.GET("/lines/:line/stations/:station/way/:direction/firstAndLast", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		direction := c.Param("direction")
		t := time.Now()
		date := t.Format("200601021504")
		firstAndLast, err := soapService.GetMissionsFirstLast(&GetMissionsFirstLast{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
			Direction: &Direction{
				Sens: &direction,
			},
			Date: &date,
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *firstAndLast.Return)
	})

	e.GET("/lines/:line/stations/:station/way/:direction/frequency", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		direction := c.Param("direction")
		frequency, err := soapService.GetMissionsFrequency(&GetMissionsFrequency{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
			Direction: &Direction{
				Sens: &direction,
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *frequency.Return)
	})

	e.GET("/lines/:line/stations/:station/way/:way/direction/:direction/frequency", func(c echo.Context) error {
		line := c.Param("line")
		station := c.Param("station")
		way := c.Param("way")
		direction := c.Param("direction")
		frequency, err := soapService.GetMissionsFrequency(&GetMissionsFrequency{
			Station: &Station{
				Id: &station,
				Line: &Line{
					Id: &line,
				},
			},
			Direction: &Direction{
				Sens: &way,
			},
			StationEndLine: &Station{
				Id: &direction,
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *frequency.Return)
	})

	e.GET("/stations/search/:req", func(c echo.Context) error {
		req := c.Param("req")
		stations, err := soapService.GetStations(&GetStations{
			Station: &Station{
				Name: &req,
			},
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *stations.Return)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
