package rand2

// Source: https://github.com/4eek/gofaker/blob/master/lorem/lorem.go

import (
	"math/rand"
	"strings"
)

var words []string = []string{
	"alias", "consequatur", "aut", "perferendis", "sit", "voluptatem",
	"accusantium", "doloremque", "aperiam", "eaque", "ipsa", "quae", "ab",
	"illo", "inventore", "veritatis", "et", "quasi", "architecto",
	"beatae", "vitae", "dicta", "sunt", "explicabo", "aspernatur", "aut",
	"odit", "aut", "fugit", "sed", "quia", "consequuntur", "magni",
	"dolores", "eos", "qui", "ratione", "voluptatem", "sequi", "nesciunt",
	"neque", "dolorem", "ipsum", "quia", "dolor", "sit", "amet",
	"consectetur", "adipisci", "velit", "sed", "quia", "non", "numquam",
	"eius", "modi", "tempora", "incidunt", "ut", "labore", "et", "dolore",
	"magnam", "aliquam", "quaerat", "voluptatem", "ut", "enim", "ad",
	"minima", "veniam", "quis", "nostrum", "exercitationem", "ullam",
	"corporis", "nemo", "enim", "ipsam", "voluptatem", "quia", "voluptas",
	"sit", "suscipit", "laboriosam", "nisi", "ut", "aliquid", "ex", "ea",
	"commodi", "consequatur", "quis", "autem", "vel", "eum", "iure",
	"reprehenderit", "qui", "in", "ea", "voluptate", "velit", "esse",
	"quam", "nihil", "molestiae", "et", "iusto", "odio", "dignissimos",
	"ducimus", "qui", "blanditiis", "praesentium", "laudantium", "totam",
	"rem", "voluptatum", "deleniti", "atque", "corrupti", "quos",
	"dolores", "et", "quas", "molestias", "excepturi", "sint",
	"occaecati", "cupiditate", "non", "provident", "sed", "ut",
	"perspiciatis", "unde", "omnis", "iste", "natus", "error",
	"similique", "sunt", "in", "culpa", "qui", "officia", "deserunt",
	"mollitia", "animi", "id", "est", "laborum", "et", "dolorum", "fuga",
	"et", "harum", "quidem", "rerum", "facilis", "est", "et", "expedita",
	"distinctio", "nam", "libero", "tempore", "cum", "soluta", "nobis",
	"est", "eligendi", "optio", "cumque", "nihil", "impedit", "quo",
	"porro", "quisquam", "est", "qui", "minus", "id", "quod", "maxime",
	"placeat", "facere", "possimus", "omnis", "voluptas", "assumenda",
	"est", "omnis", "dolor", "repellendus", "temporibus", "autem",
	"quibusdam", "et", "aut", "consequatur", "vel", "illum", "qui",
	"dolorem", "eum", "fugiat", "quo", "voluptas", "nulla", "pariatur",
	"at", "vero", "eos", "et", "accusamus", "officiis", "debitis", "aut",
	"rerum", "necessitatibus", "saepe", "eveniet", "ut", "et",
	"voluptates", "repudiandae", "sint", "et", "molestiae", "non",
	"recusandae", "itaque", "earum", "rerum", "hic", "tenetur", "a",
	"sapiente", "delectus", "ut", "aut", "reiciendis", "voluptatibus",
	"maiores", "doloribus", "asperiores", "repellat"}

// Word returns a random word.
func Word() string {
	return words[rand.Intn(len(words))]
}

// Words returns a string containing randoms words.
// The passed-in count controls how many words are used.
func Words(count int) (words string) {
	if count < 1 {
		return ""
	}
	words = Word()
	for i := 1; i < count; i++ {
		words = words + " " + Word()
	}
	return
}

// Sentence returns a string with multiple words, formatted like an English sentence.
// The passed-in count controls how many words are used.
func Sentence(count int) (sentence string) {
	if count < 1 {
		return ""
	}
	sentence = strings.Title(Word())
	for i := 1; i < count; i++ {
		sentence = sentence + " " + Word()
	}
	sentence = strings.TrimSpace(sentence) + "."
	return
}

// Sentences returns a string of multiple sentences.
// The passed-in count controls the number of sentences,
// words defines how many words per sentence are used.
func Sentences(count, words int) (sentences string) {
	if count < 1 {
		return ""
	}
	sentences = Sentence(words)
	for i := 1; i < count; i++ {
		sentences = sentences + " " + Sentence(words)
	}
	return
}

// Paragraphs returns a string with multiple sentences grouped in paragraphs.
func Paragraphs(count, sentences, words int) (paragraphs string) {
	if count < 1 {
		return ""
	}
	paragraphs = Sentences(sentences, words)
	for i := 1; i < count; i++ {
		paragraphs = paragraphs + "\n\n" + Sentences(sentences, words)
	}
	return
}
