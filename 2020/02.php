<?php

$handle = fopen("input", "r");
$valids = 0;
$total = 0;
while ($data = fscanf($handle, "%d-%d %c: %s\n")) {
    $total++;
    list ($min, $max, $letter, $pass) = $data;
    if (!$pass) break;
    $count = 0;

    $count += $pass[$max - 1] === $letter;
    $count += $pass[$min - 1] === $letter;

    $valids += $count === 1;
}
fclose($handle);

var_dump($valids, $total);
