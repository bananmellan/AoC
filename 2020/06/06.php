<?php

$file = fopen('input', 'r');
$contents = fread($file, filesize('input'));
fclose($file);
$lines = explode(PHP_EOL, $contents);

$questions = range('a', 'z');

$groups = [];
$groupid = 0;

foreach ($lines as $i => $line) {
    if (strlen($line) === 0) {
        $groupid++;
        continue;
    }

    else if (!isset($groups[$groupid])) {
        $groups[$groupid] = [];
    }

    $group = &$groups[$groupid];

    foreach ($questions as $char)
        if (!isset($group[$char])) $group[$char] = true;

    foreach ($questions as $char) {
        $flag = &$group[$char];

        if ($flag === true) {
            $flag = strpos($line, $char) !== false;
        }
    }
}

$count = 0;

foreach ($groups as &$group) {
    foreach ($group as $index => $yes) {
        if ($yes === true) $count++;
        else unset($group[$index]);
    }
}

var_dump($groups, $count);
