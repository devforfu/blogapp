Let's pretend we have a function that creates a bar plot of categorical variable 
counts using a given data frame and a name of its column. 

```python
import pandas as pd

def load_data():
    return pd.DataFrame({'x': list('aababc')})

def plot_count_bars(data: pd.DataFrame, column: str):
    cnt = data[column].value_counts().rename('counts').reset_index()
    return cnt.plot.bar(x='index', y='counts')

plot_count_bars(load_data(), column='x')
``` 

Looks good! However, what if we later create another function that generates
a similar plot but using dots connected with a line instead of bars?
```python
def plot_count_line(data: pd.DataFrame, column: str):
    cnt = data[column].value_counts().rename('counts').reset_index()
    return cnt.plot.line(x='index', y='counts', marker='o')

plot_count_line(load_data(), column='x')
```

It works as well but see that the counting code is duplicated twice. So it is 
better to keep the computational logic separate from plotting. For example, 
create a new function called `counts` that takes a data frame and generates
counts. Then, modify the plotting functions, so they take the pre-computed 
values and perform the rendering.
```python
def counts(data: pd.DataFrame, column: str):
    return data[column].value_counts().rename('counts').reset_index()

def plot_count_bars(cnt: pd.DataFrame):
    return cnt.plot.bar(x='index', y='counts')

def plot_count_line(cnt: pd.DataFrame):
    return cnt.plot.line(x='index', y='counts', marker='o')

cnt = counts(load_data())
plot_count_bars(cnt)
plot_count_line(cnt)
```

The example is intentionally super-simple but this tip gets much more relevant
for complex cases.  