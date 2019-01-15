package metric

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Millimeter
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MMToCM converts millimeters to centimeters.
func MMToCM(mm Millimeter) Centimeter { return Centimeter(mm * 0.1) }

// MMToM converts millimeters to meters.
func MMToM(mm Millimeter) Meter { return Meter(mm * 0.001) }

// MMToKM converts millimeters to kilometers.
func MMToKM(mm Millimeter) Kilometer { return Kilometer(mm * 1.0E-6) }

// MMToIN converts millimeters to inches,
func MMToIN(mm Millimeter) Inch { return Inch(mm * 0.039370078740157) }

// MMToFT converts millimeters to feet,
func MMToFT(mm Millimeter) Foot { return Foot(mm * 0.0032808398950131) }

// MMToYD converts millimeters to yards.
func MMToYD(mm Millimeter) Yard { return Yard(mm * 0.0010936132983377) }

// MMToMI converts millimeters to miles.
func MMToMI(mm Millimeter) Mile { return Mile(mm * 6.2137119223733E-7) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Centimeters
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// CMToMM converts centimeters to millimeters.
func CMToMM(cm Centimeter) Millimeter { return Millimeter(cm * 10) }

// CMToM converts centimeters to meters.
func CMToM(cm Centimeter) Meter { return Meter(cm * 0.001) }

// CMToKM converts centimeters to kilometers.
func CMToKM(cm Centimeter) Kilometer { return Kilometer(cm * 1.0E-6) }

// CMToIN converts centimeters to inches.
func CMToIN(cm Centimeter) Inch { return Inch(cm * 0.39370078740157) }

// CMToFT converts centimeters to feet.
func CMToFT(cm Centimeter) Foot { return Foot(cm * 0.032808398950131) }

// CMToYD converts centimeters to yards.
func CMToYD(cm Centimeter) Yard { return Yard(cm * 0.010936132983377) }

// CMToMI converts centimeters to miles.
func CMToMI(cm Centimeter) Mile { return Mile(cm * 6.2137119223733E-6) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Meters
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MToMM converts meters to millimeters.
func MToMM(m Meter) Millimeter { return Millimeter(m * 1000) }

// MToCM converts meters to centimeters.
func MToCM(m Meter) Centimeter { return Centimeter(m * 100) }

// MToKM converts meters to kilometers.
func MToKM(m Meter) Kilometer { return Kilometer(m * 0.001) }

// MToIN converts meters to inches.
func MToIN(m Meter) Inch { return Inch(m * 39.370078740157) }

// MToFT converts meters to feet.
func MToFT(m Meter) Foot { return Foot(m * 3.2808398950131) }

// MToYD converts meters to yards.
func MToYD(m Meter) Yard { return Yard(m * 1.0936132983377) }

// MToMI converts meters to miles.
func MToMI(m Meter) Mile { return Mile(m * 0.00062137119223733) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Kilometers
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// KMToMM converts kilometers to millimeters.
func KMToMM(km Kilometer) Millimeter { return Millimeter(km * 1000000) }

// KMToCM converts kilometers to centimeters.
func KMToCM(km Kilometer) Centimeter { return Centimeter(km * 100000) }

// KMToM converts kilometers to meters.
func KMToM(km Kilometer) Meter { return Meter(km * 1000) }

// KMToIN converts kilometers to inches.
func KMToIN(km Kilometer) Inch { return Inch(km * 39370) }

// KMToFT converts kilometers to feet.
func KMToFT(km Kilometer) Foot { return Foot(km * 3281) }

// KMToYD converts kilometers to yards.
func KMToYD(km Kilometer) Yard { return Yard(km * 1072.666666666667) }

// KMToMI converts kilometers to miles.
func KMToMI(km Kilometer) Mile { return Mile(km * 0.6214) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Inches
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// INToMM converts inches to millimeters.
func INToMM(in Inch) Millimeter { return Millimeter(in * 25.4) }

// INToCM converts inches to centimeters.
func INToCM(in Inch) Centimeter { return Centimeter(in * 2.54) }

// INToM converts inches to meters.
func INToM(in Inch) Meter { return Meter(in * 0.0254) }

// INToKM converts inches to kilometers.
func INToKM(in Inch) Kilometer { return Kilometer(in * 2.54E-5) }

// INToFT converts inches to feet.
func INToFT(in Inch) Foot { return Foot(in * 0.083333333333333) }

// INToYD converts inches to yards.
func INToYD(in Inch) Yard { return Yard(in * 0.027777777777778) }

// INToMI converts inches to miles.
func INToMI(in Inch) Mile { return Mile(in * 1.5782828282828E-5) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Foot
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FTToMM converts feet to millimeters.
func FTToMM(ft Foot) Millimeter { return Millimeter(ft * 304.8) }

// FTToCM converts feet to centimeters.
func FTToCM(ft Foot) Centimeter { return Centimeter(ft * 30.48) }

// FTToM converts feet to meters.
func FTToM(ft Foot) Meter { return Meter(ft * 0.3048) }

// FTToKM converts feet to kilometers.
func FTToKM(ft Foot) Kilometer { return Kilometer(ft * 0.0003048) }

// FTToIN converts feet to inches.
func FTToIN(ft Foot) Inch { return Inch(ft * 12) }

// FTToYD converts feet to yards.
func FTToYD(ft Foot) Yard { return Yard(ft * 0.33333333333333) }

// FTToMI converts feet to mile.
func FTToMI(ft Foot) Mile { return Mile(ft * 0.00018939393939394) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Yard
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// YDToMM converts yards to millimeters.
func YDToMM(yd Yard) Millimeter { return Millimeter(yd * 914.4) }

// YDToCM converts yards to centimeters.
func YDToCM(yd Yard) Centimeter { return Centimeter(yd * 91.44) }

// YDToM converts yards to meters.
func YDToM(yd Yard) Meter { return Meter(yd * 0.9144) }

// YDToKM converts yards to kilometers.
func YDToKM(yd Yard) Kilometer { return Kilometer(yd * 0.0009144) }

// YDToIN converts yards to inches.
func YDToIN(yd Yard) Inch { return Inch(yd * 36) }

// YDToFT converts yards to feet.
func YDToFT(yd Yard) Foot { return Foot(yd * 3) }

// YDToMI converts yards to miles.
func YDToMI(yd Yard) Mile { return Mile(yd * 0.00056818181818182) }

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//  Mile
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MIToMM converts miles to millimeters.
func MIToMM(mi Mile) Millimeter { return Millimeter(mi * 1609344) }

// MIToCM converts miles to centimeters.
func MIToCM(mi Mile) Centimeter { return Centimeter(mi * 160934.4) }

// MIToM converts miles to meters.
func MIToM(mi Mile) Meter { return Meter(mi * 1609.344) }

// MIToKM converts miles to kilometers.
func MIToKM(mi Mile) Kilometer { return Kilometer(mi * 1.609344) }

// MIToIN converts miles to inches.
func MIToIN(mi Mile) Inch { return Inch(mi * 63360) }

// MIToFT converts miles to feet.
func MIToFT(mi Mile) Foot { return Foot(mi * 5280) }

// MIToYD converts miles to yards.
func MIToYD(mi Mile) Yard { return Yard(mi * 1760) }
