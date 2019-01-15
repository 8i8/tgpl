package metric

import "fmt"

type Kilometer float64
type Meter float64
type Centimeter float64
type Millimeter float64
type Inch float64
type Foot float64
type Yard float64
type Mile float64

func (k Kilometer) String() string   { return fmt.Sprintf("%.2gkm", k) }
func (m Meter) String() string       { return fmt.Sprintf("%.2gm", m) }
func (cm Centimeter) String() string { return fmt.Sprintf("%.2gcm", cm) }
func (mm Millimeter) String() string { return fmt.Sprintf("%.2gmm", mm) }
func (in Inch) String() string       { return fmt.Sprintf("%.2g\"", in) }
func (ft Foot) String() string       { return fmt.Sprintf("%.2g'", ft) }
func (yd Yard) String() string       { return fmt.Sprintf("%.2gyd", yd) }
func (mi Mile) String() string       { return fmt.Sprintf("%.2gmi", mi) }
