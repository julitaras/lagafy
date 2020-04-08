package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Travel entity
type Travel struct {
	ID               uint           `json:"id,omitempty gorm:"primary_key"`
	HasWifi          bool           `json:"hasWifi"`
	Capacity         int            `json:"capacity" validate:"required,numeric,min=0"`
	Driver           string         `json:"driver"`
	Departure        time.Time      `json:"departure" validate:"required"`
	DepartureAddress string         `json:"departureAddress" validate:"required"`
	Arrival          time.Time      `json:"arrival"`
	ArrivalAddress   string         `json:"arrivalAddress" validate:"required"`
	Origin           string         `json:"origin" validate:"required"`
	Destination      string         `json:"destination" validate:"required"`
	Status           string         `json:"status"`
	Template         bool           `json:"template"`
	Reservations     []*Reservation `json:"reservations" gorm:"foreignkey:TravelID"`
	DeletedAt        *time.Time     `json:"deleted_at"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

//Passenger entity
type Passenger struct {
	ID        uint       `json:"id,omitepty" gorm:"primary_key"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required" gorm:"unique;not null"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

//Reservation entity
type Reservation struct {
	ID          uint `json:"id,omitepty"`
	TravelID    uint
	Travel      *Travel
	Passenger   Passenger `json:"passenger" validate:"required" gorm:"foreignkey:passenger_id"`
	PassengerID uint
	Status      string     `json:"status"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

//ClaimsData struct
type ClaimsData struct {
	Aud            string   `json:"aud"`
	Iss            string   `json:"iss"`
	Iat            int      `json:"iat"`
	Nbf            int      `json:"nbf"`
	Exp            int      `json:"exp"`
	Acct           int      `json:"acct"`
	Acr            string   `json:"acr"`
	Aio            string   `json:"aio"`
	Amr            []string `json:"amr"`
	AppDisplayname string   `json:"app_displayname"`
	Appid          string   `json:"appid"`
	Appidacr       string   `json:"appidacr"`
	FamilyName     string   `json:"family_name"`
	GivenName      string   `json:"given_name"`
	Ipaddr         string   `json:"ipaddr"`
	Name           string   `json:"name"`
	Oid            string   `json:"oid"`
	OnpremSid      string   `json:"onprem_sid"`
	Platf          string   `json:"platf"`
	Puid           string   `json:"puid"`
	Scp            string   `json:"scp"`
	Sub            string   `json:"sub"`
	Tid            string   `json:"tid"`
	UniqueName     string   `json:"unique_name"`
	Upn            string   `json:"upn"`
	Uti            string   `json:"uti"`
	Ver            string   `json:"ver"`
	XmsSt          struct {
		Sub string `json:"sub"`
	} `json:"xms_st"`
	XmsTcdt int `json:"xms_tcdt"`
	jwt.StandardClaims
}

//Travel Information Entity
type TravelInformation struct {
	ID          uint       `json:"id,omitepty" `
	Capacity    int        `json:"capacity" validate:"required,numeric,min=0"`
	Departure   time.Time  `json:"departure" validate:"required"`
	Arrival     time.Time  `json:"arrival"`
	Origin      string     `json:"origin" validate:"required"`
	Destination string     `json:"destination" validate:"required"`
	Pending     int        `json:"pending"`
	OnBoard     int        `json:"onboard"`
	Cancelled   int        `json:"cancelled"`
	Confirmed   int        `json:"confirmed"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
