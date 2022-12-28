import System.Environment
import Data.Char
import Control.Monad.State
import Data.Map (Map)
import qualified Data.Map as Map

type Pile = Map Integer String
type Stack = State Pile
type Instruction = (Integer, Integer, Integer)

main :: IO ()
main = do
  args <- getArgs
  contents <- readFile $ head args
  putStrLn "Part 1:"

  -- Fetch the lines representing the initial state.
  let initst = takeWhile (\line -> not (null line)) . lines $ contents

  -- Parse initial state and put it in lists.
  let parsedst = execState (construct . map divise . init $ initst) Map.empty

  -- Parse instructions and put them in a neat format.
  let instrs = map (\[_, n, _, p1, _, p2] -> (read n, read p1, read p2))
               . map words . tail . dropWhile (\line -> not (null line))
               . lines $ contents

  -- Perform instructions for CrateMover 9000
  let st = execState (instruct $ instrs) parsedst
  putStrLn $ tops st

  putStrLn "\nPart 2:"

  -- Perform instructions for CrateMover 9001
  let st' = execState (instruct' $ instrs) parsedst
  putStrLn $ tops st'
  where
    tops :: Pile -> String
    tops st = map (\(_,c) -> c) . Map.toList $ Map.map (\s -> head s) st

divise :: String -> String
divise ""   = []
divise line = (last $ take 2 line) : (divise $ drop 4 line)

instruct :: [Instruction] -> Stack ()
instruct ((0, _, _):rest) = instruct rest
instruct ((n, from, to):rest) = do
  val <- pop from
  push to val
  instruct $ (n - 1, from, to):rest
instruct _ = return ()

instruct' :: [Instruction] -> Stack ()
instruct' ((n, from, to):rest) = do
  vals <- pop' (fromInteger n) from
  push' to vals
  instruct' rest
instruct' _ = return ()

construct :: [String] -> Stack ()
construct (row:rows) = do
  parse 1 row
  construct rows
  where
    parse :: Integer -> String -> Stack ()
    parse n (' ':vals) = parse (n + 1) vals
    parse n (val:vals) = do
      push n val
      parse (n + 1) vals
    parse _ _ = return ()
construct rows = modify $ \st -> (Map.map (\s -> reverse s) st)

pop :: Integer -> Stack Char
pop i = do
  cols <- get
  let col = Map.findWithDefault [] i cols
  modify $ \st -> Map.insert i (tail col) cols
  return $ head col

push :: Integer -> Char -> Stack ()
push n val =
  modify $ \cols -> Map.insert n (val : Map.findWithDefault [] n cols) cols

pop' :: Int -> Integer -> Stack String
pop' n i = do
  cols <- get
  let col = Map.findWithDefault [] i cols
  modify $ \st -> Map.insert i (drop n col) cols
  return $ take n col

push' :: Integer -> String -> Stack ()
push' i vals =
  modify $ \cols -> Map.insert i (vals ++ Map.findWithDefault [] i cols) cols
