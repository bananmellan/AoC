import System.Environment
import Data.List
import Data.Char

main :: IO ()
main = do
  args <- getArgs
  contents <- readFile $ head args
  putStrLn "Part 1:"
  print . sum . map match
    . map (\bag -> (take (length bag `div` 2) bag, drop (length bag `div` 2) bag))
    . lines $ contents
  putStrLn "\nPart 2:"
  print . common $ lines $ contents

priority :: Char -> Int
priority item = ord item - if isUpper item then 38 else 96

match :: (String, String) -> Int
match (item:b1, b2) = if item `elem` b2 then priority item else match (b1, b2)
match _ = 0

common :: [String] -> Int
common (a:b:c:ds) = (priority . head $ foldl intersect a [b, c]) + common ds
common _          = 0
