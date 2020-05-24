```json
{
  "category": "blog", 
  "title": "Using Python's __new__ Method to Dynamically Switch Class Implementations",
  "tags": ["python", "magic", "meta-class"],
  "image": "new.png"
}
```

<!--preamble-->

Every Python object has a set of magic methods that could be overridden to
customize an instance creation and behavior. One of the widely used methods is
`__init__` which is used to perform the newly created instance's initialization.
But there is one more magic method that takes part in the object creation. The 
method is called `__new__`. This method actually _creates_ an instance of a class. 
The post explains how to use this method to dynamically switch the implementation 
of class' logic while hiding it under the same class name.

<!--more-->

<div class="list-of-contents">
  <h4>Post contents</h4>
  <ul></ul>
</div>

<h3 class="header" id="intro">General Information</h3>

The [Python language reference](https://docs.python.org/3/reference/datamodel.html#object.__new__)
explains the purpose of `__new__` method as being a special-cased static method that
takes the class `cls` of which an instance was required as its first argument and
should return the new instance of that class.

One of the main differences of this method from `__init__` is that it actually
_creates a new instance_, and not just binds its properties, or makes some kind of
object tuning. Therefore, appropriately overridden, this method allows one
to customize object's creation process and, for example, switch the actual
implementation of a created instance.

This method allows creating a "facade" class that declares the interface but
which implementation could be chosen at runtime based on some criterion, i.e.
using the name of a specific algorithm to use, a set of parameters, etc. This class
"hides" actual implementation and allows to use its name instead of hard-coding
a name of specific class. Talking about the languages with the static typing systems,
one could compare (from a logical point of view) this approach with coercing
the specific object types to the generic interface type and using it instead as
the following snippet shows.

```Swift
public protocol StringConvertible {
    func toString() -> String
}

struct Color: StringConvertible {
  let r: Int
  let g: Int
  let b: Int  
  func toString() -> String {
      return "Color(r: \(r), g: \(g), b: \(b))"
  }
}

struct Point: StringConvertible {
  let x: Float
  let y: Float
  func toString() -> String {
      return "Point(x: \(x), y: \(y))"
  }
}

let convertibles: [StringConvertible] = [
    Color(r: 255, g: 0, b: 0),
    Point(x: 1.5, y: 2.5)
]
for object in convertibles {
    print(object.toString())
}
```

Next section shows a simple example demonstrating one possible implementation
of the described idea. Of course, this example can be implemented in a lot
of different ways (using callback as parameter, direct inheritance, etc.), but the
major goal standing behind it is to show one of the possible usages of the 
`__new__` method overriding.

<h3 class="header" id="simple">Example: Temperature Converter</h3>

<blockquote class="tip">
<strong>TL;DR</strong>: There is <a href="https://gist.github.com/devforfu/1de5cecb96f92bd99ed595de7cdb7907">
the full source code</a> of the example shown in this post as a single Python file.
</blockquote>

Consider the following example. You need to create a family of temperature
converter classes that should be able to convert from [Kelvin temperature degrees](https://en.wikipedia.org/wiki/Kelvin) 
into different measurement scales. All these classes should share the same set of methods but apply the
different temperature converting formulas.

Let's start with the showing a short use case of the codebase we're going to create.
Here we create a list of temperature converters. Each of them has its own implementation of
the temperature converting logic. We hide these implementations under the common
`TemperatureConverter` interface, and apply them to the same temperature value.

```Python
def main():
    converters = [
        TemperatureConverter(convert_to=name)
        for name in ('celsius', 'fahrenheit')]
    temperature = 300
    for converter in converters:
        string = converter.format(temperature)
        print('%s converted %sK temperature into: %s' % (
            converter.name, temperature, string
        ))

if __name__ == '__main__':
    main()
```

The snippet above should generate the following output.
```bash
$ python main.py
CelsiusConverter converted 300K temperature into: 26.85 (°C)
FahrenheitConverter converted 300K temperature into: 80.33 (°F)
$ _
```

To implement the dynamic algorithm dispatching using the approach proposed in 
the previous section, one should create a base class with the properly overridden 
`__new__` method. A possible implementation of the base class that is used to
instantiate different types of converters depending on the value of `convert_to`
parameter goes below.
```python
class TemperatureConverter(metaclass=abc.ABCMeta):    
    """
    The base temperature converter class that implements the dynamic 
    substitution of the implementations.
    """
    symbol = 'K'

    def __new__(cls, convert_to='celsius'):
        if issubclass(cls, TemperatureConverter):
            if convert_to == 'celsius':
                cls = _CelsiusConverter
            elif convert_to == 'fahrenheit':
                cls = _FahrenheitConverter
            else:
                raise ValueError('unexpected converter: %s' % convert_to)
        return object.__new__(cls)

    def convert(self, value):        
        self.check_value(value)
        return self._convert(value)

    def format(self, value):
        return ('%.2f' % self.convert(value)) + self.symbol

    @staticmethod
    def check_value(value):
        if value < 0:
            raise ValueError('temperature should be provided in Kelvin degrees')

    @property
    def name(self):
        return self.__class__.__name__

    def _convert(self, value):
        raise NotImplementedError()
```

See how the `__new__` looks like, showing the dynamic dispatching of implementations. 
The parameter `cls` mentioned at the beginning of this post refers to one of
specific implementations of the converter class. Then an instance of that class is created. 
Also, check the public `convert()` method that makes some basic sanity check
of the provided temperature value. Then it calls the "protected" `_convert()` method that
should be overridden in the derived classes.

Now, it is time to write the implementations of the aforementioned specific classes.
Note that `_CelsiusConverter` class is written like you usually create the derived classes
in Python by explicit derivation from the base class. But `_FahrenheitConverter`
does not inherit from the `TemperatureConverter`. Instead, it is registered as a subclass via
`ABCMeta.register` method. Though in this case, as soon as classes don't share the same
hierarchy, the default implementations of the required methods are not available 
anymore and should be written from the scratch.

```python
class _CelsiusConverter(TemperatureConverter):
    """
    A concrete implementation of the temperature converter that converts from
    Kelvin into Celsius degrees.
    """
    symbol = '°C'

    def _convert(self, value):
        return value - 273.15

class _FahrenheitConverter:
    """
    A concrete implementation of the temperature converter that converts from
    Kelvin into Fahrenheit degrees.
    
    Note that this class does not directly inherit the base class but is
    registered as a derived class using ABCMeta.register method. In this
    case, there are no default implementations of the `format` method and
    the `name` property.
    """
    symbol = '°F'

    def _convert(self, value):
        return value * 9./5 - 459.67

    def format(self, value):
        return '%.2f (%s)' % (self._convert(value), self.symbol)

    @property
    def name(self):
        return 'FahrenheitConverter'

TemperatureConverter.register(_FahrenheitConverter)

```

This example shows the flexibility of duck typing where you can write custom
classes and plug-in them it into the hierarchy. The only requirement is to make 
sure that the `__new__` method appropriately handles the side extensions. A some
kind of registry or lookup is requires instead of explicit enumeration of all 
the cases using `if-else` clauses. Also, one could override the `__init__` magic
methods in the derived classes to allow the custom arguments specific for their 
internal logic. 

The following section shows an example of using the dictionary lookup instead of 
hard-coding class names and their initializers.

<h3 class="header" id="notifications">Dictionary Lookup of Registered Classes</h3>

<blockquote class="tip">
<strong>TL;DR</strong>: Again,
<a href="https://gist.github.com/devforfu/63fa7efe18133dc12f12a2e9ecbe9db4">
here is a link</a> to the source code discussed below.
</blockquote>

This example shows almost the same approach as before, but with a more realistic
example that I had on practice. Consider the case when you need to implement 
a class that should extract user's notification messages from a server. 
Before implementing the actual server API, it is convenient to build a testing 
implementation. This implementation reads the locally stored file with the 
notifications instead of firing the real network requests and converts them 
into a proper format. 

In order to switch between two different implementations (a local file mockery and a real server),
you only need to pass different configuration dictionaries like `NotificationsDispatcher(**test_config)`
The appropriate dispatcher's instance is created under the hood.

```Python
class NotificationsDispatcher(metaclass=abc.ABCMeta):
    """
    Class that retrieves a list of notifications for a specific user.
    Usually, notifications are retrieved from the remote server, but for testing
    purposes and local runs it supports reading messages from local source.
    """

    def __new__(cls, user_id: int, method: str='http', **kwargs):      
        if issubclass(cls, NotificationsDispatcher):
            cls = get_dispatcher(method)
        return object.__new__(cls)

    def __init__(self, user_id: int, dateformat='%m/%d/%Y %H:%M:%S', **kwargs):
        self.user_id = user_id
        self.dateformat = dateformat

    @abc.abstractmethod
    def get_notifications(self):
        pass
```

See that this time instead of the exhaustive `if-else` chain, we use a helper 
function that looks up `method` name in the registry and returns the appropriate 
class object (if it exists). This approach allows to write a system of plugins 
in some other modules and register them via publicly exposed interface.

```Python
# somewhere dictionary of dispatchers exists
from weakref import WeakValueDictionary
_dispatchers = WeakValueDictionary()

def get_dispatcher(name):
    """
    Public API to access dictionary with registered dispatchers.
    """
    if name not in _dispatchers:
        raise ValueError('dispatcher with name \'%s\' is not found' % name)
    return _dispatchers[name]

def register_dispatcher(name, cls):
    """
    Public API which is used to register new dispatcher class.
    """
    global _dispatchers
    _dispatchers[name] = cls

register_dispatcher('local', _LocalDispatcher)
register_dispatcher('http')
```

It is possible to override `__init__` methods to accept different set of 
arguments, depending on the actual class' implementation. For example, local 
notifications dispatcher requires the name of file with the mocked notification 
but a remote one requires the server's URL.

<h3 class="header" id="pandas">Real World Example from Pandas Library</h3>

In order to conclude the discussion, let's check a real world example of 
described technique implemented in the [Pandas library repository](https://github.com/pandas-dev/pandas/blob/master/pandas/io/excel.py#L853-L875).

The library includes a couple of classes responsible for writing the data frames 
into Excel files. The abstract class is called `ExcelWriter`. Its simplified 
implementation shown below overrides the `__new__` method to pick up an appropriate
Excel writing library depending on the provided file extension, **.xls** or
**.xlsx**.

```Python
class ExcelWriter(meta=abc.ABCMeta):
    """
    A simplified version of pandas.ExcelWriter.__new__ implementation.
    """

    def __new__(cls, path, engine=None, **kwargs):
        if subclass(cls, ExcelWriter):
            if engine is None:
                ext = 'xlsx'
            try:
                engine = config.get_option('io.excel.%s.writer' % ext)
            except KeyError:
                raise ValueError('No engine for filetype: %s' % ext)
            cls = get_writer(engine)
        return object.__new__(cls)

    @abc.abstractproperty
    def engine(self):
        pass

    @abc.abstractmethod
    def write_cells(self, cells, sheet_name=None, startrow=0, startcol=0,
                    freeze_panes=None):
        pass

    # other interface methods and properties ...
```

<h3 class="header" id="conclusion">Conclusion</h3>

The dynamic nature of Python's typing system doesn't only allow to override methods
and bind new attributes to an object but also to completely substitute its
implementation in the runtime based on the configuration parameters. Quite often
the same result can be achieved using the composition instead of inheritance, or
using the callback functions. However, if used properly, the discussed approach 
allows to clearly separate different implementations from each other while keeping
the same interface and class name by switching the configuration parameters making
the codebase cleaner and easier to maintain.

<hr class="with-margin">

### References

<ol>
  <li><a href="https://docs.python.org/3/reference/datamodel.html">Python Data Model Documentation</a></li>
  <li><a href="https://github.com/pandas-dev/pandas">Pandas Repository</a></li>
</ol>
