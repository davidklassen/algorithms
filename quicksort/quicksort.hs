module Main where

qsort :: [Int] -> [Int]
qsort [] = []
qsort (x:xs) = qsort (filter (<=x) xs) ++ [x] ++ qsort (filter (>x) xs)

main :: IO ()
main = interact (unlines . map show . qsort . map read . lines)
