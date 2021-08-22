
def func: 123;
def func($a): $a;
def func($a; $b): $a + $b;

def call1: func;
def call2: func(1);
def call3: func(1;2);

def op1: func + func;

def var: 123 as $var | $var;

def pattern1: . as {$var} | $var;
def pattern2: . as {var: $var} | $var;
def pattern2: . as [$var] | $var;
def pattern4:  . as $var | . as {("\($var)"): $var2} | $var2;

def object1: 1 as $var | {$var};
def object1: 1 as $var | {var: $var};
def object2: 1 as $var | {("\($var)"): $var};

def array1: 1 as $var | [$var];

def query1: (func);

def comma1: func, func;

def unary1: +func;

def string1: "test \(func(1)) \(func)";

def try1: func?;
def try2: try func catch func;

def if1: 1 as $var | if $var then $var end;
def if2: if func then func end;
def if3: if func then func else func end;
def if3: if func then func elif func then func else func end;

def foreach1: foreach 1 as $var ($var;$var;$var);
def foreach1($var): foreach 1 as {$var: $var2} ($var2;$var;$var);

def reduce1: reduce 1 as $var ($var;$var);
def reduce2($var): foreach 1 as {$var: $var2} ($var2;$var);

def label1: label $out | break $out | 1;
