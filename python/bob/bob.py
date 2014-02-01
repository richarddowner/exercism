class Bob:
	def hey(self, sentence):
		# if sentence is empty or whitespace
		if not sentence:
			return 'Fine. Be that way!'
		# if sentence ends with ?
		if sentence[len(sentence)-1] == '?':
			return 'Sure.'
		# if sentence is all uppercase
		if sentence == sentence.upper():
			return 'Woah, chill out!'
		# else return default
		return 'Whatever.'
