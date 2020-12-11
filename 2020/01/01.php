<?php

$file = fopen('input', 'r');
$contents = fread($file, filesize('input'));
fclose($file);

$nums = explode(PHP_EOL, $contents);

foreach ($nums as $x) {
    foreach ($nums as $y) {
        foreach ($nums as $z) {
            if (((int)$x + (int)$y + (int)$z) == 2020)
                var_dump($x*$y);
        }
    }
}
