package main

import (
	"math/rand"
	"strings"
	"time"
)

// Generate a string that is [n] words long.
// This string is generated from the static word list that is provided in this file. 
// Words are selected by random bag. 
// If the desired length is greater than the bag, the bag is reshuffled. 
func GenerateTypingPhrase(length int) string {

    shuffled := shuffleStringArray(wordlist)
    generated := make([]string, length)

    actualpos := 0 
    for i:= 0; i < length; i ++ {
        if actualpos >= len(shuffled) {
            shuffled = shuffleStringArray(shuffled)
            actualpos = 0 
        }
        generated[i] = shuffled[actualpos]
        actualpos = actualpos + 1
    }

	genString := ""
	for _, x := range generated {
		genString = genString + " " + x
	}

	return strings.TrimSpace(genString)
}

// 
func shuffleStringArray(inputArray []string) []string {
    r := rand.New(rand.NewSource(time.Now().UnixMicro()))
    out_arr := make([]string, len(inputArray))

    for i, val := range (inputArray) {
        out_arr[i] = val 
    }
    
    for i := 0;  i < len(out_arr); i++ {
        randInt := r.Intn(len(out_arr))
        swap := out_arr[i]
        out_arr[i] = out_arr[randInt] 
        out_arr[randInt] = swap 
    }

    return out_arr 
}

//200 most common english words.
//TODO replace this with a text file that is loaded into memory
var wordlist = []string{
	"the", "be", "of", "and", "a", "to", "in", "he", "have", "it", "that", "for", "they", "I", "with", "as",
	"not", "on", "she", "at", "by", "this", "we", "you", "do", "but", "from", "or", "which", "one", "would", "all",
	"will", "there", "say", "who", "make", "when", "can", "more", "if", "no", "man", "out", "other", "so", "what", "time",
	"up", "go", "about", "than", "into", "could", "state", "only", "new", "year", "some", "take", "come", "these", "know", "see",
	"use", "get", "like", "then", "first", "any", "work", "now", "may", "such", "give", "over", "think", "most", "even", "find", "day", "also", "after",
	"way", "many", "must", "look", "before", "great", "back", "through", "long", "where", "much", "should", "well", "people", "down", "own",
	"just", "because", "good", "each", "those", "feel", "seem", "how", "high", "too", "place", "little", "world", "very", "still", "nation",
	"hand", "old", "life", "tell", "write", "become", "here", "show", "house", "both", "between", "need", "mean", "call", "develop", "under",
	"last", "right", "move", "thing", "general", "school", "never", "same", "another", "begin", "while", "number", "part",
	"turn", "real", "leave", "might", "want", "point", "form", "off", "child", "few", "small", "since", "against", "ask", "late", "home",
	"interest", "large", "person", "end", "open", "public", "follow", "during", "present", "without", "again", "hold", "govern", "around", "possible", "head",
	"consider", "word", "program", "problem", "however", "lead", "system", "set", "order", "eye", "plan", "run", "keep",
	"face", "fact", "group", "play", "stand", "increase", "early", "course", "change", "help", "line",
}
