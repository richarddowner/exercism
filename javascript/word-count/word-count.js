var words = function(stringOfWords) {
	var wordCountCollection = {};

	// Remove non alphabetic symbols and extra spaces.
	stringOfWords = stringOfWords.replace(/[^\w\s]|_/g, "").replace(/\s+/g, " ");
	
	// Normalize case
	stringOfWords = stringOfWords.toLowerCase();

	// Split the words out of the string.
	var wordCollection = stringOfWords.split(" ");
	
	// Check for word matches increment if found.
	for (i = 0; i< wordCollection.length; i++) {
		if(wordCountCollection[wordCollection[i]] == null) {
			wordCountCollection[wordCollection[i]] = 1;
		} else {
			wordCountCollection[wordCollection[i]] += 1;	
		}		
	}	

	return wordCountCollection;
}
module.exports = words;