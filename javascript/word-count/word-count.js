var words = module.exports = function(sentence) {
	return sentence.toLowerCase()
	.replace(/([!:@\$%\^&,¡\?¿]+)/g, "")
	.replace(/\s+/g, " ")
	.split(" ")
	.reduce(function(wordCounts,word) {
		wordCounts.hasOwnProperty(word) ? wordCounts[word]++ : wordCounts[word] = 1;
		return wordCounts;
	}, {});	
} 