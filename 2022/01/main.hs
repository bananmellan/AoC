import Data.List

main :: IO ()
main = do
  contents <- readFile "input"
  putStrLn "Part 1:"
  print . head . reverse . sort . map sum . bags [] $ lines $ contents
  putStrLn "\nPart 2:"
  print . sum . take 3 . reverse . sort . map sum . bags [] $ lines $ contents

bags :: [Integer] -> [String] -> [[Integer]]
bags b ("":s) = b : bags [] s
bags b (s:ss) = bags (read s:b) $ ss
bags b _ = [b]
