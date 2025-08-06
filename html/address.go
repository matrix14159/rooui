package html

type AddressElement struct {
	BaseElement
}

func Address() *AddressElement {
	return &AddressElement{InitBaseElement("address")}
}
