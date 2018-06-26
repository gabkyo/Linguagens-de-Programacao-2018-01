(ns lp4.main
  (:gen-class :main true))
(load "class")
(load "list")

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (def a (Terreno. 1 "Roger" "retang" "A" 10 10 0 10))
  (def b (Casa. 2 "OK" 1 1 2 10 10 10 10))
  (def c (Apto. 3 "lol" 1 1 2 10 10 "S" 10))
  (def teste (vector c b a))
  (println (Preco a) (Preco b) (Preco c))
  (def ok (sort sortImoveis teste))
  (println ok))
