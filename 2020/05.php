<?php

$file = fopen('input', 'r');
$contents = fread($file, filesize('input'));
fclose($file);
$lines = explode(PHP_EOL, $contents);
$ids = [];
$max = 0;

function charBinPart($zero = '0', $one = '1', string $text) : int {
    return bindec(str_replace([$zero, $one], [0, 1], $text));
}

foreach ($lines as $line) {
    $row = charBinPart('F', 'B', substr($line, 0, 7));
    $col = charBinPart('L', 'R', substr($line, 7, 9));

    $product = $row * 8 + $col;

    array_push($ids, $product);

    $max = $product > $max ? $product : $max;
}

sort($ids);

foreach ($ids as $index => $id) {
    if (!isset($ids[$index + 1])) break;

    $nid = $ids[$index + 1];

    if (($id + 1) !== $nid) var_dump($id + 1);
}
