foreach (
    $missins,
    $var,
    split("")
) as $_ (
    $also;
    $not;
    $defined
),
reduce (
    $missins,
    $var
) as $_ (
    $also;
    $not
)