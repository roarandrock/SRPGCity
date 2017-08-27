package models

var (
	q1 = "This is not a load bearing sandwich"
	q2 = "The ravenous sorrow of a broken heart"
	q3 = "Choose life"
	q4 = "You are not who you say you are. Including to the mirror."
	q5 = "We live such lives of luxury. When we feel said it must be tragic. That's why we listen to the blues."
	q6 = "We fight boredom, loneliness and the creeping twins realizations that we are both meaningless dust" +
		"and resource hogging lords of this planet."
	q7  = "Your energy level is psychosomatic and not empirical"
	q8  = "It's always better to do things in a couple. LIke twins."
	q9  = "Arrogance is not your friend"
	q10 = "The problem is everyone thinks they are the heroes of their story. We cannot all be the heroes." +
		"Someone has to be the lackey, the fool and the villain."
	qCount = 10
)

//Quote1 defines the quote
type Quote1 struct {
	text      string
	situation []int //can fit a variety of situations, intro, romance, work
}
