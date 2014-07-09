
package main

import "fmt"
import "nmea"

func main() {
  sent := "$GPGSA,A,3,,,,,,16,18,,22,24,,,3.6,2.1,2.2*3C"
  s, err := nmea.Parse(sent)
  if err != nil {
    fmt.Printf("Error parsing: %s\n", err)
    return
  }

  switch t := s.(type) {
  case nmea.GPGGA:
    fmt.Printf("A GPGGA: %T\n", t)
  case nmea.GPGSA:
    fmt.Printf("A GPGSA: %T\n", t)
  }

  g, ok := s.(nmea.GPGSA)
  if !ok {
    fmt.Printf("not a GPGSA, it's a %T\n", s)
    return
  }

  fmt.Printf("%s\n", g.Raw)
  fmt.Printf("> %s\n", g.SType)
  fmt.Printf("%q\n", g.Fields)

  fmt.Printf("Mode: %s/%s\n", g.Mode, g.FixType)
  fmt.Printf("SVs: %v\n", g.SV)
  fmt.Printf("PDOP: %s\n", g.PDOP)
  fmt.Printf("HDOP: %s\n", g.HDOP)
  fmt.Printf("VDOP: %s\n", g.VDOP)
}
