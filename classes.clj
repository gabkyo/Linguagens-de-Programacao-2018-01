(ns classes)

(defn Lote []
  (let [
    id (atom nil)
    owner (atom nil)

    set (fn [i o]
          (do (reset! id i)
              (reset! owner o)))
    ID (fn [] [@id])
    OWNER (fn [] [@owner])
    PRECO (fn [] [0])
    ]
    (fn [m]
      (cond (= m :set) set
            (= m :ID) ID
            (= m :OWNER) OWNER
            (= m :PRECO) PRECO))))

(defn Apto []
  (derive Apto Lote)
  )
