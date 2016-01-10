package magnaporters

import (
  "fmt"
)

func PortersAlgorithm(aWord string) (aStem string, measure int){
  measure = FindMeasure(aWord)
  aStem = phaseOne(aWord, measure)
  aStem = phaseTwo(aStem, measure)
  aStem = phaseThree(aStem, measure)
  aStem = phaseFour(aStem, measure)
  aStem = phaseFive(aStem, measure)
  return aStem, measure
}

//FindMeasure function will calculate the measure of a word defined by Porters algorithm.
//The function takes in a string and will return the integer value of the meaure.
func FindMeasure(aWord string) int {
	measureRegexp := regexp.MustCompile("[aeiou][^aeiou]")
	aMeasure := len(measureRegexp.FindAllString(aWord, -1))
	return aMeasure
}

func phaseOne(aStem string, wordMeasure int) string {

	ssesRegexp := regexp.MustCompile(".*sses$")
	if len(ssesRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = ssesRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ss")
	}
	iesRegexp := regexp.MustCompile(".*ies$")
	if len(iesRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = iesRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"i")
	}
	ssRegexp := regexp.MustCompile(".*ss$")
	if len(ssRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = ssRegexp.ReplaceAllString(aStem, aStem)
	}
	sRegexp := regexp.MustCompile(".*s$")
	if len(sRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = sRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
	}
	p1aRegexp := regexp.MustCompile(consanantRegexp + ".*" + vowelRegexp + "eed$")
	if len(p1aRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = p1aRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"ee")
	}
	p1bRegexp := regexp.MustCompile(".*[aeiou]*(ed$|ing$)")
	if len(p1bRegexp.FindAllString(aStem, 1)) > 0 {
		edRegexp := regexp.MustCompile(".*[aeiou].*ed$")
		if len(edRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = edRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		}
		ingRegexp := regexp.MustCompile(".*[aeiou].*ing$")
		if len(ingRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ingRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}

		atRegexp := regexp.MustCompile(".*[aeiou].*at$")
		if len(atRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = atRegexp.ReplaceAllString(aStem, aStem+"e")
		}
		blRegexp := regexp.MustCompile(".*[aeiou].*bl$")
		if len(blRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = blRegexp.ReplaceAllString(aStem, aStem+"e")
		}
		izRegexp := regexp.MustCompile(".*[aeiou].*iz$")
		if len(izRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = izRegexp.ReplaceAllString(aStem, aStem+"e")
		}

	}

	yRegexp := regexp.MustCompile(".*[^aeiou]y$")
	if len(yRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = yRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1]+"i")
	}

	return aStem
}

func phaseTwo(aStem string, wordMeasure int) string {
	if wordMeasure > 0 {
		ationalRegexp := regexp.MustCompile(".*[aeiou].*ational$")
		if len(ationalRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ationalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ate")
		}
		tionalRegexp := regexp.MustCompile(".*[aeiou].*tional$")
		if len(tionalRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = tionalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-6]+"tion")
		}
		enciRegexp := regexp.MustCompile(".*[aeiou].*enci$")
		if len(enciRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = enciRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ence")
		}
		anciRegexp := regexp.MustCompile(".*[aeiou].*anci$")
		if len(anciRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = anciRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ance")
		}
		izerRegexp := regexp.MustCompile(".*[aeiou].*izer$")
		if len(izerRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = izerRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ize")
		}
		abliRegexp := regexp.MustCompile(".*[aeiou].*abli$")
		if len(abliRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = abliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"able")
		}
		alliRegexp := regexp.MustCompile(".*[aeiou].*alli$")
		if len(alliRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = alliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"al")
		}
		entliRegexp := regexp.MustCompile(".*[aeiou].*entli$")
		if len(entliRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = entliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ent")
		}
		eliRegexp := regexp.MustCompile(".*[aeiou].*eli$")
		if len(eliRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = eliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"e")
		}
		ousliRegexp := regexp.MustCompile(".*[aeiou].*ousli$")
		if len(ousliRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ousliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ous")
		}
		izationRegexp := regexp.MustCompile(".*[aeiou].*ization$")
		if len(izationRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = izationRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ize")
		}
		ationRegexp := regexp.MustCompile(".*[aeiou].*ation$")
		if len(ationRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ationRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ate")
		}
		atorRegexp := regexp.MustCompile(".*[aeiou].*ator$")
		if len(atorRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = atorRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ate")
		}
		alismRegexp := regexp.MustCompile(".*[aeiou].*alism$")
		if len(alismRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = alismRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		}
		ivenessRegexp := regexp.MustCompile(".*[aeiou].*iveness$")
		if len(ivenessRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ivenessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ive")
		}
		fulnessRegexp := regexp.MustCompile(".*[aeiou].*fulness$")
		if len(fulnessRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = fulnessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ful")
		}
		ousnessRegexp := regexp.MustCompile(".*[aeiou].*ousness$")
		if len(ousnessRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ousnessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ous")
		}
		alitiRegexp := regexp.MustCompile(".*[aeiou].*aliti$")
		if len(alitiRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = alitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		}
		ivitiRegexp := regexp.MustCompile(".*[aeiou].*iviti$")
		if len(ivitiRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ivitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ive")
		}
		bilitiRegexp := regexp.MustCompile(".*[aeiou].*biliti$")
		if len(bilitiRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = bilitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-6]+"ble")
		}
	}
	return aStem
}

