(ns lp4.main
  (:gen-class :main true))
(load "class")
(load "list")
(load "file")

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (def lista (seq nil))
  (def lista (catalogo lista))
  (def lista (atual lista))
  (def lista (sort sortImoveis lista))
  ;(println lista)
  ;(println (count lista))
  (espec lista))
