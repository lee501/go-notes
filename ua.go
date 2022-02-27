package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func GetFp(line string) string {
	fp := make([]byte, 0)
	for i, v := range line {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			fp = append(fp, line[i])
		}
	}
	return string(fp)
}

type BookingParam struct {
	MemberID               string `json:"memberId"`
	MemberPhone            string `json:"memberPhone"`
	IsDomestic             int    `json:"isDomestic"`
	Channel                string `json:"channel"`
	SpikeActivitySegmentID int    `json:"spikeActivitySegmentId"`
	SpecialIdentityMark    int    `json:"specialIdentityMark"`
	OrderType              string `json:"orderType"`
	ContactInfo            struct {
		ContactName   string `json:"contactName"`
		ContactMobile string `json:"contactMobile"`
		ContactEmail  string `json:"contactEmail"`
	} `json:"contactInfo"`
	Insurance    string        `json:"insurance"`
	Insurances   []interface{} `json:"insurances"`
	Mail         interface{}   `json:"mail"`
	PromoType    string        `json:"promoType"`
	PromoCode    string        `json:"promoCode"`
	PromoName    string        `json:"promoName"`
	PromoPrice   int           `json:"promoPrice"`
	DiscountType string        `json:"discountType"`
	Passengers   []struct {
		PassengerIndex     string `json:"passengerIndex"`
		PassengerName      string `json:"passengerName"`
		SurName            string `json:"surName"`
		FirstName          string `json:"firstName"`
		PassengerType      string `json:"passengerType"`
		PassengerSex       string `json:"passengerSex"`
		PassengerIDType    string `json:"passengerIdType"`
		PassengerIdno      string `json:"passengerIdno"`
		Birthday           string `json:"birthday"`
		ParentIndex        string `json:"parentIndex"`
		SpecialIDType      string `json:"specialIdType"`
		SpecialIDNo        string `json:"specialIdNo"`
		IssuedNationalName string `json:"issuedNationalName"`
		IDNationalName     string `json:"idNationalName"`
		IDValidityDate     string `json:"idValidityDate"`
	} `json:"passengers"`
	Segments []struct {
		FlightNo        string      `json:"flightNo"`
		FlightDate      string      `json:"flightDate"`
		SegmentIndex    string      `json:"segmentIndex"`
		DepAirportCode  string      `json:"depAirportCode"`
		DepCity         string      `json:"depCity"`
		DepDateTime     string      `json:"depDateTime"`
		ArrAirportCode  string      `json:"arrAirportCode"`
		ArrCity         string      `json:"arrCity"`
		ArrDateTime     string      `json:"arrDateTime"`
		MasterCabinCode string      `json:"masterCabinCode"`
		FloatProductID  interface{} `json:"floatProductId"`
		BasePrice       int         `json:"basePrice"`
		FlyingTime      string      `json:"flyingTime"`
		Cabins          []struct {
			Ei             string `json:"ei"`
			FareBase       string `json:"fareBase"`
			PassengerType  string `json:"passengerType"`
			AirportTax     int    `json:"airportTax"`
			OilTax         int    `json:"oilTax"`
			IntlTotalTax   int    `json:"intlTotalTax"`
			CabinCode      string `json:"cabinCode"`
			SubCabinCode   string `json:"subCabinCode"`
			OriginDiscount int    `json:"originDiscount"`
			OriginPrice    int    `json:"originPrice"`
			SalePrice      int    `json:"salePrice"`
		} `json:"cabins"`
	} `json:"segments"`
	TicketDetails []struct {
		SegmentIndex       string `json:"segmentIndex"`
		PassengerIndex     string `json:"passengerIndex"`
		OriginPrice        int    `json:"originPrice"`
		SalePrice          int    `json:"salePrice"`
		AirportTax         int    `json:"airportTax"`
		OilTax             int    `json:"oilTax"`
		InsuranceAmount    int    `json:"insuranceAmount"`
		InsuranceDeduction int    `json:"insuranceDeduction"`
		Mileage            int    `json:"mileage"`
		MileageDeduction   int    `json:"mileageDeduction"`
		CouponDeduction    int    `json:"couponDeduction"`
		FinalPrice         int    `json:"finalPrice"`
		IntlTaxJSON        string `json:"intlTaxJson"`
		IntlTotalTax       int    `json:"intlTotalTax"`
	} `json:"ticketDetails"`
	TotalInsureAmount int    `json:"totalInsureAmount"`
	TotalTicketAmount int    `json:"totalTicketAmount"`
	TotalAmount       int    `json:"totalAmount"`
	Timestamp         int    `json:"timestamp"`
	Token             string `json:"token"`
	Sign              string `json:"sign"`
}

