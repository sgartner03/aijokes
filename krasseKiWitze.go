package KrasseKiWitze

import (
	"math/rand"
)

var witze = []string{
	"1. Warum sind Bankräuber so gut im Tennis? Weil sie auf jeden Schlag vorbereitet sind!",
	"2. Warum lieben Schafe Mathe? Weil sie die ganze Zeit am Wollzählen sind!",
	"3. Was macht ein Pirat beim Yoga? Er geht in die Plank-Position!",
	"4. Warum sind Informatiker schlecht im Fußball? Weil sie jeden Ball neu erfinden!",
	"5. Was passiert, wenn ein Kühlschrank einen Witz erzählt? Es ist ein Eisbrecher!",
	"6. Warum können Geister keine Lügen erzählen? Weil man durch sie hindurchsieht!",
	"7. Warum sollte man nie Schach mit einem Baum spielen? Weil er jeden Zug verwurzelt!",
	"8. Was ist das Lieblingsessen eines Elektrikers? Spannungskartoffeln!",
	"9. Was macht ein Mathematiker, wenn er seekrank wird? Er errechnet den Kotzvektor!",
	"10. Warum hatte der Schmetterling",
}

func BringMichZumLachen() string {
	N := len(witze)
	return witze[rand.Intn(N)]
}
