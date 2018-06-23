(ns lp4.main
  (:gen-class :main true))
(load "class")

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (let [a (Terreno. 64 "Roger" "retang" "A" 10 10 0 10)] (println (Owner a ) (Preco a))))
