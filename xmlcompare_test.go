package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

type Androidx struct {
	XMLName      xml.Name `xml:"androidx.constraintlayout.widget.ConstraintLayout"`
	Text         string   `xml:",chardata"`
	Android      string   `xml:"android,attr"`
	App          string   `xml:"app,attr"`
	Tools        string   `xml:"tools,attr"`
	LayoutWidth  string   `xml:"layout_width,attr"`
	LayoutHeight string   `xml:"layout_height,attr"`
	Context      string   `xml:"context,attr"`
	ScrollView   struct {
		Text                                           string `xml:",chardata"`
		ID                                             string `xml:"id,attr"`
		LayoutWidth                                    string `xml:"layout_width,attr"`
		LayoutHeight                                   string `xml:"layout_height,attr"`
		LayoutConstraintEndToEndOf                     string `xml:"layout_constraintEnd_toEndOf,attr"`
		LayoutConstraintStartToStartOf                 string `xml:"layout_constraintStart_toStartOf,attr"`
		LayoutConstraintTopToTopOf                     string `xml:"layout_constraintTop_toTopOf,attr"`
		AndroidxConstraintlayoutWidgetConstraintLayout struct {
			Text         string `xml:",chardata"`
			LayoutWidth  string `xml:"layout_width,attr"`
			LayoutHeight string `xml:"layout_height,attr"`
			TextView     []struct {
				Text                           string `xml:",chardata"`
				ID                             string `xml:"id,attr"`
				LayoutWidth                    string `xml:"layout_width,attr"`
				LayoutHeight                   string `xml:"layout_height,attr"`
				LayoutMarginTop                string `xml:"layout_marginTop,attr"`
				AttrText                       string `xml:"text,attr"`
				TextSize                       string `xml:"textSize,attr"`
				TextStyle                      string `xml:"textStyle,attr"`
				LayoutConstraintLeftToLeftOf   string `xml:"layout_constraintLeft_toLeftOf,attr"`
				LayoutConstraintRightToRightOf string `xml:"layout_constraintRight_toRightOf,attr"`
				LayoutConstraintTopToBottomOf  string `xml:"layout_constraintTop_toBottomOf,attr"`
				LayoutMarginStart              string `xml:"layout_marginStart,attr"`
				LayoutMarginEnd                string `xml:"layout_marginEnd,attr"`
				TextAlignment                  string `xml:"textAlignment,attr"`
				LayoutConstraintEndToEndOf     string `xml:"layout_constraintEnd_toEndOf,attr"`
				LayoutConstraintStartToStartOf string `xml:"layout_constraintStart_toStartOf,attr"`
			} `xml:"TextView"`
			ImageView struct {
				Text                           string `xml:",chardata"`
				ID                             string `xml:"id,attr"`
				LayoutWidth                    string `xml:"layout_width,attr"`
				LayoutHeight                   string `xml:"layout_height,attr"`
				LayoutMarginTop                string `xml:"layout_marginTop,attr"`
				LayoutConstraintDimensionRatio string `xml:"layout_constraintDimensionRatio,attr"`
				LayoutConstraintEndToEndOf     string `xml:"layout_constraintEnd_toEndOf,attr"`
				LayoutConstraintStartToStartOf string `xml:"layout_constraintStart_toStartOf,attr"`
				LayoutConstraintTopToTopOf     string `xml:"layout_constraintTop_toTopOf,attr"`
				SrcCompat                      string `xml:"srcCompat,attr"`
			} `xml:"ImageView"`
			Button []struct {
				Text                           string `xml:",chardata"`
				ID                             string `xml:"id,attr"`
				LayoutWidth                    string `xml:"layout_width,attr"`
				LayoutHeight                   string `xml:"layout_height,attr"`
				LayoutMarginTop                string `xml:"layout_marginTop,attr"`
				OnClick                        string `xml:"onClick,attr"`
				AttrText                       string `xml:"text,attr"`
				LayoutConstraintEndToStartOf   string `xml:"layout_constraintEnd_toStartOf,attr"`
				LayoutConstraintStartToStartOf string `xml:"layout_constraintStart_toStartOf,attr"`
				LayoutConstraintTopToBottomOf  string `xml:"layout_constraintTop_toBottomOf,attr"`
				LayoutConstraintEndToEndOf     string `xml:"layout_constraintEnd_toEndOf,attr"`
				LayoutConstraintStartToEndOf   string `xml:"layout_constraintStart_toEndOf,attr"`
				LayoutConstraintTopToTopOf     string `xml:"layout_constraintTop_toTopOf,attr"`
			} `xml:"Button"`
		} `xml:"androidx.constraintlayout.widget.ConstraintLayout"`
	} `xml:"ScrollView"`
}

func TestXMLTextView(t *testing.T) {
	var androidXML Androidx
	xmlFile, _ := os.Open("test.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &androidXML)

	for _, val := range androidXML.ScrollView.AndroidxConstraintlayoutWidgetConstraintLayout.TextView {

		if val.AttrText == "Hello World!" {
			if val.LayoutMarginTop != "8dp" {
				t.Errorf("Wanted 8dp Got %v", val.LayoutMarginTop)
			}
			if val.TextStyle != "bold" {
				t.Errorf("Wanted bold Got %v", val.TextStyle)
			}
			if val.TextSize != "36sp" {
				t.Errorf("Wanted 36sp Got %v", val.TextSize)
			}
		} else {
			if val.LayoutMarginTop != "8dp" {
				t.Errorf("Wanted 8dp Got %v", val.LayoutMarginTop)
			}
			if val.LayoutMarginStart != "32dp" {
				t.Errorf("Wanted 32dp Got %v", val.LayoutMarginStart)
			}
			if val.LayoutMarginEnd != "32dp" {
				t.Errorf("Wanted 32dp Got %v", val.LayoutMarginEnd)
			}
			if val.TextAlignment != "center" {
				t.Errorf("Wanted center Got %v", val.TextAlignment)
			}
		}
	}
}

func TestXMLButton(t *testing.T) {
	var androidXML Androidx
	xmlFile, _ := os.Open("test.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &androidXML)

	for _, val := range androidXML.ScrollView.AndroidxConstraintlayoutWidgetConstraintLayout.Button {

		if val.AttrText == "Follow" {
			if val.LayoutMarginTop != "16dp" {
				t.Errorf("Wanted 16dp Got %v", val.LayoutMarginTop)
			}
		} else if val.AttrText == "Message" {

		} else {
			t.Errorf("Buttons incorrectly created!")
		}
	}
}

func TestXMLImage(t *testing.T) {
	var androidXML Androidx
	xmlFile, _ := os.Open("test.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &androidXML)

	val := androidXML.ScrollView.AndroidxConstraintlayoutWidgetConstraintLayout.ImageView

	if val.LayoutMarginTop != "32dp" {
		t.Errorf("Wanted 32dp Got %v", val.LayoutMarginTop)
	}
	if val.LayoutWidth != "128dp" {
		t.Errorf("Wanted 128dp Got %v", val.LayoutWidth)
	}
}
