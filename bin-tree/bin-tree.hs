module Main where

data BinTree a = Leaf |
                 Node { value :: a
                      , left :: BinTree a
                      , right :: BinTree a
                      } deriving Show

fromList :: Ord a => [a] -> BinTree a
fromList (x:[]) = Node x Leaf Leaf
fromList (x:xs) = insert (fromList xs) x

insert :: Ord a => BinTree a -> a -> BinTree a
insert Leaf x = Node x Leaf Leaf
insert (Node val left right) x
    | x < val   = Node val (insert left x) right
    | otherwise = Node val left (insert right x)

main :: IO ()
main = print $ fromList [7, 3, 5, 4, 1]
