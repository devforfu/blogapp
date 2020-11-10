Jupyter keeps everything in global scope. Every variable or function that was defined 
during session is stored until explicitly deleted or the kernel is reloaded. Of course,
each imported module and functions sits in the scope as well. Therefore, it is a good
habit to reload the kernel every once in a while. Especially, before committing the
results into version control system or running a scheduled job.

