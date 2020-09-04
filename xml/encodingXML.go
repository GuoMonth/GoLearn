package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

func exampleMarshalIndent() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		ID        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	//   <person id="13">
	//       <name>
	//           <first>John</first>
	//           <last>Doe</last>
	//       </name>
	//       <age>42</age>
	//       <Married>false</Married>
	//       <City>Hanga Roa</City>
	//       <State>Easter Island</State>
	//       <!-- Need more details. -->
	//   </person>
}

func writeXML() {
	filePath := "C:\\files\\test.xml"
	w, err := os.Create(filePath)
	defer w.Close()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err = xml.EscapeText(w, []byte("\n<A>nice ttto meeto</A>"))
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
}
