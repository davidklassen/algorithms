#lang racket

(define read-all (list->string (reverse
    (let loop ((char (read-char))
               (result '()))
        (if (eof-object? char)
            result
            (loop (read-char) (cons char result)))))))

(display read-all)
