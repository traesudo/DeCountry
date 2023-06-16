# DeCountry

###
simply import it into your Golang project and call the IsCountryName function with the string you wish to check as its parameter. The function will return a boolean value indicating whether the string is a valid country name or abbreviation.

For example:

```
import "github.com/DeCountry"

func main() {
   isValid := DeCountry.IsCountryName("Japan")
   fmt.Println(isValid) // Output: true
}

```
