Some files store information about their content encoded directly in their names. For example,
saved machine learning models weights. These files may include the epoch number or the value
of a tracked metric.

To extract these values from file paths, use regular exceptions. The most convenient way to do so is to
use named [regex placeholders](https://docs.python.org/3/library/re.html#regular-expression-syntax).
They allow to extract a dictionary of values from a string where the keys are matched parameters. 

```python
import ast
import re

def named_match(pattern, string):
    match = re.match(pattern, string)
    return {k: ast.literal_eval(v) for k, v in match.groupdict().items()}

filename = 'epoch=5_val_loss=0.23.ckpt'
result = named_match(r'^epoch=(?P<epoch>\d+)_val_loss=(?P<val_loss>\d+.\d+).ckpt$', filename)
assert result == {'epoch': '5', 'val_loss': '0.23'}
```

Sure enough, it is possible to go with anonymous parameters as well. But it makes the process 
more fragile because depends on the parameters' ordering.

