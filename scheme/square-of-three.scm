;;; Write a function that takes three numbers and returns the square of the two highest

(define (sq-sum x y)
    (+ (* x x) (* y y)))

(define (gr-three a b c)
    (cond ((= a b c) (sq-sum a b))
        ((and (>= a b) (>= a c))
            (if (>= b c)
                (sq-sum a b)
                (sq-sum a c)))
        ((and (>= b c) (>= b a))
            (if (>= c a)
                (sq-sum b c)
                (sq-sum b a)))
        ((and (>= c a) (>= c b))
            (if (>= a b)
                (sq-sum c a)
                (sq-sum c b)))

    )
)