func main() {
	l := "Mozilla/4.0 (compatible; tyty 7.0; Windows NT 5.1)"
	fmt.Println(GetFp(l))

	tid := "b7044573df9d7ecab5bfd57d6f5197c09752a955:y2:0bd83d43-9484-11ec-a209-0a92b7ba9948:ae97f520bd"
	re, _ := url.QueryUnescape(tid)
	fmt.Println(re)

	js := "{\"channel\":\"\",\"discountType\":\"\",\"insurance\": \"false\",\"isDomestic\":1,\"memberId\":\"Dl6oMYPXM/76Nw2pl007pg==\",\"memberPhone\":\"zd1JrRWpVZ9aUxBF6RHQTw==\",\"orderType\":\"OW\",\"promoCode\":\"\",\"promoName\":\"\",\"promoPrice\":0,\"promoType\":\"\",\"specialIdentityMark\":0,\"spikeActivitySegmentId\":0,\"timestamp\":1645615196,\"token\":\"54521d69115743b39f4257ccf6848a91\",\"totalAmount\":1692,\"totalInsureAmount\":0,\"totalTicketAmount\":1692,\"sign\":\"vWUgmaeQQWta/zzqs45WdMZ7kihb4IjqRntnSue1QRE2P/cYq1TacIfsMWuXUjPPcVbAb8vxCG04Gs5b4aAug/ySHEmfEykxYDFKSeQy1hv2JhQVaPSbDsswTjn2A93Vj9MxawhM3O9lxoKTB8LskvUmxslVoX+oynAYkcZutIku/OI6nlS9+1MRkhnKxkU3Q4fAu2cisn1IpCcOaBzV+Vlb1Ieehu5m9wHhDkI1JXooBm6TlsxcgnSJAsYG6Y45khsD8c17rO2W558iw98ZCeDPwbXYk8iN7ltSQ1uaZYWubNz/xOONXXgfn8hmxtTBwA5SKK4T5IcFSWwL3z8qado3Km2AzMulw2dRcVw+bU1SbndsTnDcHxmpoPtwjLc1LeaTNjMBSSqqYprCXXxePtpzXRKsR/Q2L2pwvQq7jTt94tjEnhNdQs3/dBHoOCViTHeIwvpfG+2ReeUlHfSiQw==\",\"insurances\":[],\"mail\":null,\"contactInfo\":{\"contactName\":\"王伟\",\"contactMobile\":\"18738558927\",\"contactEmail\":\"\"},\"passengers\":[{\"passengerIndex\":\"6e7595b6-7b77-454b-a0fb-e2d11fc76a01\",\"passengerName\":\"王伟\",\"firstName\":\"\",\"surName\":\"\",\"passengerType\":\"0\",\"passengerSex\":\"2\",\"passengerIdType\":\"01\",\"passengerIdno\":\"532626199807243314\",\"birthday\":\"1998-07-24\",\"parentIndex\":\"\",\"specialIdType\":\"\",\"specialIdNo\":\"\"},{\"passengerIndex\":\"56d6e7a9-31ab-4960-a7f5-fa09d9fc7758\",\"passengerName\":\"王风\",\"firstName\":\"\",\"surName\":\"\",\"passengerType\":\"0\",\"passengerSex\":\"2\",\"passengerIdType\":\"01\",\"passengerIdno\":\"532626198907240319\",\"birthday\":\"1989-07-24\",\"parentIndex\":\"\",\"specialIdType\":\"\",\"specialIdNo\":\"\"}],\"ticketDetails\":[{\"segmentIndex\":\"1\",\"passengerIndex\":\"6e7595b6-7b77-454b-a0fb-e2d11fc76a01\",\"originPrice\":776,\"salePrice\":776,\"airportTax\":50,\"oilTax\":20,\"insuranceAmount\":0,\"insuranceDeduction\":0,\"mileage\":0,\"mileageDeduction\":0,\"couponDeduction\":0,\"finalPrice\":846},{\"segmentIndex\":\"1\",\"passengerIndex\":\"56d6e7a9-31ab-4960-a7f5-fa09d9fc7758\",\"originPrice\":776,\"salePrice\":776,\"airportTax\":50,\"oilTax\":20,\"insuranceAmount\":0,\"insuranceDeduction\":0,\"mileage\":0,\"mileageDeduction\":0,\"couponDeduction\":0,\"finalPrice\":846}],\"segments\":[{\"flightNo\":\"KY8219\",\"flightDate\":\"2022-02-28\",\"segmentIndex\":\"1\",\"depAirportCode\":\"KMG\",\"depCity\":\"昆明\",\"depDateTime\":\"2022-02-28 08:30\",\"arrAirportCode\":\"HYN\",\"arrCity\":\"黄岩\",\"arrDateTime\":\"2022-02-28 14:15\",\"masterCabinCode\":\"P\",\"floatProductId\":null,\"basePrice\":2380,\"cabins\":[{\"ei\":\"不得签转\",\"fareBase\":\"PDT776\",\"passengerType\":\"0\",\"airportTax\":50,\"oilTax\":20,\"cabinCode\":\"P\",\"subCabinCode\":\"P\",\"originDiscount\":33,\"originPrice\":776,\"salePrice\":776}]}]}"
	js1 := "{\"memberId\":\"\",\"memberPhone\":\"\",\"tripType\":\"OW\",\"passType\":0,\"depAirportCode\":\"JHG\",\"arrAirportCode\":\"WMT\",\"depDate\":\"2022-02-25\",\"returnDate\":\"\",\"timestamp\":1645665579,\"token\":\"ddf853dbcb454032a9445e5b6852f0d2\",\"sign\":\"kRHHXo1BXW/ftgS2N3FxxrnBnDUXv17nLCFeBNc2R24PrHqLqyFsDAmcMT5s99N5R/oEfIz0mfsbczAGduFmU0X4qO18mr3M75G5i4kLIETVDFSc6E24Ju5VCmv5feeiNbDgAOVPzO+i1vCs0ziQiGF3/gMx1gNsbhVB36sgY8XywWnfbpfDeGB89gLIehfP1rZZ0B4z8BEeeM+k9U5ayMHJccQ+U7k+IyB/ybbcAsI=\"}"
	var res BookingParam
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		fmt.Println("11111", err)
	}
	fmt.Println(res)
	err1 := json.Unmarshal([]byte(js1), &res)
	if err1 != nil {
		fmt.Println("正常", err1)
	}
	fmt.Println(res)
}
