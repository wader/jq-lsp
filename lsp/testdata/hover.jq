def fn: 123;
def fn2(a): 123;

def test:
  ( fromjson
  | fn
  | fn2(1)
  );