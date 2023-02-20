package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }





*/

type FunFacts struct {
  Sun   []string
  Luna  []string
  Terra []string

}



 func  GetFunFacts(navn string ) []string{
  funFacts := FunFacts{

Luna: []string{
  "Temperatur på Månens overflate om natten -183 c", "Temperatur på Månens overflate om dagen 106 c",

},

Terra: []string{
  "Temperatur på ytre lag av Solen 5778°K", "Temperatur i Solens kjerne 15000000°C",

},

Sun: []string{
  "Temperatur på ytre lag av Solen 5778°K", "Temperatur i Solens kjerne 15000000°C",

},}


switch navn {

case "luna":
    return funFacts.Luna
case "terra":
  return funFacts.Terra

case "sun":
  return funFacts.Sun

default:
 return []string{"det er noen feil"}
}
}


