```json
{
  "category": "blog",
  "title": "Stress-Free Machine Learning,",
  "tags": ["python", "code"],
  "foreign_url": "https://towardsdatascience.com/stress-free-machine-learning-e4d0a411a56a"
}
```

<!--preamble-->

Building machine learning models isn’t easy. Heavy datasets and tricky data formats. 
A ton of hyper-parameters, models, optimization algorithms. Not talking about 
generic programming adventures like debugging, handling exceptions, and logging.
It is especially true for the R&D (Research and Development) style of work when 
models, approaches (and sometimes, even the data itself!) can change very quickly. 
You wouldn’t like to invest a large amount of time and effort to build something 
that could become irrelevant very soon. But you also don’t want to turn your 
project into a pile of Jupyter notebooks and ad hoc scripts, with a ton of files 
scattered here and there across the file system. In this post, I would like to 
share an approach that worked for me pretty well during several quickly evolving 
projects I’ve worked on recently. We’ll pick up a couple of image datasets and 
try to find a model architecture that works best for them all.

<!--more-->