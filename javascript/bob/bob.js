var Bob = module.exports = function() {
	this.hey = function(words) {
		if(words.trim() === '') {
			return "Fine. Be that way!";
		}
		if(words === words.toUpperCase() && words.toUpperCase() != words.toLowerCase()) {
			return "Woah, chill out!";
		}
		if(words[words.length - 1] === '?') {
			return "Sure.";
		}
		return "Whatever.";
	}
};