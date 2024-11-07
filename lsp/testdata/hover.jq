def fn: 123;
def fn2(a): 123;

def test:
  ( fromjson
  | @sh
  | fn
  | fn2(1)
  );