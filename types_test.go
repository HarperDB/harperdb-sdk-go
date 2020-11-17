package harperdb

func ExampleRecord() {

	type Player struct {
		Record // embed the Record type in your custom types

		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

}
