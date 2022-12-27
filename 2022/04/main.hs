import System.Environment
import Data.List
import Data.Char

main :: IO ()
main = do
  args <- getArgs
  contents <- readFile $ head args
  putStrLn "Part 1:"
  print . sum . map encompass . convert . lines $ contents
  putStrLn "\nPart 2:"
  print . sum . map overlap   . convert . lines $ contents
  where
    convert
      = map (\nums -> map (\num -> read num) nums)
      . map (\line -> words $ map (\c -> if c `elem` "-," then ' ' else c) line)


encompass :: [Integer] -> Integer
encompass [a, b, c, d] = if a' == c' || b' == c' then 1 else 0
  where c' = a' `intersect` b'; a' = [a..b]; b' = [c..d]

overlap :: [Integer] -> Integer
overlap [a, b, c, d] = if not $ null c' then 1 else 0
  where c' = a' `intersect` b'; a' = [a..b]; b' = [c..d]
