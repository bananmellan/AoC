<?php

$val = [
    'ecl' => true,
    'pid' => true,
    'eyr' => true,
    'hcl' => true,
    'byr' => true,
    'iyr' => true,
    'cid' => false,
    'hgt' => true
];

$file = fopen('input', 'r');
$contents = fread($file, filesize('input'));
fclose($file);

$lines = explode(PHP_EOL, $contents);
$valids = 0;
$text = '';
foreach ($lines as $line) {

    if (strlen($line) === 0) {
        $valid = true;

        foreach ($val as $key => $required) {
            if ($required && !isset($keys[$key])) {
                $valid = false;
                break;
            }

            else if ($required && isset($keys[$key])) {
                $value = $keys[$key];
                $error = null;

                switch ($key) {
                case 'byr':
                    $value = (int)$value;
                    $valid = $value >= 1920 && $value <= 2002;
                    break;
                case 'iyr':
                    $value = (int)$value;
                    $valid = $value >= 2010 && $value <= 2020;
                    break;
                case 'eyr':
                    $value = (int)$value;
                    $valid = $value >= 2020 && $value <= 2030;
                    break;
                case 'hgt':
                    $number = '';
                    $unit = null;

                    foreach (str_split($value) as $index => $char) {
                        if (is_numeric($char))
                            $number .= $char;
                        else {
                            $unit = substr($value, $index - 1);
                        }
                    }

                    $number = (int)$number;

                    switch ($unit) {
                    case 'in':
                        $valid = $number >= 59 && $number <= 76;
                        break;
                    case 'cm':
                        $valid = $number >= 150 && $number <= 193;
                        break;
                    default:
                        $valid = false;
                    }

                    $value = "($number $unit)";
                    break;
                case 'hcl':
                    $valid = preg_match('/^#[0-9a-f]{6}$/', $value);
                    break;
                case 'ecl':
                    $valid = preg_match('/^amb|blu|brn|gry|grn|hzl|oth$/', $value);
                    break;
                case 'pid':
                    $valid = preg_match('/^0*[0-9]{9}$/', $value);
                    break;
                }

                $keys[$key] = ($valid ? '' : 'in') . 'valid  ' . " '$value'";
                if (!$valid) break;
            }
        }

        $valids += $valid ? 1 : 0;

        $text = '';
        $keys = [];
    }

    $ex = explode(' ', $line);

    foreach ($ex as $id) {
        $pair = explode(':', $id);
        if (count($pair) !== 2) continue;
        $keys[$pair[0]] = $pair[1];
    }
}

echo "Valid passports: $valids\n";
