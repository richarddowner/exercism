// personal note - docs on reduce found here: 
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/Reduce
var words = module.exports = function(sentence) {
	
	// Normalize case.
	sentence = sentence.toLowerCase();

	// Remove non alphabetic symbols.
	// attempt at non-english regex : sentence = sentence.match(/\b([a-z\xE0-\xFF\x0-9])+/g);
	sentence = sentence.match(/\w+/g);
	
	// Check for word matches increment word counts if found.
	sentence = sentence.reduce(function(wordCounts,word) {
		wordCounts.hasOwnProperty(word) ? wordCounts[word]++ : wordCounts[word] = 1;
		return wordCounts;
	}, {});

	return sentence;
} 