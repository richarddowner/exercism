var words = module.exports = function(sentence) {
	return sentence
	.toLowerCase()
	.match(/\w+/g)
	.reduce(function(wordCounts,word) {
		wordCounts.hasOwnProperty(word) ? wordCounts[word]++ : wordCounts[word] = 1;
		return wordCounts;
	}, {});	
} 