<?php

$handle = fopen("input", "r");
$maxy = 0;

while(!feof($handle)){
  $line = fgets($handle);
  $maxy++;
}

rewind($handle);

$trees = 0;
$x = 0;
$y = 0;
$maxx;
$steps = [
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2]
];

$product = 1;

var_dump($maxy);

foreach ($steps as $data) {
    list($right, $down) = $data;

    $trees = 0;
    $skip = 0;
    var_dump("---------($right, $down)--------");

    while ($data = fscanf($handle, "%s")) { if (is_bool($data)) break;
        var_dump("($x, $y)");
        $y++;
        if ($skip > 0) {
            var_dump('skipping');
            $skip--;
            continue;
        } else var_dump('not skipping');

        $line = $data[0];
        $maxx = strlen($line);
        $trees += $line[$x] === '#';
        $x += $right;
        $x = ($x + $maxx) % $maxx;
        $skip = $down - 1;
    } rewind($handle); $y = 0; $x = 0;

    $product *= $trees;
}

fclose($handle);
var_dump($product);
