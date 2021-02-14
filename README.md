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

The main tradoff here is memory efficiency: the list will grow indefinitely with the number of words - with a lot of repetition. For simple - not so long - texts, this works fine, but this would not work for learning more extensive pieces of text (such as a full book or collection of books).

For this approach, a more sensible (more complex) solution could be to use a map instead of a slice: each third word of the trigram would appear only once (as the key), and the value would be the number of times it occurred in the text. An algorithm to compute weighted probabilities based on these values would then run to select the random word.

### Generator stop point

The current generator can potentially run indefinitely, leaving the user to wait for a long period of time. Adding a break point to avoid long running times would be a sensible approach. In the 
interest of time, a breaker was not added in this example.

### Further Improvements


One of the next steps in having a more thoroughly tested and robust service would be to cover edge cases and incorrect formats (what if someone submits a text with only 2 words?). There was a conscious decision of not tackling these cases for now, in the interest of time.

Capitalization of words is also very naive at this point - and doesn't take into account quotes, and also doesn't work for some newlines. This would be another area of improvement in the future.

## Running the program

The trigram service was built using `go 1.15.8` and `go modules`.

To run locally, please ensure you have these setup, and run `go mod install` before running the program. After that, the program should run with `go run app/web/main.go` from the root of the project.

To run all tests: `go test ./...`.

