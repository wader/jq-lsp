include "file";
import "file" as mod;
import "file" as $mod;

# this is a comment

# this is a multiline \
comment

def ns: mod::not_found;

def binop: $mod + $mod;
def update1: . |= $mod;
def update1: . += $mod;
def alt_op: 123 // $mod;
def dest_alt_op1: 123 as $v ?// $v | $v;
def dest_alt_op1: 123 as $v ?// $b | $v;
def dest_alt_op2: 123 as $v ?// $b | $b;

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

def array1: [];
def array2: [1];
def array3: 1 as $var | [$var];

def query1: (func);

def comma1: func, func;

def unary1: +func, -func;

def string1: "test \(func(1)) \(func)";

def try1: func?;
def try2: try func catch func;

def if1: 1 as $var | if $var then $var end;
def if2: if func then func end;
def if3: if func then func else func end;
def if3: if func then func elif func then func else func end;

def foreach1: foreach 1 as $var ($var;$var);
def foreach1($var): foreach 1 as {$var: $var2} ($var2;$var);
def foreach2: foreach 1 as $var ($var;$var;$var);
def foreach2($var): foreach 1 as {$var: $var2} ($var2;$var;$var);

def reduce1: reduce 1 as $var ($var;$var);
def reduce2($var): foreach 1 as {$var: $var2} ($var2;$var);

def label1: label $out | break $out | 1;

def object_val_query: {a: 1+2+3};

def recurse1: ..;
