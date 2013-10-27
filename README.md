# Реализации различных алгоритмов и структур данных на языках Haskell, Scheme и JavaScript

## Список алгоритмов и структур данных

* **radix-sort:** Поразрядная сортировка (Цифровая сортировка) — алгоритм сортировки за линейное время
* **avl-tree:** АВЛ-дерево — сбалансированное по высоте двоичное дерево поиска
* **ford-fulkerson:** Алгоритм Форда — Фалкерсона решает задачу нахождения максимального потока в транспортной сети
* **quicksort:** Быстрая сортировка — один из самых быстрых известных универсальных алгоритмов сортировки массивов (в среднем O(n log n) обменов при упорядочении n элементов)
* **floyd-warshall:** Алгоритм Флойда — Уоршелла — динамический алгоритм для нахождения кратчайших расстояний между всеми вершинами взвешенного ориентированного графа
* **lzw:** Алгори́тм Ле́мпеля — Зи́ва — Ве́лча (Lempel-Ziv-Welch, LZW) — универсальный алгоритм сжатия данных без потерь
* **heapsort:** Пирамидальная сортировка — алгоритм сортировки, работающий гарантированно за Θ(n log n) операций при сортировке n элементов
* **md5:** MD5 (Message Digest 5) — 128-битный алгоритм хеширования
* **huffman:** Алгоритм Хаффмана — адаптивный жадный алгоритм оптимального префиксного кодирования алфавита с минимальной избыточностью
* **rsa:** RSA — криптографический алгоритм с открытым ключом, основывающийся на вычислительной сложности задачи факторизации больших целых чисел
* **b-tree:** Б-дерево — сбалансированное, сильно ветвистое дерево поиска
* **red-black-tree:** Красно-чёрное дерево — это одно из самобалансирующихся двоичных деревьев поиска, гарантирующих логарифмический рост высоты дерева от числа узлов и быстро выполняющее основные операции дерева поиска: добавление, удаление и поиск узла
* **fsm:** Конечный автомат — абстрактный автомат без выходного потока, число возможных состояний которого конечно

## Запуск тестов

Для каждого алгоритма есть файл с тестовыми данными с расширением `*.in`, запускать следующим образом:

**JavaScript**
```
$ node {название\_алгоритма}.js < {название\_алгоритма}.in
```

**Haskell**
```
$ runhaskell {название\_алгоритма}.hs < {название\_алгоритма}.in
```

**Scheme**
```
$ racket {название\_алгоритма}.scm < {название\_алгоритма}.in
```
