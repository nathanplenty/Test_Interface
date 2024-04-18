package main

import (
	"log"
)

type ElectricalDevice interface {
	TurnOn()
	TurnOff()
}

type Device struct {
	Type  string
	Model string
}

func (d *Device) TurnOn() {
	// Funktion erhält die Information des Gerätes und führt damit was aus
	log.Println(d.Type, d.Model, "eingeschaltet")
}

func (d *Device) TurnOff() {
	// Funktion erhält die Information des Gerätes und führt damit was aus
	log.Println(d.Type, d.Model, "ausgeschaltet")
}

func PlugIn(device ElectricalDevice) {
	// Funktion ist hier ein Verb, etwas wird gemacht
	// Funktion nimmt das Gerät als Variable
	log.Println("Gerät an die Steckdose angeschlossen")
	// Gerät wird an die Funktion TurnOn geschickt
	device.TurnOn()
	device.TurnOff()
	// Gerät wird an die Funktion TurnOff geschickt
	log.Println("Gerät von der Steckdose getrennt")
}

func main() {
	// Toaster wird erstellt
	toaster := &Device{Type: "Toaster", Model: "FastToast"}
	log.Println("Toaster wird angeschlossen...")
	// Plugin Funktion wird mit Toaster ausgeführt
	PlugIn(toaster)

	log.Println()

	// Kühschrank wird erstellt
	refrigerator := &Device{Type: "Kühlschrank", Model: "KühlerX"}
	log.Println("Kühlschrank wird angeschlossen...")
	// Plugin Funktion wird mit Kühlschrank ausgeführt
	PlugIn(refrigerator)
}