func phaseThree(aStem string, wordMeasure int) string {
	if wordMeasure > 0 {
		icateRegexp := regexp.MustCompile(".*[aeiou].*icate$")
		if len(icateRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = icateRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ic")
		}
		ativeRegexp := regexp.MustCompile(".*[aeiou].*ative$")
		if len(ativeRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ativeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5])
		}
		alizeRegexp := regexp.MustCompile(".*[aeiou].*alize$")
		if len(alizeRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = alizeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		}
		icitiRegexp := regexp.MustCompile(".*[aeiou].*iciti$")
		if len(icitiRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = icitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ic")
		}
		icalRegexp := regexp.MustCompile(".*[aeiou].*ical$")
		if len(icalRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = icalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ic")
		}
		fulRegexp := regexp.MustCompile(".*[aeiou].*ful$")
		if len(fulRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = fulRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		nessRegexp := regexp.MustCompile(".*[aeiou].*ness$")
		if len(nessRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = nessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
	}
	return aStem
}

func phaseFour(aStem string, wordMeasure int) string {

	if wordMeasure > 1 {
		alRegexp := regexp.MustCompile(".*[aeiou].*al$")
		if len(alRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = alRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		}
		anceRegexp := regexp.MustCompile(".*[aeiou].*ance$")
		if len(anceRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = anceRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		enceRegexp := regexp.MustCompile(".*[aeiou].*ence$")
		if len(enceRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = enceRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		erRegexp := regexp.MustCompile(".*[aeiou].*er$")
		if len(erRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = erRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		}
		icRegexp := regexp.MustCompile(".*[aeiou].*ic$")
		if len(icRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = icRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		}
		ableRegexp := regexp.MustCompile(".*[aeiou].*able$")
		if len(ableRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ableRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		ibleRegexp := regexp.MustCompile(".*[aeiou].*ible$")
		if len(ibleRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ibleRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		antRegexp := regexp.MustCompile(".*[aeiou].*ant$")
		if len(antRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = antRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		ementRegexp := regexp.MustCompile(".*[aeiou].*ement$")
		if len(ementRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ementRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5])
		}
		mentRegexp := regexp.MustCompile(".*[aeiou].*ment$")
		if len(mentRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = mentRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		entRegexp := regexp.MustCompile(".*[aeiou].*ent$")
		if len(entRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = entRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}

		s_tionRegexp := regexp.MustCompile(".*[aeiou].*[s|t]ion$")
		if len(s_tionRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = s_tionRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		}
		ouRegexp := regexp.MustCompile(".*[aeiou].*ou$")
		if len(ouRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ouRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		}
		ismRegexp := regexp.MustCompile(".*[aeiou].*ism$")
		if len(ismRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ismRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		ateRegexp := regexp.MustCompile(".*[aeiou].*ate$")
		if len(ateRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ateRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		itiRegexp := regexp.MustCompile(".*[aeiou].*iti$")
		if len(itiRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = itiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		ousRegexp := regexp.MustCompile(".*[aeiou].*ous$")
		if len(ousRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = ousRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		iveRegexp := regexp.MustCompile(".*[aeiou].*ive$")
		if len(iveRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = iveRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
		izeRegexp := regexp.MustCompile(".*[aeiou].*ize$")
		if len(izeRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = izeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		}
	}
	return aStem
}

func phaseFive(aStem string, wordMeasure int) string {
	if wordMeasure > 1 {
		eRegexp := regexp.MustCompile(".*[aeiou].*e$")
		if len(eRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = eRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		}
	}
	if wordMeasure == 1 {
		eRegexp := regexp.MustCompile(".*[aeiou].*[^o]e$")
		if len(eRegexp.FindAllString(aStem, 1)) > 0 {
			aStem = eRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		}
	}
	if wordMeasure > 1 {
		lRegexp := regexp.MustCompile(".*[aeiou].*([^dL])$")
		//find second to last letter and see if they match the last letter if true
		if len(lRegexp.FindAllString(aStem, 1)) > 0 {
			lastLetter := aStem[len(aStem)-1]
			secondToLastLetter := aStem[len(aStem)-2]
			if lastLetter == secondToLastLetter {
				aStem = lRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
			}
		}
	}

	return aStem
}
