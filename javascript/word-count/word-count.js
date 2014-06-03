// personal note - docs on reduce found here: 
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/Reduce
var words = module.exports = function(sentence) {
	
	// Normalize case.
	sentence = sentence.toLowerCase();

	// Remove non alphabetic symbols.
	sentence = sentence.match(/\w+/g);

	// Check for word matches increment word counts if found.
	sentence = sentence.reduce(function(wordCounts,word) {
		wordCounts.hasOwnProperty(word) ? wordCounts[word]++ : wordCounts[word] = 1;
		return wordCounts;
	}, {});

	return sentence;	
} 