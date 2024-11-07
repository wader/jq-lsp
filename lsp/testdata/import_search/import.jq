include "b" {search: "a"};
import "b" as ns1 {search: "a"};
import "b" as ns2 {search: ["a"]};
import "b" as ns3 {search: ["missing","a"]};
import "b" as $name1 {search: "a"};
import "b" as $name2 {search: ["a"]};

c,
ns1::c,
ns2::c,
ns3::c,
$name1,
$name2
