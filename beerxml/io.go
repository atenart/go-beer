/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

package beerxml

import (
	"encoding/xml"
	"os"
)

// Open a Beer XML formated file and returns a BeerXML object.
func Import(file string) (data *BeerXML, err error) {
	f, err := os.Open(file)
	if err != nil { return nil, err }
	defer f.Close()

	d := xml.NewDecoder(f)
	if err = d.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Write a BeerXML object into a file.
func Export(data *BeerXML, file string) error {
	f, err := os.Create(file)
	if err != nil { return err }
	defer f.Close()

	e := xml.NewEncoder(f)
	e.Indent("", "    ")
	if err = e.Encode(data); err != nil {
		return err
	}

	return nil
}
