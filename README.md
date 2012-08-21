The asign golang library parses a simple template language for creating Alpha sign communication protocol packets. This allows you to control Alpha LED Signs such as the BetaBrite.

# Example Application
Example application LEDSaid - https://github.com/krussell/ledsaid

## Simple Hello World Packet
	{SOT}{STX}{WriteText}{A}{Flash}{Red}Hello {Green}World!{ETX}{EOT}

## Multi-label Packet
	{SOT}
	{STX}
	{WriteText}{A}
	{Rotate}{Green}Here is some text for {RollUp}file label {Red}A
	{ETX}

	{STX}
	{WriteText}{B}
	{Spray}{Orange}Here is some text for {Snow}file label {Rainbow1}B
	{ETX}
	{EOT}

