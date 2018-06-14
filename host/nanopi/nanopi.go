/*
	Package NanoPi provides NanoPi Duo support (other boards were not tested).
	The following features are supported on Linux kernel 4.14+ (lower versions were not tested)

	GPIO (digital (rw))
	I²C
    SPI
	LED
*/
package nanopi

import (
"github.com/kidoman/embd"
"github.com/kidoman/embd/host/generic"
)

var spiDeviceMinor = 1

var duoPins = embd.PinMap{
	&embd.PinDesc{ID: "P1_6", Aliases: []string{"198", "GPIOG6", "TX1", "UART1_TX"}, Caps: embd.CapDigital | embd.CapUART, DigitalLogical: 198},
	&embd.PinDesc{ID: "P1_7", Aliases: []string{"199", "GPIOG7", "RX1", "UART1_RX"}, Caps: embd.CapDigital | embd.CapUART, DigitalLogical: 199},
	&embd.PinDesc{ID: "P1_8", Aliases: []string{"15", "GPIOA15", "MOSI", "SPI1_MOSI"}, Caps: embd.CapDigital | embd.CapSPI, DigitalLogical: 15},
	&embd.PinDesc{ID: "P1_9", Aliases: []string{"16", "GPIOA16", "MISO", "SPI1_MISO"}, Caps: embd.CapDigital | embd.CapSPI, DigitalLogical: 16},
	&embd.PinDesc{ID: "P1_10", Aliases: []string{"14", "GPIOA14", "SCLK", "SPI1_CLK"}, Caps: embd.CapDigital | embd.CapSPI, DigitalLogical: 14},
	&embd.PinDesc{ID: "P1_11", Aliases: []string{"13", "GPIOA13", "SCS", "SPI1_CS"}, Caps: embd.CapDigital | embd.CapSPI, DigitalLogical: 13},
	&embd.PinDesc{ID: "P1_12", Aliases: []string{"12", "GPIOA12", "SDA", "I2C0_SDA"}, Caps: embd.CapDigital | embd.CapI2C, DigitalLogical: 12},
	&embd.PinDesc{ID: "P1_13", Aliases: []string{"11", "GPIOA11", "SCL", "I2C0_SCL"}, Caps: embd.CapDigital | embd.CapI2C, DigitalLogical: 11},
	&embd.PinDesc{ID: "P1_15", Aliases: []string{"4", "GPIOA4", "DEBUG_TX", "UART_TXD0"}, Caps: embd.CapDigital | embd.CapUART, DigitalLogical: 4},
	&embd.PinDesc{ID: "P1_16", Aliases: []string{"5", "GPIOA5", "DEBUG_RX", "UART_RXD0"}, Caps: embd.CapDigital | embd.CapUART, DigitalLogical: 5},

	&embd.PinDesc{ID: "P2_11", Aliases: []string{"203", "GPIOG11", "IOG11"}, Caps: embd.CapDigital, DigitalLogical: 203},
	&embd.PinDesc{ID: "P2_12", Aliases: []string{"363", "GPIOL11", "IRRX"}, Caps: embd.CapDigital, DigitalLogical: 363},
}

var ledMap = embd.LEDMap{
	"led0": []string{"0", "led0", "LED0"},
}

func init() {
	embd.Register(embd.HostCHIP, func(rev int) *embd.Descriptor {
		// Refer to http://elinux.org/RPi_HardwareHistory#Board_Revision_History
		// for details.
		pins := duoPins

		return &embd.Descriptor{
			GPIODriver: func() embd.GPIODriver {
				return embd.NewGPIODriver(pins, generic.NewDigitalPin, nil, nil)
			},
			I2CDriver: func() embd.I2CDriver {
				return embd.NewI2CDriver(generic.NewI2CBus)
			},
			LEDDriver: func() embd.LEDDriver {
				return embd.NewLEDDriver(ledMap, generic.NewLED)
			},
			SPIDriver: func() embd.SPIDriver {
				return embd.NewSPIDriver(spiDeviceMinor, generic.NewSPIBus, nil)
			},
		}
	})
}

