```json
{
  "category": "blog",
  "title": "Python Errors Done Right",
  "tags": ["python", "code"],
  "foreign_url": "https://medium.com/swlh/python-errors-done-right-faa1bfa85d02"
}
```

<!--preamble-->

Exceptions mechanism is widely adopted among modern programming languages, and 
Python is not an exception. (Pun intended!) Though this topic could seem obvious, 
like wrap your code blocks with try-catch clauses, and that’s all, there are 
some minor but important details, and taking them into account should make your 
code a bit cleaner. In this post, I’m going through some guidelines on how to 
structure error processing in Python that I derived from my personal experience.

<!--more-->