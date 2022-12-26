import System.Environment
import Data.List

main :: IO ()
main = do
  args <- getArgs
  contents <- readFile $ head args
  putStrLn "Part 1:"
  print . sum . map (score . words) . lines $ contents
  putStrLn "\nPart 2:"
  print . sum . map (score'. words) . lines $ contents

score :: [String] -> Integer
score [a, b] = mod (yours - theirs + 1) 3 * 3 + yours
  where theirs = points a; yours = points b

score' :: [String] -> Integer
score' [a, b] = score [a, choose $ mod (points a + points b - 2) 3]

points :: String -> Integer
points "A" = 1 -- Rock
points "X" = 1 -- Lose
points "B" = 2 -- Paper
points "Y" = 2 -- Draw
points "C" = 3 -- Scissors
points "Z" = 3 -- Win
choose :: Integer -> String
choose n = head . reverse . take
  (fromInteger $ if n == 0 then 3 else n) $ ["A", "B", "C"]
