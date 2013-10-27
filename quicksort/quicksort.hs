qsort :: Ord a => [a] -> [a]
qsort [] = []
qsort (x:xs) = qsort left ++ [x] ++ qsort right
               where left = [ a | a <- xs, a <= x ]
                     right = [ a | a <- xs, a > x ]

-- input handling

main = do
  content <- getContents
  let arr = map read $ lines content :: [Int]
  mapM_ print $ qsort arr

