#lang racket

(define (qsort arr)
  (if (null? arr) '()
  (let ((left (filter (lambda (el) (<= el (car arr))) (cdr arr)))
        (right (filter (lambda (el) (> el (car arr))) (cdr arr))))
    (append (qsort left) (list (car arr)) (qsort right)))))

; input handling

(define (arr)
  (let ((line (read-line)))
    (if (eof-object? line) '()
    (cons (string->number line) (arr)))))

(for ([i (qsort (arr))]) (display i) (newline))
