# Trigramerator

A basic implementation of a service to generate Trigrams.

## Trigrams

In natural language processing, a _trigram_ is a sequence of three consecutive words in a given body of text. For example, the sentence "To be or not to be, that is the question" contains the following trigrams:

```
[to, be, or]
[be, or, not]
[or, not, to]
[not, to, be]
[to, be, that]
[be, that, is]
[that, is, the]
[is, the, question]
```

Given a series of trigrams, it is possible to generate a random piece of text that resembles the original. For example, if we start with the two words "to be", we can find two trigrams that match that prefix:

```
[to, be, or]
[to, be, that]
```

At this point, we can make a random choice, and end up with "to be that". We can repeat the process by taking the last two words, "be that", and looking up what words may come next. Although with such a small example we are very limited in our choice, given a large enough body of text, such as a novel, we can produce text that, while completely devoid of meaning, appears to be in the same style as the novel that it was trained on.

## Design

### Database

The most critical design choice for this service has to do with the storage: for simplicity, and for an MVP, the decision was to use a simple in memory database, where the unerlying data 
structure would be a hash map of string to lists of strings. 

The map keys are the first 2 words or a trigram (e.g. `"to be"`), and it is mapped to a list of all the words the service found that follow the trigram key. The benefit is that by randomly selecting a word in this list, it is more likely to get words that follow the key more often - hence we have a natural weighted probability.

The main tradeoff here is memory efficiency: the list will grow indefinitely with the number of words - with a lot of repeated words. For simple - not so long - texts, this works fine, but this would not work for learning more extensive pieces of text (such as a full book or collection of books).

For this approach, a more sensible (more complex) solution could be to use a map instead of a slice: each third word of the trigram would appear only once (as the key), and the value would be the number of times it occurred in the text. An algorithm to compute weighted probabilities based on these values would then run to select the random word.

### Generated text size

It is a conscious design decision to keep the size of the texts short. We are assuming that we want just a one paragraph excerpt generated from the learned trigrams.

We could easily improve this to generate large pieces of text over multiple paragraphs, by preserving carriage returns (newlines) from the original text as one of the suffixes.

### Router

The router used in this project was taken from the [Go Service Template](https://github.com/tgmendes/go-service-template). It's a wrapper around [httprouter](https://github.com/julienschmidt/httprouter) and makes it easy to define the endpoints for the API. Without this template, the standard Go http router would be more than suitable for this project.

### Further Improvements
* More thoroughly tested and robust service would be to cover edge cases and incorrect formats (what if someone submits a text with only 2 words?). There was a conscious decision of not tackling these cases for now, in the interest of time.

* Capitalization of words is also very naive at this point - for example, it doesn't take into account quotes. 

* Learning a text can be done asynchronously (i.e. the user can send a text to learn and this can process in the background) using a go routine. However, the learning process (even big books) is quite fast, and for simplicity the decision was made to run this task synchronously.

* The text may finish on random words that might not make sense.

* Integration tests.

## Running the program

The trigram service was built using `go 1.15.8` using Go modules.

* Running: `go run app/main.go`
* Testing: `go test ./...`

## Testing

The following endpoints are exposed:

* `http://localhost:8080/learn` - `POST` with a text to learn trigrams (e.g. `curl --data-binary @harrypotter_test.txt localhost:8080/learn`).
* `http://localhost:8080/generate` - `GET` to generate a random piece of text.