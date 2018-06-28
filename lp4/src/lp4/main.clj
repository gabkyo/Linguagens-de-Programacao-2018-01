(ns lp4.main
  (:gen-class :main true))
(load "class")
(load "list")
(load "file")

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (def lista (catalogo))
  ;(def lista (sort sortImoveis lista))
  (println lista))
